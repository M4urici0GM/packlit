package packlit

// WithHLSMasterPlaylistOutputFlag Adds flag '--hls_master_playlist_output <path>'
func WithHLSMasterPlaylistOutputFlag(val string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(HLSMasterPlaylistOutputFlag(val))
	}
}

// WithHLSBaseURLFlag Adds flag '--hls_base_url <url>'
func WithHLSBaseURLFlag(val string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(HLSBaseURLFlag(val))
	}
}

// WithHLSKeyURIFlag Adds flag '--hls_key_uri <uri>'
func WithHLSKeyURIFlag(val string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(HLSKeyURIFlag(val))
	}
}

// WithHLSPlaylistTypeFlag Adds flag '--hls_playlist_type <VOD|EVENT|LIVE>'
func WithHLSPlaylistTypeFlag(val HLSPlaylistTypeFlag) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(val)
	}
}

// WithHLSMediaSequenceNumberFlag Adds flag '--hls_media_sequence_number <number>'
func WithHLSMediaSequenceNumberFlag(val int) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(HLSMediaSequenceNumberFlag(val))
	}
}

// WithHLSStartTimeOffsetFlag Adds flag '--hls_start_time_offset <offset>'
func WithHLSStartTimeOffsetFlag(val float32) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(HLSStartTimeOffsetFlag(val))
	}
}

// WithHLSCreateSessionKeysFlag Adds flag '--create_session_keys'
func WithHLSCreateSessionKeysFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(HLSCreateSessionKeysFlag{})
	}
}
