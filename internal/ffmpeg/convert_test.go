package ffmpeg

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/photoprism/photoprism/internal/ffmpeg/encode"
	"github.com/photoprism/photoprism/pkg/fs"
)

func TestAvcConvertCmd(t *testing.T) {
	runCmd := func(t *testing.T, encoder encode.Encoder, srcName, destName string, cmd *exec.Cmd) {
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		cmd.Env = append(cmd.Env, []string{
			fmt.Sprintf("HOME=%s", fs.Abs("./testdata")),
		}...)

		// Transcode source media file to AVC.
		start := time.Now()
		if err := cmd.Run(); err != nil {
			if stderr.String() != "" {
				err = errors.New(stderr.String())
			}

			// Remove broken video file.
			if !fs.FileExists(destName) {
				// Do nothing.
			} else if removeErr := os.Remove(destName); removeErr != nil {
				t.Logf("%s: failed to remove %s after error (%s)", encoder, srcName, removeErr)
			}

			// Log ffmpeg output for debugging.
			if err.Error() != "" {
				t.Error(err)
				t.Fatalf("%s: failed to transcode %s [%s]", encoder, srcName, time.Since(start))
			}
		}

		// Log filename and transcoding time.
		t.Logf("%s: created %s [%s]", encoder, destName, time.Since(start))

		if removeErr := os.Remove(destName); removeErr != nil {
			t.Fatalf("%s: failed to remove %s after successful test (%s)", encoder, srcName, removeErr)
		}
	}

	t.Run("NoSource", func(t *testing.T) {
		opt := encode.Options{
			Bin:      "",
			Encoder:  "intel",
			Size:     1500,
			Bitrate:  "50M",
			MapVideo: MapVideoDefault,
			MapAudio: MapAudioDefault,
		}
		_, _, err := AvcConvertCmd("", "", opt)

		assert.Equal(t, "empty source filename", err.Error())
	})
	t.Run("NoDestination", func(t *testing.T) {
		opt := encode.Options{
			Bin:      "",
			Encoder:  "intel",
			Size:     1500,
			Bitrate:  "50M",
			MapVideo: MapVideoDefault,
			MapAudio: MapAudioDefault,
		}
		_, _, err := AvcConvertCmd("VID123.mov", "", opt)

		assert.Equal(t, "empty destination filename", err.Error())
	})
	t.Run("Animation", func(t *testing.T) {
		opt := encode.Options{
			Bin:      "",
			Encoder:  "intel",
			Size:     1500,
			Bitrate:  "50M",
			MapVideo: MapVideoDefault,
			MapAudio: MapAudioDefault,
		}
		r, _, err := AvcConvertCmd("VID123.gif", "VID123.gif.avc", opt)

		if err != nil {
			t.Fatal(err)
		}

		assert.Contains(t, r.String(), "bin/ffmpeg -y -strict -2 -i VID123.gif -pix_fmt yuv420p -vf scale='trunc(iw/2)*2:trunc(ih/2)*2' -f mp4 -movflags +faststart VID123.gif.avc")
	})
	t.Run("VP9", func(t *testing.T) {
		encoder := encode.SoftwareAvc

		opt := encode.Options{
			Bin:      "/usr/bin/ffmpeg",
			Encoder:  encoder,
			Size:     1500,
			Bitrate:  "50M",
			MapVideo: MapVideoDefault,
			MapAudio: MapAudioDefault,
		}

		srcName := fs.Abs("./testdata/25fps.vp9")
		destName := fs.Abs("./testdata/25fps.avc")

		cmd, _, err := AvcConvertCmd(srcName, destName, opt)

		if err != nil {
			t.Fatal(err)
		}

		cmdStr := cmd.String()
		cmdStr = strings.Replace(cmdStr, srcName, "SRC", 1)
		cmdStr = strings.Replace(cmdStr, destName, "DEST", 1)

		assert.Equal(t, "/usr/bin/ffmpeg -y -strict -2 -i SRC -c:v libx264 -map 0:v:0 -map 0:a:0? -c:a aac -vf scale='if(gte(iw,ih), min(1500, iw), -2):if(gte(iw,ih), -2, min(1500, ih))',format=yuv420p -max_muxing_queue_size 1024 -crf 23 -r 30 -b:v 50M -f mp4 -movflags +faststart DEST", cmdStr)

		// Performs software transcoding.
		runCmd(t, encoder, srcName, destName, cmd)
	})
	t.Run("Vaapi", func(t *testing.T) {
		encoder := encode.VaapiAvc
		opt := encode.Options{
			Bin:      "/usr/bin/ffmpeg",
			Encoder:  encoder,
			Size:     1500,
			Bitrate:  "50M",
			MapVideo: MapVideoDefault,
			MapAudio: MapAudioDefault,
		}

		srcName := fs.Abs("./testdata/25fps.vp9")
		destName := fs.Abs("./testdata/25fps.vaapi.avc")

		cmd, _, err := AvcConvertCmd(srcName, destName, opt)

		if err != nil {
			t.Fatal(err)
		}

		cmdStr := cmd.String()
		cmdStr = strings.Replace(cmdStr, srcName, "SRC", 1)
		cmdStr = strings.Replace(cmdStr, destName, "DEST", 1)

		assert.Equal(t, "/usr/bin/ffmpeg -y -strict -2 -hwaccel vaapi -i SRC -c:a aac -vf scale='if(gte(iw,ih), min(1500, iw), -2):if(gte(iw,ih), -2, min(1500, ih))',format=nv12,hwupload -c:v h264_vaapi -map 0:v:0 -map 0:a:0? -r 30 -b:v 50M -f mp4 -movflags +faststart DEST", cmdStr)

		// Requires Video4Linux transcoding device.
		// runCmd(t, encoder, srcName, destName, cmd)
	})
	t.Run("QSV", func(t *testing.T) {
		opt := encode.Options{
			Bin:      "",
			Encoder:  "h264_qsv",
			Size:     1500,
			Bitrate:  "50M",
			MapVideo: MapVideoDefault,
			MapAudio: MapAudioDefault,
		}
		r, _, err := AvcConvertCmd("VID123.mov", "VID123.mov.avc", opt)

		if err != nil {
			t.Fatal(err)
		}

		assert.Contains(t, r.String(), "/bin/ffmpeg -y -strict -2 -hwaccel qsv -hwaccel_output_format qsv -qsv_device /dev/dri/renderD128 -i VID123.mov -c:a aac -vf scale_qsv=w='if(gte(iw,ih), min(1500, iw), -1)':h='if(gte(iw,ih), -1, min(1500, ih))':format=nv12 -c:v h264_qsv -map 0:v:0 -map 0:a:0? -r 30 -b:v 50M -bitrate 50M -f mp4 -movflags +faststart VID123.mov.avc")
	})
	t.Run("Apple", func(t *testing.T) {
		opt := encode.Options{
			Bin:      "",
			Encoder:  "h264_videotoolbox",
			Size:     1500,
			Bitrate:  "50M",
			MapVideo: MapVideoDefault,
			MapAudio: MapAudioDefault,
		}
		r, _, err := AvcConvertCmd("VID123.mov", "VID123.mov.avc", opt)

		if err != nil {
			t.Fatal(err)
		}

		assert.Contains(t, r.String(), "bin/ffmpeg -y -strict -2 -i VID123.mov -c:v h264_videotoolbox -map 0:v:0 -map 0:a:0? -c:a aac -vf scale='if(gte(iw,ih), min(1500, iw), -2):if(gte(iw,ih), -2, min(1500, ih))',format=yuv420p -profile high -level 51 -r 30 -b:v 50M -f mp4 -movflags +faststart VID123.mov.avc")
	})
	t.Run("Nvidia", func(t *testing.T) {
		opt := encode.Options{
			Bin:      "",
			Encoder:  "h264_nvenc",
			Size:     1500,
			Bitrate:  "50M",
			MapVideo: MapVideoDefault,
			MapAudio: MapAudioDefault,
		}
		r, _, err := AvcConvertCmd("VID123.mov", "VID123.mov.avc", opt)

		if err != nil {
			t.Fatal(err)
		}

		assert.Contains(t, r.String(), "bin/ffmpeg -y -strict -2 -hwaccel auto -i VID123.mov -pix_fmt yuv420p -c:v h264_nvenc -map 0:v:0 -map 0:a:0? -c:a aac -preset 15 -pixel_format yuv420p -gpu any -vf scale='if(gte(iw,ih), min(1500, iw), -2):if(gte(iw,ih), -2, min(1500, ih))',format=yuv420p -rc:v constqp -cq 0 -tune 2 -r 30 -b:v 50M -profile:v 1 -level:v auto -coder:v 1 -f mp4 -movflags +faststart VID123.mov.avc")
	})
	t.Run("Video4Linux", func(t *testing.T) {
		opt := encode.Options{
			Bin:      "",
			Encoder:  "h264_v4l2m2m",
			Size:     1500,
			Bitrate:  "50M",
			MapVideo: MapVideoDefault,
			MapAudio: MapAudioDefault,
		}
		r, _, err := AvcConvertCmd("VID123.mov", "VID123.mov.avc", opt)

		if err != nil {
			t.Fatal(err)
		}

		assert.Contains(t, r.String(), "bin/ffmpeg -y -strict -2 -i VID123.mov -c:v h264_v4l2m2m -map 0:v:0 -map 0:a:0? -c:a aac -vf scale='if(gte(iw,ih), min(1500, iw), -2):if(gte(iw,ih), -2, min(1500, ih))',format=yuv420p -num_output_buffers 72 -num_capture_buffers 64 -max_muxing_queue_size 1024 -crf 23 -r 30 -b:v 50M -f mp4 -movflags +faststart VID123.mov.avc")
	})
}
