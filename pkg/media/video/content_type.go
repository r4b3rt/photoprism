package video

import (
	"fmt"
	"mime"
	"strings"

	"github.com/photoprism/photoprism/pkg/clean"
	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/photoprism/photoprism/pkg/media/http/header"
)

// ContentType returns a normalized video content type strings based on the video file type and codec.
func ContentType(mediaType, fileType, videoCodec string, hdr bool) string {
	if mediaType == "" && fileType == "" && videoCodec == "" {
		return header.ContentTypeBinary
	}

	if mediaType == "" {
		videoCodec = Codecs[videoCodec]

		if fs.VideoMov.Equal(fileType) {
			fileType = fs.VideoMp4.String()
		}

		switch {
		case fs.VideoMp4.Equal(fileType) && videoCodec == CodecAvc3:
			if hdr {
				mediaType = header.ContentTypeMp4Avc3High // AVC (H.264) Bitstream, High Profile
			} else {
				mediaType = header.ContentTypeMp4Avc3Main // AVC (H.264) Bitstream, High Profile
			}
		case fs.VideoAvc.Equal(fileType) || fs.VideoMp4.Equal(fileType) && videoCodec == CodecAvc:
			if hdr {
				mediaType = header.ContentTypeMp4AvcHigh // MPEG-4 AVC, High Profile
			} else {
				mediaType = header.ContentTypeMp4AvcMain // MPEG-4 AVC, Main Profile
			}
		case fs.VideoHvc.Equal(fileType) || fs.VideoMp4.Equal(fileType) && videoCodec == CodecHvc:
			if hdr {
				mediaType = header.ContentTypeMp4HvcMain10 // MPEG-4 HEVC, Main 10-Bit HDR
			} else {
				mediaType = header.ContentTypeMp4HvcMain // MPEG-4 HEVC, Main
			}
		case fs.VideoHev.Equal(fileType) || fs.VideoMp4.Equal(fileType) && videoCodec == CodecHev:
			if hdr {
				mediaType = header.ContentTypeMp4HevMain10 // HEVC Bitstream, Main 10-Bit HDR
			} else {
				mediaType = header.ContentTypeMp4HevMain // HEVC Bitstream, Main
			}
		case fs.VideoVvc.Equal(fileType) || fs.VideoMp4.Equal(fileType) && videoCodec == CodecVvc:
			mediaType = header.ContentTypeMp4Vvc // Versatile Video Coding (VVC), also known as H.266
		case fs.VideoEvc.Equal(fileType) || fs.VideoMp4.Equal(fileType) && videoCodec == CodecEvc:
			mediaType = header.ContentTypeMp4Evc // MPEG-5 Essential Video Coding (EVC), also known as ISO/IEC 23094-1
		case videoCodec == CodecVp8:
			mediaType = header.ContentTypeWebmVp8
		case videoCodec == CodecVp9:
			mediaType = header.ContentTypeWebmVp9
		case fs.VideoAv1.Equal(fileType) || fs.VideoWebm.Equal(fileType) && videoCodec == CodecAv1:
			mediaType = header.ContentTypeWebmAv1
		case fs.VideoMkv.Equal(fileType) && videoCodec == CodecAv1:
			mediaType = header.ContentTypeMkvAv1
		case fs.VideoTheora.Equal(fileType) || videoCodec == CodecTheora:
			mediaType = header.ContentTypeOggTheora
		case fs.VideoWebm.Equal(fileType):
			mediaType = header.ContentTypeWebm
		case fs.VideoMp4.Equal(fileType):
			mediaType = header.ContentTypeMp4
		case fs.VideoMkv.Equal(fileType):
			mediaType = header.ContentTypeMkv
		}
	}

	// Add codec parameter, if possible.
	if mediaType != "" && !strings.Contains(mediaType, ";") {
		if codec, found := Codecs[videoCodec]; found && codec != "" {
			mediaType = fmt.Sprintf("%s; codecs=\"%s\"", mediaType, codec)
		}
	}

	if hdr {
		switch mediaType {
		case
			header.ContentTypeMovAvc,
			header.ContentTypeMovAvcMain,
			header.ContentTypeMp4Avc,
			header.ContentTypeMp4AvcBaseline,
			header.ContentTypeMp4AvcMain:
			if Codecs[videoCodec] == CodecAvc3 {
				mediaType = header.ContentTypeMp4Avc3High
			} else {
				mediaType = header.ContentTypeMp4AvcHigh
			}
		case
			header.ContentTypeMp4Avc3,
			header.ContentTypeMp4Avc3Main:
			mediaType = header.ContentTypeMp4Avc3High
		case
			header.ContentTypeMp4Hvc,
			header.ContentTypeMp4HvcMain:
			mediaType = header.ContentTypeMp4HvcMain10
		case
			header.ContentTypeMp4Hev,
			header.ContentTypeMp4HevMain:
			mediaType = header.ContentTypeMp4HevMain10
		}
	}

	return clean.ContentType(mediaType)
}

// Compatible tests if the video content types are expected to be compatible,
func Compatible(contentType1, contentType2 string) bool {
	// Content is likely compatible if the content type strings match exactly (case-insensitive).
	if contentType1 == "" || contentType2 == "" {
		return false
	} else if strings.EqualFold(contentType1, contentType2) {
		return true
	}

	// Sanitize and normalize content type strings.
	contentType1 = clean.ContentType(contentType1)
	contentType2 = clean.ContentType(contentType2)

	// Parse content type strings.
	mediaType1, params1, err1 := mime.ParseMediaType(contentType1)
	mediaType2, params2, err2 := mime.ParseMediaType(contentType2)

	// If parsing fails, assume the content is invalid or incompatible.
	if err1 != nil || err2 != nil {
		return false
	} else if len(params1) == 0 && len(params2) == 0 {
		return strings.EqualFold(mediaType1, mediaType2)
	}

	// If the media types don't match, assume the content is incompatible.
	if !strings.EqualFold(mediaType1, mediaType2) {
		return false
	}

	// Compare the media codecs.
	codec1 := params1["codecs"]
	codec2 := params2["codecs"]

	// Content is likely compatible if the full codec details match (case-insensitive).
	if strings.EqualFold(codec1, codec2) {
		return true
	}

	// Compare main codec names.
	codec1, _, _ = strings.Cut(codec1, ",")
	codec2, _, _ = strings.Cut(codec2, ",")

	codecName1, _, _ := strings.Cut(strings.TrimSpace(codec1), ".")
	codecName2, _, _ := strings.Cut(strings.TrimSpace(codec2), ".")

	// Content is likely compatible if the name of the main codec matches (case-insensitive).
	return strings.EqualFold(codecName1, codecName2)
}
