package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dustin/go-humanize/english"
	"github.com/gin-gonic/gin"

	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/event"
	"github.com/photoprism/photoprism/internal/form"
	"github.com/photoprism/photoprism/internal/photoprism/get"
	"github.com/photoprism/photoprism/internal/server/limiter"
	"github.com/photoprism/photoprism/pkg/authn"
	"github.com/photoprism/photoprism/pkg/clean"
	"github.com/photoprism/photoprism/pkg/media/http/header"
)

// OAuthToken creates a new access token for clients that authenticate with valid OAuth2 client credentials.
//
//	@Tags	Authentication
//	@Router	/api/v1/oauth/token [post]
func OAuthToken(router *gin.RouterGroup) {
	router.POST("/oauth/token", func(c *gin.Context) {
		// Prevent CDNs from caching this endpoint.
		if header.IsCdn(c.Request) {
			AbortNotFound(c)
			return
		}

		// Get client IP address for logs and rate limiting checks.
		clientIp := ClientIP(c)
		actor := "unknown client"
		action := "create token"

		// Abort if running in public mode.
		if get.Config().Public() {
			event.AuditErr([]string{clientIp, "oauth2", actor, action, authn.ErrDisabledInPublicMode.Error()})
			AbortForbidden(c)
			return
		}

		// Disable caching of responses.
		c.Header(header.CacheControl, header.CacheControlNoStore)

		// Token create request form.
		var frm form.OAuthCreateToken
		var sess *entity.Session
		var client *entity.Client
		var err error

		// Allow authentication with basic auth and form values.
		if clientId, clientSecret, _ := header.BasicAuth(c); clientId != "" && clientSecret != "" {
			frm.GrantType = authn.GrantClientCredentials
			frm.ClientID = clientId
			frm.ClientSecret = clientSecret
		} else if err = c.ShouldBind(&frm); err != nil {
			event.AuditWarn([]string{clientIp, "oauth2", actor, action, "%s"}, err)
			AbortBadRequest(c)
			return
		}

		// Check the credentials for completeness and the correct format.
		if err = frm.Validate(); err != nil {
			event.AuditWarn([]string{clientIp, "oauth2", actor, action, "%s"}, err)
			AbortInvalidCredentials(c)
			return
		}

		// Check request rate limit.
		r := limiter.Login.Request(clientIp)

		// Abort if request rate limit is exceeded.
		if r.Reject() || limiter.Auth.Reject(clientIp) {
			limiter.AbortJSON(c)
			return
		}

		if frm.ClientID != "" {
			actor = fmt.Sprintf("client %s", clean.Log(frm.ClientID))
		} else if frm.Username != "" {
			actor = fmt.Sprintf("user %s", clean.Log(frm.Username))
		} else if frm.GrantType == authn.GrantPassword {
			actor = "unknown user"
		}

		// Create a new session (access token) based on the grant type specified in the request.
		switch frm.GrantType {
		case authn.GrantClientCredentials, authn.GrantUndefined:
			// Find client with the specified ID.
			client = entity.FindClientByUID(frm.ClientID)

			// Check if a client has been found, it is enabled, and the credentials are valid.
			if client == nil {
				event.AuditWarn([]string{clientIp, "oauth2", actor, action, authn.ErrInvalidClientID.Error()})
				AbortInvalidCredentials(c)
				return
			} else if !client.AuthEnabled {
				event.AuditWarn([]string{clientIp, "oauth2", actor, action, authn.ErrAuthenticationDisabled.Error()})
				AbortInvalidCredentials(c)
				return
			} else if method := client.Method(); !method.IsDefault() && method != authn.MethodOAuth2 {
				event.AuditWarn([]string{clientIp, "oauth2", actor, action, "method %s not supported"}, clean.LogQuote(method.String()))
				AbortInvalidCredentials(c)
				return
			} else if client.InvalidSecret(frm.ClientSecret) {
				event.AuditWarn([]string{clientIp, "oauth2", actor, action, authn.ErrInvalidClientSecret.Error()})
				AbortInvalidCredentials(c)
				return
			}

			// Update time of last activity.
			client.UpdateLastActive(true)

			// Cancel failure rate limit reservation.
			r.Success()

			// Create new client session.
			sess = client.NewSession(c, authn.GrantClientCredentials)
		case authn.GrantPassword, authn.GrantSession:
			// Generate an app password for a user account and check the password for confirmation.
			s := Session(clientIp, AuthToken(c))

			if s == nil {
				AbortInvalidCredentials(c)
				return
			} else if s.Username() == "" || s.IsClient() || !s.IsRegistered() {
				event.AuditErr([]string{clientIp, "oauth2", actor, action, authn.ErrInvalidGrantType.Error()})
				AbortInvalidCredentials(c)
				return
			}

			actor = fmt.Sprintf("user %s", clean.Log(s.Username()))

			if s.User().Provider().SupportsPasswordAuthentication() {
				loginForm := form.Login{
					Username: s.Username(),
					Password: frm.Password,
				}

				authUser, authProvider, authMethod, authErr := entity.Auth(loginForm, nil, c)

				if authProvider.IsClient() {
					event.AuditErr([]string{clientIp, "oauth2", actor, action, authn.Denied})
					AbortInvalidCredentials(c)
					return
				} else if authMethod.Is(authn.Method2FA) && errors.Is(authErr, authn.ErrPasscodeRequired) {
					// Ok.
				} else if authErr != nil {
					event.AuditErr([]string{clientIp, "oauth2", actor, action, "%s"}, strings.ToLower(clean.Error(authErr)))
					AbortInvalidCredentials(c)
					return
				} else if !authUser.Equal(s.User()) {
					event.AuditErr([]string{clientIp, "oauth2", actor, action, authn.ErrUserDoesNotMatch.Error()})
					AbortInvalidCredentials(c)
					return
				}

				frm.GrantType = authn.GrantPassword
			} else {
				frm.GrantType = authn.GrantSession
			}

			sess = entity.NewClientSession(frm.ClientName, frm.ExpiresIn, frm.Scope, frm.GrantType, s.User())

			// Return the reserved request rate limit tokens after successful authentication.
			r.Success()
		default:
			event.AuditErr([]string{clientIp, "oauth2", actor, action, authn.ErrInvalidGrantType.Error()})
			AbortInvalidCredentials(c)
			return
		}

		// Save new session.
		if sess, err = get.Session().Save(sess); err != nil {
			event.AuditErr([]string{clientIp, "oauth2", actor, action, err.Error()})
			AbortInvalidCredentials(c)
			return
		} else if sess == nil {
			event.AuditErr([]string{clientIp, "oauth2", actor, action, StatusFailed.String()})
			AbortUnexpectedError(c)
			return
		} else {
			event.AuditInfo([]string{clientIp, "oauth2", actor, action, authn.Created})
		}

		// Delete any existing client sessions above the configured limit.
		if client == nil {
			// Skip deletion if not created by a client.
		} else if deleted := client.EnforceAuthTokenLimit(); deleted > 0 {
			event.AuditInfo([]string{clientIp, "oauth2", actor, action, "deleted %s to enforce token limit"}, english.Plural(deleted, "session", "sessions"))
		}

		// Send response with access token, token type, and token lifetime.
		response := gin.H{
			"status":       StatusSuccess,
			"session_id":   sess.ID,
			"access_token": sess.AuthToken(),
			"token_type":   sess.AuthTokenType(),
			"expires_in":   sess.ExpiresIn(),
			"client_name":  sess.ClientName,
			"scope":        sess.Scope(),
		}

		c.JSON(http.StatusOK, response)
	})
}
