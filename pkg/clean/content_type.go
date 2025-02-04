package clean

import (
	"strings"

	"github.com/photoprism/photoprism/pkg/media/http/header"
)

// ContentType normalizes media content type strings, see https://en.wikipedia.org/wiki/Media_type.
func ContentType(s string) string {
	if s == "" {
		return header.ContentTypeBinary
	}

	s = Type(s)

	// Replace "video/quicktime" with "video/mp4" as the container formats are largely compatible.
	s = strings.Replace(s, header.ContentTypeMov, header.ContentTypeMp4, 1)

	switch s {
	case "":
		return header.ContentTypeBinary
	case "text/json", "application/json":
		return header.ContentTypeJsonUtf8
	case "text/htm", "text/html":
		return header.ContentTypeHtml
	case "text/plain":
		return header.ContentTypeText
	case "text/pdf",
		"text/x-pdf",
		"application/x-pdf",
		"application/acrobat":
		return header.ContentTypePDF
	case "image/svg":
		return header.ContentTypeSVG
	case "image/jpe", "image/jpg":
		return header.ContentTypeJpeg
	case "video/mp4; codecs=\"avc\"",
		"video/mp4; codecs=\"avc1\"":
		return header.ContentTypeMp4Avc // Advanced Video Coding (AVC), also known as H.264
	case "video/mp4; codecs=\"hvc\"",
		"video/mp4; codecs=\"hvc1\"",
		"video/mp4; codecs=\"hevc\"":
		return header.ContentTypeMp4Hevc // HEVC Mp4 Main10 Profile
	case "video/mp4; codecs=\"hev\"",
		"video/mp4; codecs=\"hev1\"":
		return header.ContentTypeMp4Hev1 // HEVC bitstream with the parameter sets stored in the samples, not supported on macOS
	case "video/webm; codecs=\"vp8\"",
		"video/webm; codecs=\"vp08\"":
		return header.ContentTypeWebmVp8 // Google WebM container with VP8 video
	case "video/webm; codecs=\"vp9\"",
		"video/webm; codecs=\"vp09\"":
		return header.ContentTypeWebmVp9 // Google WebM container with VP9 video
	case "video/webm; codecs=\"av1\"",
		"video/webm; codecs=\"av01\"":
		return header.ContentTypeWebmAv1 // Google WebM container with AV1 video
	}

	return s
}
