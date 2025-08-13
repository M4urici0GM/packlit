package packlit

// WithTimeShiftBufferDepthFlag Adds flag '--time_shift_buffer_depth <duration>'
func WithTimeShiftBufferDepthFlag(val float32) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(TimeShiftBufferDepthFlag(val))
	}
}

// WithPreservedSegmentsOutsideLiveWindowFlag Adds flag '--preserved_segments_outside_live_window <count>'
func WithPreservedSegmentsOutsideLiveWindowFlag(val int) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(PreservedSegmentsOutsideLiveWindowFlag(val))
	}
}

// WithDefaultLanguageFlag Adds flag '--default_language <language>'
func WithDefaultLanguageFlag(val string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(DefaultLanguageFlag(val))
	}
}

// WithDefaultTextLanguageFlag Adds flag '--default_text_language <language>'
func WithDefaultTextLanguageFlag(val string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(DefaultTextLanguageFlag(val))
	}
}

// WithForceCLIndexFlag Adds flag '--force_cl_index'
func WithForceCLIndexFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(ForceCLIndexFlag{})
	}
}
