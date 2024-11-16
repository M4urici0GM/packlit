package gaka


// Adds option 'input' to the stream descriptor.
// Example:
// gaka.WithInput('file.mp4') will produce a stream descriptor '...input=file.mp4'
func WithInput(input string) func(*StreamOptions) {
	return func(sd *StreamOptions) {
		sd.Add(InputSelector(input))
	}
}


// Adds option 'stream' to the stream descriptor.
// Example:
// gaka.WithStream('file.mp4') will produce a stream descriptor '...stream=audio|video|text|0..'
func WithStream(stream string) func(*StreamOptions) {
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
func WithStaticLiveMpd() func(*StreamOptions) {
	return func(so *StreamOptions) {
		so.Add(GenerateStaticLiveMpd{})
	}
}
