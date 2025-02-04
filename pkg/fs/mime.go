package fs

import (
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

const (
	MimeTypeUnknown = ""
	MimeTypeBinary  = "application/octet-stream"
	MimeTypeJpeg    = "image/jpeg"
	MimeTypeJpegXL  = "image/jxl"
	MimeTypePng     = "image/png"
	MimeTypeAPng    = "image/vnd.mozilla.apng"
	MimeTypeGif     = "image/gif"
	MimeTypeBmp     = "image/bmp"
	MimeTypeTiff    = "image/tiff"
	MimeTypeDNG     = "image/dng"
	MimeTypeAvif    = "image/avif"
	MimeTypeAvifS   = "image/avif-sequence"
	MimeTypeHeic    = "image/heic"
	MimeTypeHeicS   = "image/heic-sequence"
	MimeTypeWebp    = "image/webp"
	MimeTypeMp4     = "video/mp4"
	MimeTypeMov     = "video/quicktime"
	MimeTypeSVG     = "image/svg+xml"
	MimeTypeAI      = "application/vnd.adobe.illustrator"
	MimeTypePS      = "application/postscript"
	MimeTypeEPS     = "image/eps"
	MimeTypeText    = "text/plain"
	MimeTypeXml     = "text/xml"
	MimeTypeJson    = "application/json"
)

// MimeType returns the mimetype of a file, or an empty string if it could not be determined.
//
// The IANA and IETF use the term "media type", and consider the term "MIME type" to be obsolete,
// since media types have become used in contexts unrelated to email, such as HTTP:
// https://en.wikipedia.org/wiki/Media_type#Structure
func MimeType(filename string) (mimeType string) {
	if filename == "" {
		return MimeTypeUnknown
	}

	// Detect file type based on the filename extension.
	fileType := Extensions[strings.ToLower(filepath.Ext(filename))]

	// Determine mime type based on the extension for the following
	// formats, which otherwise cannot be reliably distinguished:
	switch fileType {
	// Apple QuickTime Container
	case VideoMov:
		return MimeTypeMov
	// MPEG-4 AVC Video
	case VideoAvc:
		return MimeTypeMp4
	// Adobe Digital Negative
	case ImageDNG:
		return MimeTypeDNG
	// Adobe PostScript
	case VectorPS:
		return MimeTypePS
	// Adobe Embedded PostScript
	case VectorEPS:
		return MimeTypeEPS
	// Adobe Illustrator
	case VectorAI:
		return MimeTypeAI
	// Scalable Vector Graphics
	case VectorSVG:
		return MimeTypeSVG
	}

	// Detect mime type based on the file content.
	detectedType, err := mimetype.DetectFile(filename)

	if detectedType != nil && err == nil {
		mimeType = detectedType.String()
	}

	// Treat "application/octet-stream" as unknown.
	if mimeType == MimeTypeBinary {
		mimeType = MimeTypeUnknown
	}

	// If it could be detected, try to determine mime type from extension:
	if mimeType == MimeTypeUnknown {
		switch fileType {
		// Mp4 Multimedia Container
		case VideoMp4, VideoHevc, VideoHev1:
			return MimeTypeMp4
		// AV1 Image File
		case ImageAvif:
			return MimeTypeAvif
		// AV1 Image File Sequence
		case ImageAvifS:
			return MimeTypeAvifS
		// High Efficiency Image Container
		case ImageHeic, ImageHeif:
			return MimeTypeHeic
		// High Efficiency Image Container Sequence
		case ImageHeicS:
			return MimeTypeHeicS
		}
	}

	return mimeType
}

// BaseType returns the media type string without any optional parameters.
func BaseType(mimeType string) string {
	if mimeType == "" {
		return ""
	}

	mimeType, _, _ = strings.Cut(mimeType, ";")

	return strings.ToLower(mimeType)
}

// IsType tests if the specified mime types are matching, except for any optional parameters.
func IsType(mime1, mime2 string) bool {
	if mime1 == mime2 {
		return true
	}

	return BaseType(mime1) == BaseType(mime2)
}
