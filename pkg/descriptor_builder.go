package packlit

type StreamDescriptorBuilder struct {
	options []StreamDescriptorFn
}

func NewStreamBuilder() *StreamDescriptorBuilder {
	return &StreamDescriptorBuilder{
		options: make([]StreamDescriptorFn, 0),
	}
}

func (b *StreamDescriptorBuilder) WithInput(input string) *StreamDescriptorBuilder {
	return b.AddOption(WithInput(input))
}

func (b *StreamDescriptorBuilder) WithStream(stream StreamType) *StreamDescriptorBuilder {
	return b.AddOption(WithStream(stream))
}

func (b *StreamDescriptorBuilder) WithOutput(output string) *StreamDescriptorBuilder {
	return b.AddOption(WithOutput(output))
}

func (b *StreamDescriptorBuilder) AddOption(opt StreamDescriptorFn) *StreamDescriptorBuilder {
	b.options = append(b.options, opt)
	return b
}

func (builder *StreamDescriptorBuilder) Build() *StreamOptions {
	opts := &StreamOptions{}
	for _, fn := range builder.options {
		fn(opts)
	}

	return opts
}
