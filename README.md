# Packlit
Packlit is a Golang library providing easy-to-use bindings for Shaka Packager,
enabling streamlined video and audio packaging workflows. With Packlit,
you can integrate powerful media packaging capabilities into your Go applications,
supporting formats like DASH and HLS with encryption and DRM.

## Warning
Not all features of Shaka-packer is yet implemented.
If you need a descriptor or a flag that is not implemented yet, you can implement your own, (Or open a PR, please)
Take a loook at `examples/custom_descriptor` to see how to implement a custom descriptor, and in `examples/custom_flag` to
see more information.
Feel free to implement flags/descriptor as you wish.

## Features
- Easy integration with Shaka Packager CLI
- Support for packaging DASH and HLS formats
- DRM support for Widevine, PlayReady, and FairPlay
- Simplified API for common use cases
- Highly customizable for advanced workflows

## Prerequisites
- Shaka Packager installed and available in your system's PATH
- Go version 1.18 or later

## Installation
To install Packlit, use:

```bash
$ go get github.com/yourusername/packlit  
```

## How to use

- Basic usage, described [here](https://shaka-project.github.io/shaka-packager/html/tutorials/basic_usage.html)
```bash
$ packager input=some_content.mp4 --dump_stream_info
```
```go
func main() {
	opts := packlit.NewShakaOptions(packlit.WithDumpStream())
	runner := packlit.NewBuilder().
		WithStream(packlit.NewStreamDescriptor(packlit.WithInput("some_content.mp4")))).
		WithFlag(opts).
		Build()

    if err := runner.Run(); err != nil {
        fmt.Fatalf("error when running shaka-packager: %v", err)
    }
}
```

- Dash usage, described [here](https://shaka-project.github.io/shaka-packager/html/tutorials/dash.html)
```go

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
```

You can achieve the same result as the above with
```go
func buildDescriptors() []*packlit.StreamOptions {
	return []*packlit.StreamOptions{
		{Options: []packlit.OptionParser{
				packlit.InputSelector("h264_baseline_360p_600.mp4"),
				packlit.StreamSelector("audio"),
				packlit.OutputSelector("audio.mp4"),
		}},
		{Options: []packlit.OptionParser{
				packlit.InputSelector("input_text.vtt"),
				packlit.StreamSelector("text"),
				packlit.OutputSelector("output_text.vtt"),
		}},
        ...
	}
}

func buildFlags() *packlit.ShakaFlags {
	return &packlit.ShakaFlags{
		Flags: []packlit.OptionParser{
			packlit.MpdOuput("h264.mpd"),
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
```

## Contributions
Contributions are welcome! If you'd like to contribute, please open an issue or submit a pull request.

## License
Packlit is licensed under the MIT License. See LICENSE for more details.

## Acknowledgements
Packlit is built around the excellent Shaka Packager project.

## License
```
This file is part of Packlit
                                                                     
Packlit is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation,
                                                                     
Packlit is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Lesser General Public License for more details.
                                                                     
You should have received a copy of the GNU Lesser General Public License
along with Packlit. If not, see <http://www.gnu.org/licenses/>.
```
