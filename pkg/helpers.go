package gaka

func WithInput(input string) func(*StreamOptions) {
	return func(sd *StreamOptions) {
		sd.Add(InputSelector(input))
	}
}

func WithStream(stream string) func(*StreamOptions) {
	return func(sd *StreamOptions) {
		sd.Add(StreamSelector(stream))
	}
}

// Adds flag ouput to the stream descriptor.
// Example:
// gaka.WithOutput('file.mp4') will produce a stream descriptor '...output=file.mp4'
func WithOutput(output string) func(*StreamOptions) {
	return func(sd *StreamOptions) {
		sd.Add(OutputSelector(output))
	}
}
