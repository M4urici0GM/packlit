package packlit

// WithEnableRawKeyEncryptionFlag Adds flag '--enable_raw_key_encryption'
func WithEnableRawKeyEncryptionFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(EnableRawKeyEncryptionFlag{})
	}
}

// WithEnableRawKeyDecryptionFlag Adds flag '--enable_raw_key_decryption'
func WithEnableRawKeyDecryptionFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(EnableRawKeyDecryptionFlag{})
	}
}

// WithKeysFlag Adds flag '--keys=<value>'
func WithKeysFlag(value []string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(KeysFlag(value))
	}
}

// WithIvFlag Adds flag '--iv=<value>'
func WithIvFlag(value string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(IvFlag(value))
	}
}

// WithPsshFlag Adds flag '--pssh=<value>'
func WithPsshFlag(value string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(PsshFlag(value))
	}
}
