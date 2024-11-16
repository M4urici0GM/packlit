package gaka

type ShakaRunnerBuilder struct {
	runner *ShakaRunner
}

func NewBuilder() *ShakaRunnerBuilder {
	return &ShakaRunnerBuilder{
		runner: NewRunner(""),
	}
}

func (b *ShakaRunnerBuilder) WithFlag(flag *ShakaOptions) *ShakaRunnerBuilder {
	b.runner.flags = flag
	return b
}

func (b *ShakaRunnerBuilder) WithStream(streamOption *StreamOptions) *ShakaRunnerBuilder {
	b.runner.streamOptions = append(b.runner.streamOptions, streamOption)

	return b
}

func (b *ShakaRunnerBuilder) Build() *ShakaRunner {
	return b.runner
}
