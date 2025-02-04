package meta

import (
	"github.com/photoprism/photoprism/pkg/media/video"
)

const CodecUnknown = ""
const CodecAv1 = video.CodecAv1
const CodecAvc1 = video.CodecAvc
const CodecHvc1 = video.CodecHevc
const CodecJpeg = "jpeg"
const CodecHeic = "heic"
const CodecXMP = "xmp"

// CodecAvc returns true if the video codec is AVC.
func (data Data) CodecAvc() bool {
	return data.Codec == CodecAvc1
}
