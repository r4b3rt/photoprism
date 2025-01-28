package video

// Types maps identifiers to standards.
var Types = Standards{
	"":       AVC,
	"mp4":    MP4,
	"mpeg4":  MP4,
	"avc":    AVC,
	"avc1":   AVC,
	"hevc":   HEVC,
	"hevC":   HEVC,
	"hvc":    HEVC,
	"hvc1":   HEVC,
	"v_hvc":  HEVC,
	"v_hvc1": HEVC,
	"hev":    HEVC,
	"hev1":   HEVC,
	"evc":    EVC,
	"evc1":   EVC,
	"evcC":   EVC,
	"vvc":    VVC,
	"vvc1":   VVC,
	"vvcC":   VVC,
	"vp8":    VP8,
	"vp08":   VP8,
	"vp80":   VP8,
	"vp9":    VP9,
	"vp09":   VP9,
	"vp90":   VP9,
	"av1":    AV1,
	"av01":   AV1,
	"ogg":    OGV,
	"ogv":    OGV,
	"webm":   WebM,
}

// Standards maps names to standardized formats.
type Standards map[string]Type
