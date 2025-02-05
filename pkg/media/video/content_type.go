package video

import (
	"fmt"
	"strings"

	"github.com/photoprism/photoprism/pkg/clean"
	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/photoprism/photoprism/pkg/media/http/header"
)

// ContentType returns a normalized video content type strings based on the video file type and codec.
func ContentType(mediaType, fileType, videoCodec string) string {
	if mediaType == "" && fileType == "" && videoCodec == "" {
		return header.ContentTypeBinary
	}

	if mediaType == "" {
		var c string

		if videoCodec != "" {
			c = Codecs[videoCodec]
		} else {
			c = fileType
		}

		switch c {
		case "mov", "mp4":
			mediaType = header.ContentTypeMp4
		case CodecAvc3:
			mediaType = header.ContentTypeMp4Avc3High // AVC (H.264) Bitstream, High Profile
		case string(fs.VideoAvc), CodecAvc:
			mediaType = header.ContentTypeMp4AvcHigh // Advanced Video Coding (H.264), High Profile
		case string(fs.VideoHevc), CodecHevc, "hvc":
			mediaType = header.ContentTypeMp4Hevc // High Efficiency Video Coding (H.265)
		case CodecHev1, "hev":
			mediaType = header.ContentTypeMp4Hev1 // High Efficiency Video Coding (HEVC) Bitstream
		case CodecVvc, "vvc":
			mediaType = header.ContentTypeMp4Vvc // Versatile Video Coding (VVC), also known as H.266
		case CodecEvc, "evc":
			mediaType = header.ContentTypeMp4Evc // MPEG-5 Essential Video Coding (EVC), also known as ISO/IEC 23094-1
		case CodecVp8, "vp08":
			mediaType = header.ContentTypeWebmVp8
		case CodecVp9, "vp9":
			mediaType = header.ContentTypeWebmVp9
		case CodecAv1, "av1":
			mediaType = header.ContentTypeWebmAv1
		case CodecTheora, "ogg":
			mediaType = header.ContentTypeOggTheora
		case string(fs.VideoWebm):
			mediaType = header.ContentTypeWebm
		}
	}

	// Add codec parameter, if possible.
	if mediaType != "" && !strings.Contains(mediaType, ";") {
		if codec, found := Codecs[videoCodec]; found && codec != "" {
			mediaType = fmt.Sprintf("%s; codecs=\"%s\"", mediaType, codec)
		}
	}

	return clean.ContentType(mediaType)
}
