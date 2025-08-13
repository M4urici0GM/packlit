package packlit

import (
	"fmt"
)

var (
	_ ShakaParser = (*TimeShiftBufferDepthFlag)(nil)
	_ ShakaParser = (*PreservedSegmentsOutsideLiveWindowFlag)(nil)
	_ ShakaParser = (*DefaultLanguageFlag)(nil)
	_ ShakaParser = (*DefaultTextLanguageFlag)(nil)
	_ ShakaParser = (*ForceCLIndexFlag)(nil)
)

// TimeShiftBufferDepthFlag represents the guaranteed duration of the time shifting buffer for HLS LIVE playlists and DASH dynamic media presentations.
type TimeShiftBufferDepthFlag float32

// Parse returns the string representation of TimeShiftBufferDepth for use in the command line flags.
func (t TimeShiftBufferDepthFlag) Parse() string {
	return fmt.Sprintf("--time_shift_buffer_depth=%v", t)
}

// Validate checks if the value of TimeShiftBufferDepth is valid.
// The value should be a positive duration in seconds.
func (t TimeShiftBufferDepthFlag) Validate() error {
	return nil
}

// PreservedSegmentsOutsideLiveWindowFlag represents the number of segments to preserve outside the live window.
type PreservedSegmentsOutsideLiveWindowFlag int

// Parse returns the string representation of PreservedSegmentsOutsideLiveWindow for use in the command line flags.
func (p PreservedSegmentsOutsideLiveWindowFlag) Parse() string {
	return fmt.Sprintf("--preserved_segments_outside_live_window=%d", p)
}

// Validate checks if the value of PreservedSegmentsOutsideLiveWindow is valid.
// The value must be non-negative.
func (p PreservedSegmentsOutsideLiveWindowFlag) Validate() error {
	return nil
}

// DefaultLanguageFlag represents the default language for audio and text tracks.
type DefaultLanguageFlag string

// Parse returns the string representation of DefaultLanguage for use in the command line flags.
func (d DefaultLanguageFlag) Parse() string {
	return fmt.Sprintf("--default_language=%s", d)
}

// Validate checks if the value of DefaultLanguage is valid.
func (d DefaultLanguageFlag) Validate() error {
	// No specific validation for language strings
	return nil
}

// DefaultTextLanguageFlag represents the default language for text tracks.
type DefaultTextLanguageFlag string

// Parse returns the string representation of DefaultTextLanguage for use in the command line flags.
func (d DefaultTextLanguageFlag) Parse() string {
	return fmt.Sprintf("--default_text_language=%s", d)
}

// Validate checks if the value of DefaultTextLanguage is valid.
func (d DefaultTextLanguageFlag) Validate() error {
	// No specific validation for text language strings
	return nil
}

// ForceCLIndexFlag represents the flag to force the muxer to order streams in the order given on the command-line.
type ForceCLIndexFlag struct{}

// Parse returns the string representation of ForceCLIndex for use in the command line flags.
func (f ForceCLIndexFlag) Parse() string {
	return "--force_cl_index"
}

// Validate checks if the value of ForceCLIndex is valid.
func (f ForceCLIndexFlag) Validate() error {
	// No specific validation logic for this flag
	return nil
}
