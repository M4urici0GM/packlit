package packlit

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type StreamType string

const (
	StreamTypeVideo StreamType = "video"
	StreamTypeAudio StreamType = "audio"
	StreamTypeText  StreamType = "text"
)

// Wrapper around the fmt.Sprinter with Validate method
// Used for making sure that the provided value is correct.
type OptionParser interface {
	fmt.Stringer

	Validate() error
}

// Represents a function that can add a new option to a given StremaOption
// Ex:
//
//	func WithInput(input string) StreamDescriptorFn {
//	  return func(opts *StreamOptions) {
//	    opts.Add(InputStream(input))
//	  }
//	}
type StreamDescriptorFn = func(*StreamOptions)

// Represents a function that can add new flags to *ShakaFlags
// Ex:
//
//	func WithStaticLiveMpd() StreamDescriptorFn {
//	  return func(so *StreamOptions) {
//		   so.Add(GenerateStaticLiveMpd{})
//		 }
//	}
type ShakaFlagFn = func(*ShakaFlags)

// Responsible for calling Shaka-packager.
type ShakaRunner struct {
	// Location of the binary executable of Shaka-Packager.
	// Ex: /bin/packager
	Binary string

	// Flags to be used when running Shaka-Packager
	Flags *ShakaFlags

	// Stream descriptors.
	StreamOptions []*StreamOptions
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

func (r *ShakaRunner) BuildAndValidate() ([]string, error) {
	args := []string{}
	validationErrors := []string{}

	for i, stream := range r.StreamOptions {
		for j, opt := range stream.Options {
			if err := opt.Validate(); err != nil {
				validationErrors = append(validationErrors, fmt.Sprintf("stream[%d].option[%d]: %v", i, j, err))
				continue
			}

			args = append(args, opt.String())
		}
	}

	for i, flag := range r.Flags.Flags {
		if err := flag.Validate(); err != nil {
			validationErrors = append(validationErrors, fmt.Sprintf("flag[%d]: %v", i, err))
			continue
		}

		args = append(args, flag.String())
	}

	if len(validationErrors) > 0 {
		return nil, fmt.Errorf("validation errors: %s", validationErrors)
	}

	return args, nil
}

func (r *ShakaRunner) PreviewCommand() (string, error) {
	args, err := r.BuildAndValidate()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", r.Binary, strings.Join(args, " ")), nil
}

func (r *ShakaRunner) Run() error {
	args, err := r.BuildAndValidate()
	if err != nil {
		return err
	}

	cmd := exec.Command(r.Binary, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Wrapper around a slice of implementations of interface `OptionParser`
type ShakaFlags struct {
	Flags []OptionParser
}

func NewFlags(opts ...ShakaFlagFn) *ShakaFlags {
	options := &ShakaFlags{Flags: make([]OptionParser, 0)}
	for _, fn := range opts {
		fn(options)
	}

	return options
}

func (s *ShakaFlags) Add(flag OptionParser) {
	s.Flags = append(s.Flags, flag)
}

// Wrapper around a slice of implementations of interface `OptionParser`
// Both this and `ShakaFlags` type are for isolating and handling both concerns independently
type StreamOptions struct {
	Options []OptionParser
}

func NewStreamDescriptor(opts ...StreamDescriptorFn) *StreamOptions {
	streamOptions := &StreamOptions{Options: make([]OptionParser, 0)}
	for _, fn := range opts {
		fn(streamOptions)
	}

	return streamOptions
}

func (s *StreamOptions) Add(option OptionParser) {
	s.Options = append(s.Options, option)
}
