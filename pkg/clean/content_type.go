package clean

import (
	"github.com/photoprism/photoprism/pkg/header"
)

// ContentType normalizes media content type strings, see https://en.wikipedia.org/wiki/Media_type.
func ContentType(s string) string {
	if s == "" {
		return header.ContentTypeBinary
	}

	s = Type(s)

	switch s {
	case "":
		return header.ContentTypeBinary
	case "text/json", "application/json":
		return header.ContentTypeJsonUtf8
	case "text/htm", "text/html":
		return header.ContentTypeHtml
	case "text/plain":
		return header.ContentTypeText
	case "text/pdf", "text/x-pdf", "application/x-pdf", "application/acrobat":
		return header.ContentTypePDF
	case "image/svg":
		return header.ContentTypeSVG
	case "image/jpe", "image/jpg":
		return header.ContentTypeJPEG
	case "video/mp4; codecs=\"avc\"":
		return header.ContentTypeAVC
	case "video/mp4; codecs=\"hvc1\"", "video/mp4; codecs=\"hvc\"", "video/mp4; codecs=\"hevc\"":
		return header.ContentTypeHEVC
	case "video/webm; codecs=\"av01\"":
		return header.ContentTypeAV1
	}

	return s
}
