package packlit

import (
	"fmt"
	"strings"
)

var (
	_ ShakaParser = (*OutputMediaInfoFlag)(nil)
	_ ShakaParser = (*MpdOutputFlag)(nil)
	_ ShakaParser = (*BaseUrlsFlag)(nil)
	_ ShakaParser = (*MinBufferTimeFlag)(nil)
	_ ShakaParser = (*MinimumUpdatePeriodFlag)(nil)
	_ ShakaParser = (*SuggestedPresentationDelayFlag)(nil)
	_ ShakaParser = (*UtCTimingsFlag)(nil)
	_ ShakaParser = (*GenerateDashIfIopCompliantMpdFlag)(nil)
	_ ShakaParser = (*AllowApproximateSegmentTimelineFlag)(nil)
	_ ShakaParser = (*AllowCodecSwitchingFlag)(nil)
	_ ShakaParser = (*IncludeMsprProForPlayReadyFlag)(nil)
	_ ShakaParser = (*DashForceSegmentListFlag)(nil)
	_ ShakaParser = (*LowLatencyDashModeFlag)(nil)
)

// OutputMediaInfoFlag represents the flag to output media information.
type OutputMediaInfoFlag struct{}

// Parse returns the string representation of OutputMediaInfo for use in the command line flags.
func (o OutputMediaInfoFlag) Parse() string {
	return "--output_media_info"
}

// Validate checks if the OutputMediaInfoFlag value is valid.
func (o OutputMediaInfoFlag) Validate() error {
	// No specific validation logic for this boolean flag.
	return nil
}

// BaseUrlsFlag represents the comma-separated BaseURLs for the MPD.
type BaseUrlsFlag []string

// Parse returns the string representation of BaseUrls for use in the command line flags.
func (b BaseUrlsFlag) Parse() string {
	return fmt.Sprintf("--base_urls=%s", strings.Join(b, ","))
}

// Validate checks if the BaseUrls value is valid.
func (b BaseUrlsFlag) Validate() error {
	return nil
}

// MinBufferTimeFlag represents the minimum buffer time for MPD.
type MinBufferTimeFlag float32

// Parse returns the string representation of MinBufferTime for use in the command line flags.
func (m MinBufferTimeFlag) Parse() string {
	return fmt.Sprintf("--min_buffer_time=%v", m)
}

// Validate checks if the MinBufferTime value is valid.
func (m MinBufferTimeFlag) Validate() error {
	return nil
}

// MinimumUpdatePeriodFlag represents the minimum update period for the MPD.
type MinimumUpdatePeriodFlag float32

// Parse returns the string representation of MinimumUpdatePeriod for use in the command line flags.
func (m MinimumUpdatePeriodFlag) Parse() string {
	return fmt.Sprintf("--minimum_update_period=%v", m)
}

// Validate checks if the MinimumUpdatePeriod value is valid.
func (m MinimumUpdatePeriodFlag) Validate() error {
	return nil
}

// SuggestedPresentationDelayFlag represents the suggested presentation delay for the MPD.
type SuggestedPresentationDelayFlag float32

// Parse returns the string representation of SuggestedPresentationDelay for use in the command line flags.
func (s SuggestedPresentationDelayFlag) Parse() string {
	return fmt.Sprintf("--suggested_presentation_delay=%v", s)
}

// Validate checks if the SuggestedPresentationDelay value is valid.
func (s SuggestedPresentationDelayFlag) Validate() error {
	// Placeholder validation logic. It could be any required check.
	return nil
}

// UtCTimingsFlag represents the UTC Timing for the MPD.
type UtCTimingsFlag []string

// Parse returns the string representation of UtCTimings for use in the command line flags.
func (u UtCTimingsFlag) Parse() string {
	return fmt.Sprintf("--utc_timings=%s", strings.Join(u, ","))
}

// Validate checks if the UtCTimings value is valid.
func (u UtCTimingsFlag) Validate() error {
	// You can add specific validation for the UTC Timing here.
	return nil
}

// GenerateDashIfIopCompliantMpdFlag represents the flag to generate DASH-IF IOP compliant MPD.
type GenerateDashIfIopCompliantMpdFlag struct{}

// Parse returns the string representation of GenerateDashIfIopCompliantMpd for use in the command line flags.
func (g GenerateDashIfIopCompliantMpdFlag) Parse() string {
	return "--generate_dash_if_iop_compliant_mpd"
}

// Validate checks if the GenerateDashIfIopCompliantMpdFlag value is valid.
func (g GenerateDashIfIopCompliantMpdFlag) Validate() error {
	// No specific validation logic for this boolean flag.
	return nil
}

// AllowApproximateSegmentTimelineFlag represents the flag to allow approximate segment timeline for live profiles.
type AllowApproximateSegmentTimelineFlag struct{}

// Parse returns the string representation of AllowApproximateSegmentTimeline for use in the command line flags.
func (a AllowApproximateSegmentTimelineFlag) Parse() string {
	return "--allow_approximate_segment_timeline"
}

// Validate checks if the AllowApproximateSegmentTimelineFlag value is valid.
func (a AllowApproximateSegmentTimelineFlag) Validate() error {
	// No specific validation logic for this boolean flag.
	return nil
}

// AllowCodecSwitchingFlag represents the flag to allow adaptive switching between different codecs.
type AllowCodecSwitchingFlag struct{}

// Parse returns the string representation of AllowCodecSwitching for use in the command line flags.
func (a AllowCodecSwitchingFlag) Parse() string {
	return "--allow_codec_switching"
}

// Validate checks if the AllowCodecSwitchingFlag value is valid.
func (a AllowCodecSwitchingFlag) Validate() error {
	// No specific validation logic for this boolean flag.
	return nil
}

// IncludeMsprProForPlayReadyFlag represents the flag to include mspr:pro for PlayReady content protection.
type IncludeMsprProForPlayReadyFlag struct{}

// Parse returns the string representation of IncludeMsprProForPlayReady for use in the command line flags.
func (i IncludeMsprProForPlayReadyFlag) Parse() string {
	return "--include_mspr_pro_for_playready"
}

// Validate checks if the IncludeMsprProForPlayReadyFlag value is valid.
func (i IncludeMsprProForPlayReadyFlag) Validate() error {
	// No specific validation logic for this boolean flag.
	return nil
}

// DashForceSegmentListFlag represents the flag to force SegmentList in DASH MPD.
type DashForceSegmentListFlag struct{}

// Parse returns the string representation of DashForceSegmentList for use in the command line flags.
func (d DashForceSegmentListFlag) Parse() string {
	return "--dash_force_segment_list"
}

// Validate checks if the DashForceSegmentListFlag value is valid.
func (d DashForceSegmentListFlag) Validate() error {
	// No specific validation logic for this boolean flag.
	return nil
}

// LowLatencyDashModeFlag represents the flag to enable low latency DASH mode.
type LowLatencyDashModeFlag bool

// Parse returns the string representation of LowLatencyDashMode for use in the command line flags.
func (l LowLatencyDashModeFlag) Parse() string {
	return fmt.Sprintf("--low_latency_dash_mode=%v", l)
}

// Validate checks if the LowLatencyDashModeFlag value is valid.
func (l LowLatencyDashModeFlag) Validate() error {
	// No specific validation logic for this boolean flag.
	return nil
}
