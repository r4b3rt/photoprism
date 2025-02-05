package video

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/photoprism/photoprism/pkg/media/http/header"
)

func TestContentType(t *testing.T) {
	t.Run("QuickTime", func(t *testing.T) {
		assert.Equal(t, header.ContentTypeMp4, ContentType(header.ContentTypeMov, "", ""))
	})
	t.Run("QuickTime_HVC", func(t *testing.T) {
		assert.Equal(t, header.ContentTypeMp4HvcMain, ContentType(header.ContentTypeMov, "mov", CodecHvc))
	})
	t.Run("Mp4", func(t *testing.T) {
		assert.Equal(t, header.ContentTypeMp4, ContentType(header.ContentTypeMp4, "", ""))
	})
	t.Run("Mp4_AVC", func(t *testing.T) {
		assert.Equal(t, header.ContentTypeMp4AvcMain, ContentType(header.ContentTypeMp4, "", CodecAvc))
	})
	t.Run("Mp4_HVC", func(t *testing.T) {
		assert.Equal(t, header.ContentTypeMp4HvcMain, ContentType(header.ContentTypeMp4, "", CodecHvc))
	})
}
