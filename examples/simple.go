package main

import (
	"log"

	gaka "github.com/m4urici0gm/gaka/pkg"
)

func main() {
    // Equivalent to 
    // $ packager in=some_content.mp4,stream=video,out=video.mp4 \
    //            in=some_content.mp4,stream=audio,out=audio.mp4

	descriptors := []*gaka.StreamDescriptor{
		gaka.NewStreamDescriptor(
			gaka.WithInput("some_content.mp4"),
			gaka.WithStream("video"),
			gaka.WithOutput("video.mp4"),
		),
		gaka.NewStreamDescriptor(
			gaka.WithInput("some_content.mp4"),
			gaka.WithStream("audio"),
			gaka.WithOutput("audio.mp4"),
		),
	}

	if err := gaka.Run(descriptors...); err != nil {
		log.Fatalf("error trying to run shaka-packager: %v", err)
	}
}
