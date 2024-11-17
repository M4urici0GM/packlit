package packlit

type ShakaRunnerBuilder struct {
	runner *ShakaRunner
}

func NewBuilder() *ShakaRunnerBuilder {
	return &ShakaRunnerBuilder{
		runner: NewRunner(""),
	}
}

func (b *ShakaRunnerBuilder) WithFlag(flag *ShakaOptions) *ShakaRunnerBuilder {
	b.runner.Flags = flag
	return b
}

func (b *ShakaRunnerBuilder) WithStream(streamOption *StreamOptions) *ShakaRunnerBuilder {
	b.runner.StreamOptions = append(b.runner.StreamOptions, streamOption)

	return b
}

func (b *ShakaRunnerBuilder) Build() *ShakaRunner {
	return b.runner
}

func (s *ShakaRunner) Args() ([]string, string, error) {
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
