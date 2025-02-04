package video

type CodecProfile struct {
	Codec      Codec
	Profile    string
	Level      int
	SubLevel   int
	Bitrate    int
	Resolution int
	Framerate  int
	ID         string
}

type CodecProfiles []CodecProfile
