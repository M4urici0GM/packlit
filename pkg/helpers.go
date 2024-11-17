// Copyright 2024 Mauricio Barbosa. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

package packlit

// Adds option 'input' to the stream descriptor.
// Example:
// gaka.WithInput('file.mp4') will produce a stream descriptor '...input=file.mp4'
func WithInput(input string) StreamDescriptorFn {
	return func(sd *StreamOptions) {
		sd.Add(InputSelector(input))
	}
}

// Adds option 'stream' to the stream descriptor.
// Example:
// gaka.WithStream('file.mp4') will produce a stream descriptor '...stream=audio|video|text|0..'
func WithStream(stream StreamType) StreamDescriptorFn {
	return func(sd *StreamOptions) {
		sd.Add(StreamSelector(stream))
	}
}

// Adds option 'ouput' to the stream descriptor.
// Example:
// gaka.WithOutput('file.mp4') will produce a stream descriptor '...output=file.mp4'
func WithOutput(output string) func(*StreamOptions) {
	return func(sd *StreamOptions) {
		sd.Add(OutputSelector(output))
	}
}

// Adds flag '--generate_static_live_mpd' to cmd
func WithStaticLiveMpd() StreamDescriptorFn {
	return func(so *StreamOptions) {
		so.Add(GenerateStaticLiveMpd{})
	}
}

func WithMpdOutput(output string) func(*ShakaFlags) {
	return func(so *ShakaFlags) {
		so.Add(MpdOutput(output))
	}
}

func WithDumpStream() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(DumpStreamInfo{})
	}
}
