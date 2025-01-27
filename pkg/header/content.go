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
	ContentTypeAVC       = "video/mp4; codecs=\"avc1\""
	ContentTypeHEVC      = "video/mp4; codecs=\"hvc1.1.6.L93.90\""
	ContentTypeOGG       = "video/ogg"
	ContentTypeWebM      = "video/webm"
	ContentTypeVP8       = "video/webm; codecs=\"vp8\""
	ContentTypeVP9       = "video/webm; codecs=\"vp9\""
	ContentTypeAV1       = "video/webm; codecs=\"av01.0.08M.08\""
)
