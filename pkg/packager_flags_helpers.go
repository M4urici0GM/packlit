package packlit

// WithLicensesFlag Adds flag '--licenses'
func WithLicensesFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(LicensesFlag{})
	}
}

// WithQuietFlag Adds flag '--quiet'
func WithQuietFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(QuietFlag{})
	}
}

// WithUseFakeClockForMuxerFlag Adds flag '--use_fake_clock_for_muxer'
func WithUseFakeClockForMuxerFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(UseFakeClockForMuxerFlag{})
	}
}

// WithTestPackagerVersionFlag Adds flag '--test_packager_version'
func WithTestPackagerVersionFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(TestPackagerVersionFlag{})
	}
}

// WithSingleThreadedFlag Adds flag '--single_threaded'
func WithSingleThreadedFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(SingleThreadedFlag{})
	}
}

// WithVModuleFlag Adds flag '--vmodule=<value>'
func WithVModuleFlag(value string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(VModuleFlag(value))
	}
}

// WithVersionFlag Adds flag '--version'
func WithVersionFlag() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add((VersionFlag{}))
	}
}
