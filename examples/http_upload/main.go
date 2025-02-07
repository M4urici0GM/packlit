// Refer to the Shaka Packager HTTP upload tutorial for more details:
// https://shaka-project.github.io/shaka-packager/html/tutorials/http_upload.html
// This implementation follows the HTTP upload process, ensuring media segments
// and MPD files are uploaded correctly to the specified server.
package main

import (
	"fmt"
	"log"
	"path"

	packlit "github.com/m4urici0gm/packlit/pkg"
)

func buildFlags(uploadURL string) *packlit.ShakaFlags {
	return packlit.NewFlags(
		packlit.WithUserAgentFlag("Packlit-client/1.0"),
		packlit.WithCaFileFlag("/etc/ssl/certs/ca-certificates.crt"),
		packlit.WithClientCertFileFlag("/etc/ssl/certs/client-cert.pem"),
		packlit.WithClientCertPrivateKeyFileFlag("/etc/ssl/private/client-key.pem"),
		packlit.WithClientCertPrivateKeyPasswordFlag("supersecurepassword"),
		packlit.WithDisablePeerVerificationFlag(),
		packlit.WithIgnoreHttpOutputFailuresFlag(),
		packlit.WithStaticLiveMpd(),
		packlit.WithSegmentDuration("4"),
		packlit.WithMpdOutput(path.Join(uploadURL, "h264.mpd")),
	)
}

type stream struct {
	streamType packlit.StreamType
	input      string
}

func main() {
	uploadURL := "http://localhost:6767/vod"

	streams := []stream{
		{streamType: "audio", input: "audio.mp4"},
		{streamType: "video", input: "video.mp4"},
	}

	builder := packlit.NewBuilder()
	flags := buildFlags(uploadURL)

	for _, stream := range streams {
		initSegment := path.Join(uploadURL, fmt.Sprintf("%v-init.mp4", stream.streamType))
		segmentTemplate := path.Join(uploadURL, fmt.Sprintf("%v-$Time$.m4s", stream.streamType))

		builder.WithStream(
			packlit.NewStreamBuilder().
				WithInput(stream.input).
				WithStream(stream.streamType).
				AddOption(packlit.WithInitSegment(initSegment)).
				AddOption(packlit.WithSegmentTemplate(segmentTemplate)).
				Build(),
		)
	}

	packager := builder.WithFlag(flags).Build()

	cmd, err := packager.PreviewCommand()
	if err != nil {
		log.Fatalf("error: when previewing command %v", err)
	}

	// Print command
	log.Println(cmd)

	packagerExecutor := packlit.NewExecutor(packager)
	if err := packagerExecutor.Run(); err != nil {
		log.Fatalf("error when running shaka-packager: %v", err)
	}
}
