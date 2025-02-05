package header

import (
	"crypto/rand"
	"fmt"
	"hash/crc32"
	"log"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const CharsetBase62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomToken generates a random hexadecimal character token for authenticating client applications.
//
// Examples: 9fa8e562564dac91b96881040e98f6719212a1a364e0bb25
func RandomToken() string {
	b := make([]byte, 24)

	if _, err := rand.Read(b); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", b)
}

// RandomAppPassword generates a random, human-friendly authentication token that can also be used as
// password replacement for client applications. It is separated by 3 dashes for better readability
// and has a total length of 27 characters.
//
// Example: OXiV72-wTtiL9-d04jO7-X7XP4p
func RandomAppPassword() string {
	m := big.NewInt(int64(len(CharsetBase62)))
	b := make([]byte, 0, 27)

	for i := 0; i < 27; i++ {
		if (i+1)%7 == 0 {
			b = append(b, '-')
		} else if i == 27-1 {
			b = append(b, CharsetBase62[crc32.ChecksumIEEE(b)%62])
			return string(b)
		} else if r, err := rand.Int(rand.Reader, m); err == nil {
			b = append(b, CharsetBase62[r.Int64()])
		}
	}

	return string(b)
}

func TestAuth(t *testing.T) {
	t.Run("Header", func(t *testing.T) {
		assert.Equal(t, "X-Auth-Token", XAuthToken)
		assert.Equal(t, "X-Session-ID", XSessionID)
		assert.Equal(t, "Authorization", Auth)
	})
	t.Run("Values", func(t *testing.T) {
		assert.Equal(t, "Basic", AuthBasic)
		assert.Equal(t, "Bearer", AuthBearer)
	})
}

func TestAuthToken(t *testing.T) {
	t.Run("None", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}

		// No headers have been set, so no token should be returned.
		token := AuthToken(c)
		assert.Equal(t, "", token)
	})
	t.Run("BearerToken", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}

		// Set Bearer Authorization header to a random value generated by RandomToken().
		expected := RandomToken()
		SetAuthorization(c.Request, expected)

		// Check header for expected token.
		authToken := AuthToken(c)
		assert.Equal(t, expected, authToken)
		bearerToken := BearerToken(c)
		assert.Equal(t, authToken, bearerToken)
	})
	t.Run("XAuthToken", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}

		// Set X-Auth-Token header to a random value generated by RandomToken().
		expected := RandomToken()
		c.Request.Header.Add(XAuthToken, expected)

		// Check header for expected token.
		authToken := AuthToken(c)
		assert.Equal(t, expected, authToken)
		bearerToken := BearerToken(c)
		assert.Equal(t, "", bearerToken)
	})
	t.Run("XSessionID", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}

		// Set X-Session-ID header to a random value generated by RandomToken().
		expected := RandomToken()
		c.Request.Header.Add(XSessionID, expected)

		// Check header for expected token.
		authToken := AuthToken(c)
		assert.Equal(t, expected, authToken)
		bearerToken := BearerToken(c)
		assert.Equal(t, "", bearerToken)
	})
	t.Run("AppPassword", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}

		// Set X-Auth-Token header to a random value generated by RandomAppPassword().
		expected := RandomAppPassword()
		c.Request.Header.Add(XAuthToken, expected)

		// Check header for expected token.
		authToken := AuthToken(c)
		assert.Equal(t, expected, authToken)
		bearerToken := BearerToken(c)
		assert.Equal(t, "", bearerToken)
	})
}

func TestBearerToken(t *testing.T) {
	t.Run("None", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}

		// No headers have been set, so no token should be returned.
		token := BearerToken(c)
		assert.Equal(t, "", token)
	})
	t.Run("Found", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}

		// Add authorization header.
		SetAuthorization(c.Request, "69be27ac5ca305b394046a83f6fda18167ca3d3f2dbe7ac0")

		// Check result.
		token := BearerToken(c)
		assert.Equal(t, "69be27ac5ca305b394046a83f6fda18167ca3d3f2dbe7ac0", token)
	})
}

func TestAuthorization(t *testing.T) {
	t.Run("None", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}

		// No headers have been set, so no token should be returned.
		authType, authToken := Authorization(c)
		assert.Equal(t, "", authType)
		assert.Equal(t, "", authToken)
	})
	t.Run("BearerToken", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}

		// Add authorization header.
		c.Request.Header.Add(Auth, "Bearer 69be27ac5ca305b394046a83f6fda18167ca3d3f2dbe7ac0")

		// Check result.
		authType, authToken := Authorization(c)
		assert.Equal(t, AuthBearer, authType)
		assert.Equal(t, "69be27ac5ca305b394046a83f6fda18167ca3d3f2dbe7ac0", authToken)
	})
}

func TestBasicAuth(t *testing.T) {
	t.Run("None", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}

		// No headers have been set, so no token should be returned.
		user, pass, key := BasicAuth(c)
		assert.Equal(t, "", user)
		assert.Equal(t, "", pass)
		assert.Equal(t, "", key)
	})
	t.Run("Found", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}

		// Add authorization header.
		c.Request.Header.Add(Auth, AuthBasic+" QWxhZGRpbjpvcGVuIHNlc2FtZQ==")

		// Check result.
		user, pass, key := BasicAuth(c)
		assert.Equal(t, "Aladdin", user)
		assert.Equal(t, "open sesame", pass)
		assert.Equal(t, "0cdb723383eb144043424a4a254461658d887396", key)
	})
}
