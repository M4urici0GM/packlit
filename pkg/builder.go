// Copyright 2024 Mauricio Barbosa. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

package packlit

type ShakaRunnerBuilder struct {
	runner *ShakaPackager
}

func NewBuilder(binary ...string) *ShakaRunnerBuilder {
    binaryPath := ""
    if len(binary) > 0 && binary[0] != "" {
        binaryPath = binary[0]
    }

	return &ShakaRunnerBuilder{
		runner: NewShakaPackager(binaryPath),
	}
}

func (b *ShakaRunnerBuilder) WithFlag(flag *ShakaFlags) *ShakaRunnerBuilder {
	b.runner.Flags = flag
	return b
}

func (b *ShakaRunnerBuilder) WithStream(streamOption *StreamOptions) *ShakaRunnerBuilder {
	b.runner.StreamOptions = append(b.runner.StreamOptions, streamOption)

	return b
}

func (b *ShakaRunnerBuilder) Build() *ShakaPackager {
	return b.runner
}

func (s *ShakaPackager) Args() ([]string, string, error) {
	streamOptions, err := buildStreamDescriptors(s.StreamOptions...)
	if err != nil {
		return nil, "", err
	}

	flags, err := buildFlags(s.Flags)
	if err != nil {
		return nil, "", err
	}

	return streamOptions, flags, nil
}
