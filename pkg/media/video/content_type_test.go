package video

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/photoprism/photoprism/pkg/media/http/header"
)

func TestContentType(t *testing.T) {
	t.Run("QuickTime", func(t *testing.T) {
		assert.Equal(t, fs.MimeTypeMp4, ContentType(fs.MimeTypeMov, "", ""))
	})
	t.Run("QuickTime_HVC", func(t *testing.T) {
		assert.Equal(t, header.ContentTypeMp4Hevc, ContentType(fs.MimeTypeMov, "mov", CodecHevc))
	})
	t.Run("Mp4", func(t *testing.T) {
		assert.Equal(t, fs.MimeTypeMp4, ContentType(fs.MimeTypeMp4, "", ""))
	})
	t.Run("Mp4_AVC", func(t *testing.T) {
		assert.Equal(t, header.ContentTypeMp4Avc, ContentType(fs.MimeTypeMp4, "", CodecAvc))
	})
	t.Run("Mp4_HVC", func(t *testing.T) {
		assert.Equal(t, header.ContentTypeMp4Hevc, ContentType(fs.MimeTypeMp4, "", CodecHevc))
	})
}
