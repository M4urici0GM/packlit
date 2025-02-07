package packlit

// WithUserAgentFlag Adds flag '--user_agent <value>'
func WithUserAgentFlag(val string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(UserAgentFlag(val))
	}
}

// WithCaFileFlag Adds flag '--ca_file <value>'
func WithCaFileFlag(val string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(CaFileFlag(val))
	}
}

// WithClientCertFileFlag Adds flag '--client_cert_file <value>'
func WithClientCertFileFlag(val string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(ClientCertFileFlag(val))
	}
}

// WithClientCertPrivateKeyFileFlag Adds flag '--client_cert_private_key_file <value>'
func WithClientCertPrivateKeyFileFlag(val string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(ClientCertPrivateKeyFileFlag(val))
	}
}

// WithClientCertPrivateKeyPasswordFlag Adds flag '--client_cert_private_key_password <value>'
func WithClientCertPrivateKeyPasswordFlag(val string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(ClientCertPrivateKeyPasswordFlag(val))
	}
}

// WithDisablePeerVerificationFlag Adds flag '--disable_peer_verification'
func WithDisablePeerVerificationFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(DisablePeerVerificationFlag{})
	}
}

// WithIgnoreHttpOutputFailuresFlag Adds flag '--ignore_http_output_failures'
func WithIgnoreHttpOutputFailuresFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(IgnoreHttpOutputFailuresFlag{})
	}
}

// WithIoCacheSizeFlag Adds flag '--io_cache_size <value>'
func WithIoCacheSizeFlag(val uint) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(IoCacheSizeFlag(val))
	}
}
