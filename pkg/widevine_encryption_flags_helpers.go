package packlit

// WithEnableWidevineEncryptionFlag Adds flag '--enable_widevine_encryption'
func WithEnableWidevineEncryptionFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(EnableWidevineEncryptionFlag{})
	}
}

// WithEnableWidevineDecryptionFlag Adds flag '--enable_widevine_decryption'
func WithEnableWidevineDecryptionFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(EnableWidevineDecryptionFlag{})
	}
}

// WithKeyServerURLFlag Adds flag '--key_server_url=<value>'
func WithKeyServerURLFlag(value string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(KeyServerURLFlag(value))
	}
}

// WithContentIDFlag Adds flag '--content_id=<value>'
func WithContentIDFlag(value string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(ContentIDFlag(value))
	}
}

// WithPolicyFlag Adds flag '--policy=<value>'
func WithPolicyFlag(value string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(PolicyFlag(value))
	}
}

// WithMaxSDPixelsFlag Adds flag '--max_sd_pixels=<value>'
func WithMaxSDPixelsFlag(value int) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(MaxSDPixelslag(value))
	}
}

// WithMaxHDPixelsFlag Adds flag '--max_hd_pixels=<value>'
func WithMaxHDPixelsFlag(value int) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(MaxHDPixelslag(value))
	}
}

// WithMaxUHD1PixelsFlag Adds flag '--max_uhd1_pixels=<value>'
func WithMaxUHD1PixelsFlag(value int) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(MaxUHD1Pixelslag(value))
	}
}

// WithSignerFlag Adds flag '--signer=<value>'
func WithSignerFlag(value string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(SignerFlag(value))
	}
}

// WithAESSigningKeyFlag Adds flag '--aes_signing_key=<value>'
func WithAESSigningKeyFlag(value string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(AESSigningKeyFlag(value))
	}
}

// WithAESSigningIVFlag Adds flag '--aes_signing_iv=<value>'
func WithAESSigningIVFlag(value string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(AESSigningIVFlag(value))
	}
}

// WithRSASigningKeyPathFlag Adds flag '--rsa_signing_key_path=<value>'
func WithRSASigningKeyPathFlag(value string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(RSASigningKeyPathFlag(value))
	}
}

// WithCryptoPeriodDurationFlag Adds flag '--crypto_period_duration=<value>'
func WithCryptoPeriodDurationFlag(value int) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(CryptoPeriodDurationFlag(value))
	}
}

// WithGroupIDFlag Adds flag '--group_id=<value>'
func WithGroupIDFlag(value string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(GroupIDFlag(value))
	}
}

// WithEnableEntitlementLicenseFlag Adds flag '--enable_entitlement_license'
func WithEnableEntitlementLicenseFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(EnableEntitlementLicenseFlag{})
	}
}
