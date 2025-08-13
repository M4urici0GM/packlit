package packlit

// WithCryptByteBlockFlag Adds flag '--crypt_byte_block <value>'
func WithCryptByteBlockFlag(val int) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(CryptByteBlockFlag(val))
	}
}

// WithProtectionSchemeFlag Adds flag '--protection_scheme <cenc|cbc1|cens|cbcs>'
func WithProtectionSchemeFlag(val string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(ProtectionSchemeFlag(val))
	}
}

// WithSkipByteBlockFlag Adds flag '--skip_byte_block <value>'
func WithSkipByteBlockFlag(val int) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(SkipByteBlockFlag(val))
	}
}

// WithVP9SubsampleEncryptionFlag Adds flag '--vp9_subsample_encryption'
func WithVP9SubsampleEncryptionFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(VP9SubsampleEncryptionFlag{})
	}
}

// WithPlayReadyExtraHeaderDataFlag Adds flag '--playready_extra_header_data <XML>'
func WithPlayReadyExtraHeaderDataFlag(val string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(PlayReadyExtraHeaderDataFlag(val))
	}
}
