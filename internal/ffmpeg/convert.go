package ffmpeg

import (
	"fmt"
	"os/exec"

	"github.com/photoprism/photoprism/internal/ffmpeg/apple"
	"github.com/photoprism/photoprism/internal/ffmpeg/encode"
	"github.com/photoprism/photoprism/internal/ffmpeg/intel"
	"github.com/photoprism/photoprism/internal/ffmpeg/nvidia"
	"github.com/photoprism/photoprism/internal/ffmpeg/v4l"
	"github.com/photoprism/photoprism/internal/ffmpeg/vaapi"
	"github.com/photoprism/photoprism/pkg/fs"
)

// AvcConvertCmd returns the command for converting video files to MPEG-4 AVC.
func AvcConvertCmd(srcName, destName string, opt encode.Options) (cmd *exec.Cmd, useMutex bool, err error) {
	if srcName == "" {
		return nil, false, fmt.Errorf("empty source filename")
	} else if destName == "" {
		return nil, false, fmt.Errorf("empty destination filename")
	}

	// Don't transcode more than one video at the same time.
	useMutex = true

	// Use default ffmpeg command name.
	if opt.Bin == "" {
		opt.Bin = DefaultBin
	}

	// Don't use hardware transcoding for animated images.
	if fs.TypeAnimated[fs.FileType(srcName)] != "" {
		cmd = exec.Command(
			opt.Bin,
			"-y",
			"-strict", "-2",
			"-i", srcName,
			"-pix_fmt", encode.FormatYUV420P.String(),
			"-vf", "scale='trunc(iw/2)*2:trunc(ih/2)*2'",
			"-f", "mp4",
			"-movflags", "+faststart", // puts headers at the beginning for faster streaming
			destName,
		)

		return cmd, useMutex, nil
	}

	// Display encoder info.
	if opt.Encoder != encode.SoftwareAvc {
		log.Infof("convert: ffmpeg encoder %s selected", opt.Encoder.String())
	}

	switch opt.Encoder {
	case encode.IntelAvc:
		cmd = intel.AvcConvertCmd(srcName, destName, opt)

	case encode.AppleAvc:
		cmd = apple.AvcConvertCmd(srcName, destName, opt)

	case encode.VaapiAvc:
		cmd = vaapi.AvcConvertCmd(srcName, destName, opt)

	case encode.NvidiaAvc:
		cmd = nvidia.AvcConvertCmd(srcName, destName, opt)

	case encode.V4LAvc:
		cmd = v4l.AvcConvertCmd(srcName, destName, opt)

	default:
		cmd = encode.AvcConvertCmd(srcName, destName, opt)
	}

	return cmd, useMutex, nil
}
