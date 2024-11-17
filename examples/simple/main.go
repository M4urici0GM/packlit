// Copyright 2024 Mauricio Barbosa. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

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
	packager := packlit.NewBuilder().
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

	packagerExecutor := packlit.NewExecutor(packager)
	if err := packagerExecutor.Run(); err != nil {
		log.Fatalf("error when running shaka-packager: %v", err)
	}
}