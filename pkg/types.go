package gaka

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
	binary        string
	flags         *ShakaOptions
	streamOptions []*StreamOptions
}

type ShakaOptions struct {
	flags []ShakaFlag
}

func (s *ShakaOptions) Add(flag ShakaFlag) {
	s.flags = append(s.flags, flag)
}

type StreamOptions struct {
	options []StreamOption
}
