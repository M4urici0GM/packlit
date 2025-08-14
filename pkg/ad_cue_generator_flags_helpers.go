package packlit

// WithAdCuesFlag Adds flag '--ad_cues <start_time[,duration]>[;<start_time[,duration]>]...'
func WithAdCuesFlag(val string) ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(AdCuesFlag(val))
	}
}
