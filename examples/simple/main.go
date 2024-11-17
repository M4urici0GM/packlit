package main

import (
	"log"

	"github.com/m4urici0gm/packlit/pkg"
)

func main() {
	// Equivalent to
	// $ packager in=some_content.mp4,stream=video,out=video.mp4 \
	//            in=some_content.mp4,stream=audio,out=audio.mp4 \
	//      --mpd_output=file.mpd

	opts := packlit.NewFlags(packlit.WithMpdOutput("file.mpd"))
	runner := packlit.NewBuilder().
		WithStream(
			packlit.NewStreamBuilder().
				WithInput("some_content.mp4").
				WithStream("video").
				WithOutput("video.mp4").
				Build(),
		).
		WithStream(
			packlit.NewStreamBuilder().
				WithInput("some_content.mp4").
				WithStream("audio").
				WithOutput("audio.mp4").
				Build(),
		).
		WithFlag(opts).
		Build()

	if err := runner.Run(); err != nil {
		log.Fatalf("error running shaka-packager: %v", err)
	}
}
