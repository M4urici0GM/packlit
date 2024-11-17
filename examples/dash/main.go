package main

import (
	"log"

	packlit "github.com/m4urici0gm/packlit/pkg"
)

func main() {
	opts := packlit.NewFlags(packlit.WithMpdOutput("h264.mpd"))
	runner := packlit.NewBuilder().
		WithStream(
			packlit.NewStreamBuilder().
				WithInput("h264_baseline_360p_600.mp4").
				WithStream(packlit.StreamTypeAudio).
				WithOutput("audio.mp4").
				Build(),
		).
		WithStream(
			packlit.NewStreamBuilder().
				WithInput("h264_main_480p_1000.mp4").
				WithStream(packlit.StreamTypeVideo).
				WithOutput("h264_480p.mp4").
				Build(),
		).
		WithStream(
			packlit.NewStreamBuilder().
				WithInput("h264_main_720p_3000.mp4").
				WithStream(packlit.StreamTypeVideo).
				WithOutput("h264_720p.mp4").
				Build(),
		).
		WithStream(
			packlit.NewStreamBuilder().
				WithInput("h264_high_1080p_6000.mp4").
				WithStream(packlit.StreamTypeVideo).
				WithOutput("h264_1080p.mp4").
				Build(),
		).
		WithFlag(opts).
		Build()

	if err := runner.Run(); err != nil {
		log.Fatalf("error when running shaka-packager: %v", err)
	}
}
