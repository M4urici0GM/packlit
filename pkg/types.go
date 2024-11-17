package packlit 

import "fmt"

type StringerValidate interface {
	fmt.Stringer

	Validate() error
}

type StreamOpt = func(*StreamOptions)
type ShakaOpt = func(*ShakaOptions)

type StreamOption interface {
	StringerValidate
}

type ShakaFlag interface {
	StringerValidate
}

type ShakaRunner struct {
    // Location of the binary executable of Shaka-Packager.
    // Ex: /bin/packager
	Binary        string
    
    // Flags to be used when running Shaka-Packager
	Flags         *ShakaOptions

    // Stream descriptors.
	StreamOptions []*StreamOptions
}

type ShakaOptions struct {
	Flags []ShakaFlag
}

func (s *ShakaOptions) Add(flag ShakaFlag) {
	s.Flags = append(s.Flags, flag)
}

type StreamOptions struct {
	Options []StreamOption
}
