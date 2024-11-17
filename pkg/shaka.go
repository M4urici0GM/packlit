package packlit 

import "strings"

func NewStreamDescriptor(opts ...StreamDescriptorFn) *StreamOptions {
	streamOptions := &StreamOptions{Options: make([]OptionParser, 0)}
	for _, fn := range opts {
		fn(streamOptions)
	}

	return streamOptions
}

func NewFlags(opts ...ShakaFlagFn) *ShakaFlags {
	options := &ShakaFlags{Flags: make([]OptionParser, 0)}
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
		Binary:        binary,
		Flags:         &ShakaFlags{},
		StreamOptions: make([]*StreamOptions, 0),
	}
}

func buildFlags(flags *ShakaFlags) (string, error) {
	built := make([]string, 0)
	for _, flag := range flags.Flags {
		if err := flag.Validate(); err != nil {
			return "", err
		}

		built = append(built, flag.String())
	}

	return strings.Join(built, " "), nil
}

func buildStreamDescriptors(descriptors ...*StreamOptions) ([]string, error) {
	built := make([]string, 0)
	for _, desc := range descriptors {
		builtDesc, err := buildStreamDescriptor(desc)
		if err != nil {
			return nil, err
		}

		built = append(built, builtDesc)
	}

	return built, nil
}

func (r *ShakaRunner) Run() error {
	panic("implement me.")
}

func (d *StreamOptions) Add(flag OptionParser) {
	d.Options = append(d.Options, flag)
}
