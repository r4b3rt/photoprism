package header

// Standard content request and response header names.
const (
	Accept             = "Accept"
	AcceptEncoding     = "Accept-Encoding"
	AcceptLanguage     = "Accept-Language"
	AcceptRanges       = "Accept-Ranges"
	ContentType        = "Content-Type"
	ContentDisposition = "Content-Disposition"
	ContentEncoding    = "Content-Encoding"
	ContentRange       = "Content-Range"
	Location           = "Location"
	Origin             = "Origin"
	Vary               = "Vary"
)

// Standard ContentType header values.
//
// See https://dmnsgn.github.io/media-codecs/ and https://cconcolato.github.io/media-mime-support/
// for codec support in browsers and the corresponding profile strings to use in the content type.
const (
	ContentTypeBinary    = "application/octet-stream"
	ContentTypeForm      = "application/x-www-form-urlencoded"
	ContentTypeMultipart = "multipart/form-data"
	ContentTypeJson      = "application/json"
	ContentTypeJsonUtf8  = "application/json; charset=utf-8"
	ContentTypeHtml      = "text/html; charset=utf-8"
	ContentTypeText      = "text/plain; charset=utf-8"
	ContentTypePDF       = "application/pdf"
	ContentTypePNG       = "image/png"
	ContentTypeJPEG      = "image/jpeg"
	ContentTypeSVG       = "image/svg+xml"
	ContentTypeAVC       = "video/mp4; codecs=\"avc1\""             // Advanced Video Coding (AVC), also known as H.264
	ContentTypeHEVC      = "video/mp4; codecs=\"hev1.2.4.L120.B0\"" // HEVC Main10 Profile, Main Tier, Level 4.0
	ContentTypeVVC       = "video/mp4; codecs=\"vvc1\""             // Versatile Video Coding (VVC), also known as H.266
	ContentTypeEVC       = "video/mp4; codecs=\"evc1\""             // MPEG-5 Essential Video Coding (EVC), also known as ISO/IEC 23094-1
	ContentTypeOGV       = "video/ogg"
	ContentTypeWebM      = "video/webm"
	ContentTypeVP8       = "video/webm; codecs=\"vp08.02.41.10\""
	ContentTypeVP9       = "video/webm; codecs=\"vp09.00.50.08\""
	ContentTypeAV1       = "video/webm; codecs=\"av01.2.10M.10\""
)
