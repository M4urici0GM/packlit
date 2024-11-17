package packlit

import "fmt"

type OptionParser interface {
	fmt.Stringer

	Validate() error
}

type StreamDescriptorFn = func(*StreamOptions)
type ShakaFlagFn = func(*ShakaFlags)

type ShakaRunner struct {
	// Location of the binary executable of Shaka-Packager.
	// Ex: /bin/packager
	Binary string

	// Flags to be used when running Shaka-Packager
	Flags *ShakaFlags

	// Stream descriptors.
	StreamOptions []*StreamOptions
}

type ShakaFlags struct {
	Flags []OptionParser
}

type StreamOptions struct {
	Options []OptionParser
}

func (s *ShakaFlags) Add(flag OptionParser) {
	s.Flags = append(s.Flags, flag)
}
