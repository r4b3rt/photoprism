package header

/*
	Standard content types for use in HTTP headers and the web interface.

	Browser support can be tested on one or more of the following sites:
	- https://ott.dolby.com/codec_test/index.html
	- https://dmnsgn.github.io/media-codecs/
	- https://cconcolato.github.io/media-mime-support/
	- https://thorium.rocks/misc/h265-tester.html
*/

// Standard ContentType identifiers for audio and video files.
const (
	ContentTypeMov       = "video/quicktime"
	ContentTypeMp4       = "video/mp4"
	ContentTypeMp4Avc720 = ContentTypeMp4 + "; codecs=\"avc1.640020\""      // MPEG-4 AVC, High Profile Level 3.2
	ContentTypeMp4Avc    = ContentTypeMp4 + "; codecs=\"avc1.640028\""      // MPEG-4 AVC, High Profile Level 4.0
	ContentTypeMp4Hevc   = ContentTypeMp4 + "; codecs=\"hvc1.2.4.L120.B0\"" // HEVC Mp4 Main10 Profile, Main Tier, Level 4.0
	ContentTypeMp4Hev1   = ContentTypeMp4 + "; codecs=\"hev1.2.4.L120.B0\"" // HEVC Bitstream, not supported on macOS
	ContentTypeMp4Vvc    = ContentTypeMp4 + "; codecs=\"vvc1\""             // Versatile Video Coding (VVC), also known as H.266
	ContentTypeMp4Evc    = ContentTypeMp4 + "; codecs=\"evc1\""             // MPEG-5 Essential Video Coding (EVC), also known as ISO/IEC 23094-1
	ContentTypeTheora    = "video/ogg"
	ContentTypeWebm      = "video/webm"
	ContentTypeWebmVp8   = "video/webm; codecs=\"vp8\""
	ContentTypeWebmVp9   = "video/webm; codecs=\"vp09.00.10.08\""
	ContentTypeWebmAv1   = "video/webm; codecs=\"av01.2.10M.10\""
)

// Standard ContentType identifiers for images and vector graphics.
const (
	ContentTypePng  = "image/png"
	ContentTypeJpeg = "image/jpeg"
	ContentTypeSVG  = "image/svg+xml"
)

// Standard ContentType identifiers for markup and sidecar files.
const (
	ContentTypeBinary    = "application/octet-stream"
	ContentTypeForm      = "application/x-www-form-urlencoded"
	ContentTypeMultipart = "multipart/form-data"
	ContentTypeJson      = "application/json"
	ContentTypeJsonUtf8  = "application/json; charset=utf-8"
	ContentTypeHtml      = "text/html; charset=utf-8"
	ContentTypeText      = "text/plain; charset=utf-8"
	ContentTypePDF       = "application/pdf"
)
