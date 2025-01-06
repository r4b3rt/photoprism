package config

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientAssets_Load(t *testing.T) {
	c := NewConfig(CliTestContext())

	t.Run("Ok", func(t *testing.T) {
		testBuildPath := "testdata/static/build"
		a := NewClientAssets(testBuildPath, c.StaticUri())

		err := a.Load("assets.json")

		assert.NoError(t, err)

		assert.Equal(t, "/static", a.BaseUri)
		assert.Equal(t, "app.6ea2755a30e3f9eb1169.css", a.AppCss) // app.6ea2755a30e3f9eb1169.css
		assert.Equal(t, "/static/build/app.6ea2755a30e3f9eb1169.css", a.AppCssUri())
		assert.Equal(t, "app.dbde125e18ba925d22fe.js", a.AppJs)
		assert.Equal(t, "/static/build/app.dbde125e18ba925d22fe.js", a.AppJsUri())
		assert.Equal(t, "share.2259c0edcc020e7af593.css", a.ShareCss)
		assert.Equal(t, "/static/build/share.2259c0edcc020e7af593.css", a.ShareCssUri())
		assert.Equal(t, "share.61248f7eb0aa9c8a7b21.js", a.ShareJs)
		assert.Equal(t, "/static/build/share.61248f7eb0aa9c8a7b21.js", a.ShareJsUri())
		assert.Equal(t, "/static/build/splash.a62e8b4d5ec0c8dc4ed4.css", a.SplashCssUri())
		assert.Equal(t, "splash.a62e8b4d5ec0c8dc4ed4.css", a.SplashCssFile())
		assert.NotEmpty(t, a.SplashCssFileContents())
	})

	t.Run("Error", func(t *testing.T) {
		testBuildPath := "testdata/foo"
		a := NewClientAssets(testBuildPath, c.StaticUri())

		err := a.Load("assets.json")

		assert.Error(t, err)

		assert.Equal(t, "/static", a.BaseUri)
		assert.Equal(t, "", a.AppCss)
		assert.Equal(t, "", a.AppCssUri())
		assert.Equal(t, "", a.AppJs)
		assert.Equal(t, "", a.AppJsUri())
		assert.Equal(t, "", a.ShareCss)
		assert.Equal(t, "", a.ShareCssUri())
		assert.Equal(t, "", a.ShareJs)
		assert.Equal(t, "", a.ShareJsUri())
	})
}

func TestConfig_ClientAssets(t *testing.T) {
	c := NewConfig(CliTestContext())

	c.options.AssetsPath = "testdata"
	c.options.CdnUrl = "https://mycdn.com/foo/"
	c.SetWallpaperUri("default")

	a := c.ClientAssets()

	assert.Equal(t, "https://mycdn.com/foo/static", a.BaseUri)
	assert.Equal(t, "app.6ea2755a30e3f9eb1169.css", a.AppCss)
	assert.Equal(t, "https://mycdn.com/foo/static/build/app.6ea2755a30e3f9eb1169.css", a.AppCssUri())
	assert.Equal(t, "app.dbde125e18ba925d22fe.js", a.AppJs)
	assert.Equal(t, "https://mycdn.com/foo/static/build/app.dbde125e18ba925d22fe.js", a.AppJsUri())
	assert.Equal(t, "share.2259c0edcc020e7af593.css", a.ShareCss)
	assert.Equal(t, "https://mycdn.com/foo/static/build/share.2259c0edcc020e7af593.css", a.ShareCssUri())
	assert.Equal(t, "share.61248f7eb0aa9c8a7b21.js", a.ShareJs)
	assert.Equal(t, "https://mycdn.com/foo/static/build/share.61248f7eb0aa9c8a7b21.js", a.ShareJsUri())
	assert.Equal(t, "https://mycdn.com/foo/static/img/wallpaper/default.jpg", c.WallpaperUri())
	assert.Equal(t, "https://mycdn.com/foo/static/build/splash.a62e8b4d5ec0c8dc4ed4.css", a.SplashCssUri())
	assert.Equal(t, "splash.a62e8b4d5ec0c8dc4ed4.css", a.SplashCssFile())
	assert.NotEmpty(t, a.SplashCssFileContents())

	c.options.AssetsPath = "testdata/invalid"
	c.options.CdnUrl = ""
	c.options.SiteUrl = "http://myhost/foo"
	c.SetWallpaperUri("default")

	a = c.ClientAssets()

	assert.Equal(t, "/foo/static", a.BaseUri)
	assert.Equal(t, "", a.AppCss)
	assert.Equal(t, "", a.AppCssUri())
	assert.Equal(t, "", a.AppJs)
	assert.Equal(t, "", a.AppJsUri())
	assert.Equal(t, "", a.ShareCss)
	assert.Equal(t, "", a.ShareCssUri())
	assert.Equal(t, "", a.ShareJs)
	assert.Equal(t, "", a.ShareJsUri())
	assert.Equal(t, "", c.WallpaperUri())
}

func TestClientManifestUri(t *testing.T) {
	c := NewConfig(CliTestContext())

	assert.True(t, strings.HasPrefix(c.ClientManifestUri(), "/manifest.json?2e5b4b86"))

	c.options.SiteUrl = ""

	assert.True(t, strings.HasPrefix(c.ClientManifestUri(), "/manifest.json?2e5b4b86"))

	c.options.SiteUrl = "http://myhost/foo"

	assert.True(t, strings.HasPrefix(c.ClientManifestUri(), "/foo/manifest.json?2e5b4b86"))
}
