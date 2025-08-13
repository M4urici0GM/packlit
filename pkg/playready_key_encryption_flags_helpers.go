package packlit

// WithPlayReadyEncryptionFlag Adds flag '--enable_playready_encryption'
func WithPlayReadyEncryptionFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(PlayReadyEncryptionFlag{})
	}
}

// WithPlayReadyServerURLFlag Adds flag '--playready_server_url=<value>'
func WithPlayReadyServerURLFlag(value string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(PlayReadyServerURLFlag(value))
	}
}

// WithProgramIdentifierFlag Adds flag '--program_identifier=<value>'
func WithProgramIdentifierFlag(value string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(ProgramIdentifierFlag(value))
	}
}
