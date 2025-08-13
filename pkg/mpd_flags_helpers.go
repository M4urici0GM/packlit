package packlit

// WithOutputMediaInfoFlag Adds flag '--output_media_info'
func WithOutputMediaInfoFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(OutputMediaInfoFlag{})
	}
}

// WithBaseUrlsFlag Adds flag '--base_urls <url1,url2,...>'
func WithBaseUrlsFlag(values []string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(BaseUrlsFlag(values))
	}
}

// WithMinBufferTimeFlag Adds flag '--min_buffer_time <duration>'
func WithMinBufferTimeFlag(val float32) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(MinBufferTimeFlag(val))
	}
}

// WithMinimumUpdatePeriodFlag Adds flag '--minimum_update_period <duration>'
func WithMinimumUpdatePeriodFlag(val float32) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(MinimumUpdatePeriodFlag(val))
	}
}

// WithSuggestedPresentationDelayFlag Adds flag '--suggested_presentation_delay <duration>'
func WithSuggestedPresentationDelayFlag(val float32) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(SuggestedPresentationDelayFlag(val))
	}
}

// WithUtCTimingsFlag Adds flag '--utc_timings <timing>'
func WithUtCTimingsFlag(val []string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(UtCTimingsFlag(val))
	}
}

// WithGenerateDashIfIopCompliantMpdFlag Adds flag '--generate_dash_if_iop_compliant_mpd'
func WithGenerateDashIfIopCompliantMpdFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(GenerateDashIfIopCompliantMpdFlag{})
	}
}

// WithAllowApproximateSegmentTimelineFlag Adds flag '--allow_approximate_segment_timeline'
func WithAllowApproximateSegmentTimelineFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(AllowApproximateSegmentTimelineFlag{})
	}
}

// WithAllowCodecSwitchingFlag Adds flag '--allow_codec_switching'
func WithAllowCodecSwitchingFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(AllowCodecSwitchingFlag{})
	}
}

// WithIncludeMsprProForPlayReadyFlag Adds flag '--include_mspr_pro_for_playready'
func WithIncludeMsprProForPlayReadyFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(IncludeMsprProForPlayReadyFlag{})
	}
}

// WithDashForceSegmentListFlag Adds flag '--dash_force_segment_list'
func WithDashForceSegmentListFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(DashForceSegmentListFlag{})
	}
}

// WithLowLatencyDashModeFlag Adds flag '--low_latency_dash_mode=boolean'
func WithLowLatencyDashModeFlag(val bool) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(LowLatencyDashModeFlag(val))
	}
}
