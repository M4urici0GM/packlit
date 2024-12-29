// Copyright 2024 Mauricio Barbosa. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

package packlit

import (
	"fmt"
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
type ShakaParser interface {
	Validate() error
	Parse() string
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
type ShakaPackager struct {
	// Location of the binary executable of Shaka-Packager.
	// Ex: /bin/packager
	Binary string

	// Flags to be used when running Shaka-Packager
	Flags *ShakaFlags

	// Stream descriptors.
	StreamOptions []*StreamOptions
}

func NewShakaPackager(binary string) *ShakaPackager {
	if binary == "" {
		binary = "/bin/packager"
	}

	return &ShakaPackager{
		Binary:        binary,
		Flags:         &ShakaFlags{},
		StreamOptions: make([]*StreamOptions, 0),
	}
}

func (r *ShakaPackager) BuildAndValidate() ([]string, error) {
	args := []string{}
	validationErrors := []string{}

    args, argsError := buildStreamDescriptors(r.StreamOptions...)
    flags, flagError := buildFlags(r.Flags)

    if argsError != nil || flagError != nil {
		return nil, fmt.Errorf("validation errors: %s", validationErrors)
    }

    args = append(args, flags)


	return args, nil
}

func (r *ShakaPackager) PreviewCommand() (string, error) {
	args, err := r.BuildAndValidate()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", r.Binary, strings.Join(args, " ")), nil
}

// Wrapper around a slice of implementations of interface `OptionParser`
type ShakaFlags struct {
	Flags []ShakaParser
}

func NewFlags(opts ...ShakaFlagFn) *ShakaFlags {
	options := &ShakaFlags{Flags: make([]ShakaParser, 0)}
	for _, fn := range opts {
		fn(options)
	}

	return options
}

func (s *ShakaFlags) Add(flag ShakaParser) {
	s.Flags = append(s.Flags, flag)
}

// Wrapper around a slice of implementations of interface `OptionParser`
// Both this and `ShakaFlags` type are for isolating and handling both concerns independently
type StreamOptions struct {
	Options []ShakaParser
}

func NewStreamDescriptor(opts ...StreamDescriptorFn) *StreamOptions {
	streamOptions := &StreamOptions{Options: make([]ShakaParser, 0)}
	for _, fn := range opts {
		fn(streamOptions)
	}

	return streamOptions
}

func (s *StreamOptions) Add(option ShakaParser) {
	s.Options = append(s.Options, option)
}