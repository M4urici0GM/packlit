package main

import (
	"log"

	gaka "github.com/m4urici0gm/gaka/pkg"
)

func main() {
	// Equivalent to
	// $ packager in=some_content.mp4,stream=video,out=video.mp4 \
	//            in=some_content.mp4,stream=audio,out=audio.mp4 \
    //      --mpd_output=file.mpd

	opts := gaka.NewShakaOptions(gaka.WithMpdOutput("file.mpd"))
	runner := gaka.NewBuilder().
		WithStream(gaka.NewStreamDescriptor(gaka.WithInput("some_content.mp4"), gaka.WithStream("video"), gaka.WithOutput("video.mp4"))).
		WithStream(gaka.NewStreamDescriptor(gaka.WithInput("some_content.mp4"), gaka.WithStream("audio"), gaka.WithOutput("audio.mp4"))).
		WithFlag(opts).
		Build()

	if err := runner.Run(); err != nil {
		log.Fatalf("error running shaka-packager: %v", err)
	}
}
