package video

// Profile represents a video codec profile name,
// see https://en.wikipedia.org/wiki/Advanced_Video_Coding#Profiles.
type Profile = string

const (
	ProfileBaseline Profile = "BaseLine"
	ProfileMain     Profile = "Main"
	ProfileHigh     Profile = "High"
)

// Profiles contains common video codec profiles together with their ContentType ID,
// maximum bitrate, resolution, and frame rate (if known).
var Profiles = CodecProfiles{
	{Codec: CodecAvc, Profile: ProfileBaseline, Level: 30, Bitrate: 10000, ID: "avc1.66.30"}, // iOS friendly
	{Codec: CodecAvc, Profile: ProfileBaseline, Level: 30, Bitrate: 10000, ID: "avc1.42001e"},
	{Codec: CodecAvc, Profile: ProfileBaseline, Level: 31, Bitrate: 14000, ID: "avc1.42001f"},
	{Codec: CodecAvc, Profile: ProfileMain, Level: 30, Bitrate: 10000, ID: "avc1.77.30"}, // iOS friendly
	{Codec: CodecAvc, Profile: ProfileMain, Level: 30, Bitrate: 10000, ID: "avc1.4d001e"},
	{Codec: CodecAvc, Profile: ProfileMain, Level: 31, Bitrate: 14000, ID: "avc1.4d001f"},
	{Codec: CodecAvc, Profile: ProfileMain, Level: 40, Bitrate: 20000, ID: "avc1.4d0028"},
	{Codec: CodecAvc, Profile: ProfileHigh, Level: 31, Bitrate: 17500, ID: "avc1.64001f"},
	{Codec: CodecAvc, Profile: ProfileHigh, Level: 40, Bitrate: 25000, ID: "avc1.640028"},
	{Codec: CodecAvc, Profile: ProfileHigh, Level: 41, Bitrate: 62500, ID: "avc1.640029"},
}
