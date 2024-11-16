package main

import (
	"log"

	gaka "github.com/m4urici0gm/gaka/pkg"
)

func buildDescriptors() []*gaka.StreamOptions {
	return []*gaka.StreamOptions{
		{
			Options: []gaka.StreamOption{
				gaka.InputSelector("file.mp4"),
				gaka.StreamSelector("audio"),
				gaka.OutputSelector("audio.mp4"),
			},
		},
		{
			Options: []gaka.StreamOption{
				gaka.InputSelector("file.mp4"),
				gaka.StreamSelector("video"),
				gaka.OutputSelector("video.mp4"),
			},
		},
	}
}

func buildFlags() *gaka.ShakaOptions {
	return &gaka.ShakaOptions{
		Flags: []gaka.ShakaFlag{
			gaka.GenerateStaticLiveMpd{},
		},
	}
}

func main() {
	descriptors := buildDescriptors()
	flags := buildFlags()

	runner := gaka.ShakaRunner{
		StreamOptions: descriptors,
		Flags:         flags,
	}

	if err := runner.Run(); err != nil {
		log.Fatalf("error when trying to run shaka-packager: %v", err)
	}
}
