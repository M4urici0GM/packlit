package packlit

// WithProtectionSystemsFlag Adds flag '--protection_systems=<value>'
func WithProtectionSystemsFlag(value []string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(ProtectionSystemsFlag(value))
	}
}
