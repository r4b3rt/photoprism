package video

import (
	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/photoprism/photoprism/pkg/media/http/header"
)

// Unknown represents an unknown and/or unsupported video format.
var Unknown = Type{
	Codec:    CodecUnknown,
	FileType: fs.TypeUnknown,
}

// Mp4 specifies the MPEG-4 Part 14 multimedia container format.
var Mp4 = Type{
	Codec:       CodecAvc,
	FileType:    fs.VideoMp4,
	ContentType: header.ContentTypeMp4Avc,
	WidthLimit:  8192,
	HeightLimit: 4320,
	Public:      true,
}

// Mov specifies the Apple QuickTime (QT) container format.
var Mov = Type{
	Codec:       CodecAvc,
	FileType:    fs.VideoMov,
	WidthLimit:  8192,
	HeightLimit: 4320,
	Public:      true,
}

// Avc specifies the MPEG-4 Advanced Video Coding (H.264) format.
var Avc = Type{
	Codec:       CodecAvc,
	FileType:    fs.VideoAvc,
	WidthLimit:  8192,
	HeightLimit: 4320,
	Public:      true,
}

// Hevc specifies the generally compatible High Efficiency Video Coding (H.265) format.
var Hevc = Type{
	Codec:       CodecHevc,
	FileType:    fs.VideoHevc,
	WidthLimit:  0,
	HeightLimit: 0,
	Public:      false,
}

// Hev1 specifies the HEVC bitstream format with the parameter sets stored in the samples (not supported on macOS):
// https://ott.dolby.com/codec_test/index.html
var Hev1 = Type{
	Codec:       CodecHev1,
	FileType:    fs.VideoHev1,
	WidthLimit:  0,
	HeightLimit: 0,
	Public:      false,
}

// Vvc specifies the Versatile Video Coding (H.266) format.
var Vvc = Type{
	Codec:       CodecVvc,
	FileType:    fs.VideoVvc,
	WidthLimit:  0,
	HeightLimit: 0,
	Public:      false,
}

// Evc specifies the Essential Video Coding (MPEG-5 Part 1) format.
var Evc = Type{
	Codec:       CodecEvc,
	FileType:    fs.VideoEvc,
	WidthLimit:  0,
	HeightLimit: 0,
	Public:      false,
}

// Vp8 specifies a Google VP8 video in a WebM multimedia container.
var Vp8 = Type{
	Codec:       CodecVp8,
	FileType:    fs.VideoWebm,
	WidthLimit:  0,
	HeightLimit: 0,
	Public:      false,
}

// Vp9 specifies a Google VP9 video in a WebM multimedia container.
var Vp9 = Type{
	Codec:       CodecVp9,
	FileType:    fs.VideoWebm,
	WidthLimit:  0,
	HeightLimit: 0,
	Public:      false,
}

// Av1 specifies an AV1 (AOMedia Video 1) video in a WebM multimedia container.
var Av1 = Type{
	Codec:       CodecAv1,
	FileType:    fs.VideoWebm,
	WidthLimit:  0,
	HeightLimit: 0,
	Public:      false,
}

// Theora specifies OGV video with Vorbis audio in an OGG multimedia container.
var Theora = Type{
	Codec:       CodecTheora,
	FileType:    fs.VideoTheora,
	WidthLimit:  0,
	HeightLimit: 0,
	Public:      false,
}

// Webm specifies the Google WebM multimedia container format.
var Webm = Type{
	Codec:       CodecUnknown,
	FileType:    fs.VideoWebm,
	WidthLimit:  0,
	HeightLimit: 0,
	Public:      false,
}
