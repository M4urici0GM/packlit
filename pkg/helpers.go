// Copyright 2024 Mauricio Barbosa. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

package packlit

// WithInput Adds option 'input' to the stream descriptor.
func WithInput(input string) StreamDescriptorFn {
	return func(sd *StreamOptions) {
		sd.Add(InputSelector(input))
	}
}

// WithStream Adds option 'stream' to the stream descriptor.
func WithStream(stream StreamType) StreamDescriptorFn {
	return func(sd *StreamOptions) {
		sd.Add(StreamSelector(stream))
	}
}

// WithOutput Adds option 'ouput' to the stream descriptor.
func WithOutput(output string) func(*StreamOptions) {
	return func(sd *StreamOptions) {
		sd.Add(OutputSelector(output))
	}
}

// WithStaticLiveMpd Adds flag '--generate_static_live_mpd' to cmd
func WithStaticLiveMpd() ShakaFlagFn {
	return func(flags *ShakaFlags) {
		flags.Add(GenerateStaticLiveMpd{})
	}
}

// WithMpdOutput Adds flag '--mpd_output <name>'
func WithMpdOutput(output string) func(*ShakaFlags) {
	return func(so *ShakaFlags) {
		so.Add(MpdOutput(output))
	}
}

func WithApproximateSegmentTimeline() ShakaFlagFn {
	return func(flags *ShakaFlags) {
		flags.Add(AllowApproximateSegmentTimeline{})
	}
}

// WithDumpStream Adds flag '--dump_stream_info'
func WithDumpStream() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(DumpStreamInfo{})
	}
}

func WithSegmentDuration(seconds string) func(*ShakaFlags) {
	return func(so *ShakaFlags) {
		so.Add(SegmentDuration(seconds))
	}
}

func WithSegmentTemplate(val string) StreamDescriptorFn {
	return func(options *StreamOptions) {
		options.Add(SegmentTemplate(val))
	}
}

func WithInitSegment(val string) StreamDescriptorFn {
	return func(options *StreamOptions) {
		options.Add(InitSegment(val))
	}
}

func WithBaseUrls(baseUrls []string) ShakaFlagFn {
    return func(sf *ShakaFlags) {
        sf.Add(BaseUrls(baseUrls))
    }
}
