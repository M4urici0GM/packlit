package main

import (
	"log"

	packlit "github.com/m4urici0gm/packlit/pkg"
)

func buildDescriptors() []*packlit.StreamOptions {
	return []*packlit.StreamOptions{
		{
			Options: []packlit.OptionParser{
				packlit.InputSelector("file.mp4"),
				packlit.StreamSelector("audio"),
				packlit.OutputSelector("audio.mp4"),
			},
		},
		{
			Options: []packlit.OptionParser{
				packlit.InputSelector("file.mp4"),
				packlit.StreamSelector("video"),
				packlit.OutputSelector("video.mp4"),
			},
		},
	}
}

func buildFlags() *packlit.ShakaFlags {
	return &packlit.ShakaFlags{
		Flags: []packlit.OptionParser{
			packlit.GenerateStaticLiveMpd{},
		},
	}
}

func main() {
	descriptors := buildDescriptors()
	flags := buildFlags()

	runner := packlit.ShakaRunner{
		StreamOptions: descriptors,
		Flags:         flags,
	}

	if err := runner.Run(); err != nil {
		log.Fatalf("error when trying to run shaka-packager: %v", err)
	}
}
