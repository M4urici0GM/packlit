// Copyright 2024 Mauricio Barbosa. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

package packlit

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	_ ShakaParser = (*SegmentDuration)(nil)
	_ ShakaParser = (*SegmentTemplate)(nil)
	_ ShakaParser = (*InitSegment)(nil)
)

// Build the list of StreamOptions into a slice of strings.
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

// builds the given stream descriptor into its string version.
// Se its tests for more information
func buildStreamDescriptor(descriptor *StreamOptions) (string, error) {
	flags := make([]string, 0)
	for _, f := range descriptor.Options {
		if err := f.Validate(); err != nil {
			return "", err
		}

		flags = append(flags, f.Parse())
	}

	return strings.Join(flags, ","), nil
}

// InputSelector input/source media “file” path, which can be regular files, pipes, udp streams.
// See https://shaka-project.github.io/shaka-packager/html/options/udp_file_options.html on additional options for UDP files.
type InputSelector string

func (i InputSelector) Parse() string {
	return fmt.Sprintf("input=%s", string(i))
}

func (i InputSelector) Validate() error {
	if string(i) == "" {
		return fmt.Errorf("input empty is not allowed.")
	}

	return nil
}

// StreamSelector Required field with value ‘audio’, ‘video’, ‘text’ or stream number (zero based).
type StreamSelector string

func (s StreamSelector) Parse() string {
	return fmt.Sprintf("stream_selector=%s", string(s))
}

func (s StreamSelector) Validate() error {
	switch string(s) {
	case "video", "audio", "text":
		return nil
	default:
		if string(s) == "" {
			return fmt.Errorf("empty stream selector is not allowed.")
		}

		result, ok := strconv.ParseInt(string(s), 10, 32)
		if ok != nil {
			return fmt.Errorf("invalid stream selector: %s, %v", string(s), ok)
		}

		if result < 0 {
			return fmt.Errorf("invalid zero based stream selector: %d", result)
		}

		return nil
	}
}

// Required output file path (single file).
type OutputSelector string

func (o OutputSelector) Parse() string {
	return fmt.Sprintf("output=%s", string(o))
}

func (o OutputSelector) Validate() error {
	return nil
}

// MAYBE IT SHOULD NOT BE HERE.

type SegmentDuration string

func (s SegmentDuration) Validate() error {
	return nil
}

func (s SegmentDuration) Parse() string {
	return fmt.Sprintf("--segment_duration=%s", s)
}

type SegmentTemplate string

func (s SegmentTemplate) Validate() error {
	return nil
}

func (s SegmentTemplate) Parse() string {
	return fmt.Sprintf("segment_template=%s", string(s))
}

type InitSegment string

func (i InitSegment) Validate() error {
	// Maybe check if the folder exists?





	return nil
}

func (i InitSegment) Parse() string {
	return fmt.Sprintf("init_segment=%s", string(i))
}
