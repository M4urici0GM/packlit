package gaka

import "fmt"

type StreamDescriptorFlag interface {
	fmt.Stringer

	// Validate given flag.
	Validate() error
}

type StreamDescriptor struct {
	flags []StreamDescriptorFlag
}

func (d *StreamDescriptor) Add(flag StreamDescriptorFlag) {
	d.flags = append(d.flags, flag)
}

func WithInput(input string) func(*StreamDescriptor) {
	return func(sd *StreamDescriptor) {
        sd.Add(InputSelector(input))
	}
}

func WithStream(stream string) func(*StreamDescriptor) {
	return func(sd *StreamDescriptor) {
		sd.Add(StreamSelector(stream))
	}
}

// Adds flag ouput to the stream descriptor.
// Example:
// gaka.WithOutput('file.mp4') will produce a stream descriptor '...output=file.mp4' 
func WithOutput(output string) func(*StreamDescriptor) {
	return func(sd *StreamDescriptor) {
        sd.Add(OutputSelector(output))
	}
}

func NewStreamDescriptor(...func(*StreamDescriptor)) *StreamDescriptor {
	return &StreamDescriptor{
		flags: make([]StreamDescriptorFlag, 0),
	}
}

func Run(...*StreamDescriptor) error {
	panic("implement me.")
}
