package gaka

import "strings"

func NewStreamDescriptor(opts ...StreamOpt) *StreamOptions {
	streamOptions := &StreamOptions{options: make([]StreamOption, 0)}
	for _, fn := range opts {
		fn(streamOptions)
	}

	return streamOptions
}

func NewShakaOptions(opts ...ShakaOpt) *ShakaOptions {
	options := &ShakaOptions{flags: make([]ShakaFlag, 0)}
	for _, fn := range opts {
		fn(options)
	}

	return options
}

func NewRunner(binary string) *ShakaRunner {
	if binary == "" {
		binary = "/bin/packager"
	}

	return &ShakaRunner{
		binary:        binary,
		flags:         &ShakaOptions{},
		streamOptions: make([]*StreamOptions, 0),
	}
}

func buildFlags(flags *ShakaOptions) (string, error) {
	built := make([]string, 0)
	for _, flag := range flags.flags {
		if err := flag.Validate(); err != nil {
			return "", err
		}

		built = append(built, flag.String())
	}

	return strings.Join(built, " "), nil
}

func buildStreamDescriptors(descriptors ...*StreamOptions) (string, error) {
	built := make([]string, 0)
	for _, desc := range descriptors {
		builtDesc, err := buildStreamDescriptor(desc)
		if err != nil {
			return "", err
		}

		built = append(built, builtDesc)
	}

	return strings.Join(built, " "), nil
}

func (r *ShakaRunner) Run() error {
	panic("implement me.")
}

func (d *StreamOptions) Add(flag StreamOption) {
	d.options = append(d.options, flag)
}
