package fs

// TypeAnimated maps animated file types to their mime type.
var TypeAnimated = TypeMap{
	ImageGif:   MimeTypeGif,
	ImagePng:   MimeTypeAPng,
	ImageWebp:  MimeTypeWebp,
	ImageAvif:  MimeTypeAvifS,
	ImageAvifS: MimeTypeAvifS,
	ImageHeic:  MimeTypeHeicS,
	ImageHeicS: MimeTypeHeicS,
}
