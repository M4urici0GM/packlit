package packlit

// WithClearLeadFlag Adds flag '--clear_lead <seconds>'
func WithClearLeadFlag(val float32) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(ClearLeadFlag(val))
	}
}

// WithSegmentDurationFlag Adds flag '--segment_duration <seconds>'
func WithSegmentDurationFlag(val float32) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(SegmentDurationFlag(val))
	}
}

// WithSegmentSapAlignedFlag Adds flag '--segment_sap_aligned'
func WithSegmentSapAlignedFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(SegmentSapAlignedFlag{})
	}
}

// WithFragmentDurationFlag Adds flag '--fragment_duration <seconds>'
func WithFragmentDurationFlag(val float32) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(FragmentDurationFlag(val))
	}
}

// WithFragmentSapAlignedFlag Adds flag '--fragment_sap_aligned'
func WithFragmentSapAlignedFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(FragmentSapAlignedFlag{})
	}
}

// WithGenerateSidxInMediaSegmentsFlag Adds flag '--generate_sidx_in_media_segments'
func WithGenerateSidxInMediaSegmentsFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(GenerateSidxInMediaSegmentsFlag{})
	}
}

// WithTempDirFlag Adds flag '--temp_dir <directory>'
func WithTempDirFlag(val string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(TempDirFlag(val))
	}
}

// WithMp4IncludePsshInStreamFlag Adds flag '--mp4_include_pssh_in_stream'
func WithMp4IncludePsshInStreamFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(Mp4IncludePsshInStreamFlag{})
	}
}

// WithTransportStreamTimestampOffsetMsFlag Adds flag '--transport_stream_timestamp_offset_ms <milliseconds>'
func WithTransportStreamTimestampOffsetMsFlag(val int) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(TransportStreamTimestampOffsetMsFlag(val))
	}
}

// WithDefaultTextZeroBiasMsFlag Adds flag '--default_text_zero_bias_ms <milliseconds>'
func WithDefaultTextZeroBiasMsFlag(val int) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(DefaultTextZeroBiasMsFlag(val))
	}
}

// WithStartSegmentNumberFlag Adds flag '--start_segment_number <number>'
func WithStartSegmentNumberFlag(val int) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(StartSegmentNumberFlag(val))
	}
}

// WithStartSegmentNumberFlag Adds flag '--use_dovi_supplemental_codecs'
func WithUseDoviSupplementalCodecsFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(UseDoviSupplementalCodecsFlag{})
	}
}
