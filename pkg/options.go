package gaka

import (
	"fmt"
	"strconv"
	"strings"
)

// builds the given stream descritpor into its string version.
func buildStreamDescriptor(descriptor *StreamDescriptor) (string, error) {
	flags := make([]string, 0)
	for _, f := range descriptor.flags {
		if err := f.Validate(); err != nil {
			return "", err
		}

		flags = append(flags, f.String())
	}

	return strings.Join(flags, ","), nil
}

// input/source media “file” path, which can be regular files, pipes, udp streams.
// See https://shaka-project.github.io/shaka-packager/html/options/udp_file_options.html on additional options for UDP files.
type InputSelector string

func (i InputSelector) String() string {
	return fmt.Sprintf("input=%s", string(i))
}

func (i InputSelector) Validate() error {
	if string(i) == "" {
		return fmt.Errorf("input empty is not allowed.")
	}

	return nil
}

// Required field with value ‘audio’, ‘video’, ‘text’ or stream number (zero based).
type StreamSelector string

func (s StreamSelector) String() string {
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

func (o OutputSelector) String() string {
	return fmt.Sprintf("output=%s", string(o))
}

func (o OutputSelector) Validate() error {
	return nil
}
