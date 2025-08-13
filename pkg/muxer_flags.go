package packlit

import (
	"fmt"
)

var (
	_ ShakaParser = (*ClearLeadFlag)(nil)
	_ ShakaParser = (*SegmentDurationFlag)(nil)
	_ ShakaParser = (*SegmentSapAlignedFlag)(nil)
	_ ShakaParser = (*FragmentDurationFlag)(nil)
	_ ShakaParser = (*FragmentSapAlignedFlag)(nil)
	_ ShakaParser = (*GenerateSidxInMediaSegmentsFlag)(nil)
	_ ShakaParser = (*TempDirFlag)(nil)
	_ ShakaParser = (*Mp4IncludePsshInStreamFlag)(nil)
	_ ShakaParser = (*TransportStreamTimestampOffsetMsFlag)(nil)
	_ ShakaParser = (*DefaultTextZeroBiasMsFlag)(nil)
	_ ShakaParser = (*StartSegmentNumberFlag)(nil)
	_ ShakaParser = (*UseDoviSupplementalCodecsFlag)(nil)
)

// ClearLeadFlag represents the clear lead time in seconds when encryption is enabled.
type ClearLeadFlag float32

// Parse returns the string representation of ClearLead for use in the command line flags.
func (c ClearLeadFlag) Parse() string {
	return fmt.Sprintf("--clear_lead=%v", c)
}

// Validate checks if the ClearLead value is valid.
func (c ClearLeadFlag) Validate() error {
	if c < 0 {
		return fmt.Errorf("clear_lead must be a positive value")
	}
	return nil
}

// SegmentDurationFlag represents the segment duration in seconds.
type SegmentDurationFlag float32

// Parse returns the string representation of SegmentDuration for use in the command line flags.
func (s SegmentDurationFlag) Parse() string {
	return fmt.Sprintf("--segment_duration=%v", s)
}

// Validate checks if the SegmentDuration value is valid.
func (s SegmentDurationFlag) Validate() error {
	if s <= 0 {
		return fmt.Errorf("segment_duration must be greater than zero")
	}
	return nil
}

// SegmentSapAlignedFlag represents the flag to force segments to begin with stream access points.
type SegmentSapAlignedFlag struct{}

// Parse returns the string representation of SegmentSapAligned for use in the command line flags.
func (s SegmentSapAlignedFlag) Parse() string {
	return "--segment_sap_aligned"
}

// Validate checks if the SegmentSapAligned value is valid.
func (s SegmentSapAlignedFlag) Validate() error {
	return nil
}

// FragmentDurationFlag represents the fragment duration in seconds.
type FragmentDurationFlag float32

// Parse returns the string representation of FragmentDuration for use in the command line flags.
func (f FragmentDurationFlag) Parse() string {
	return fmt.Sprintf("--fragment_duration=%v", f)
}

// Validate checks if the FragmentDuration value is valid.
func (f FragmentDurationFlag) Validate() error {
	if f < 0 {
		return fmt.Errorf("fragment_duration must not be negative")
	}
	return nil
}

// FragmentSapAlignedFlag represents the flag to force fragments to begin with stream access points.
type FragmentSapAlignedFlag struct{}

// Parse returns the string representation of FragmentSapAligned for use in the command line flags.
func (f FragmentSapAlignedFlag) Parse() string {
	return "--fragment_sap_aligned"
}

// Validate checks if the FragmentSapAligned value is valid.
func (f FragmentSapAlignedFlag) Validate() error {
	return nil
}

// GenerateSidxInMediaSegmentsFlag represents the flag to generate 'sidx' box in media segments.
type GenerateSidxInMediaSegmentsFlag struct{}

// Parse returns the string representation of GenerateSidxInMediaSegments for use in the command line flags.
func (g GenerateSidxInMediaSegmentsFlag) Parse() string {
	return "--generate_sidx_in_media_segments"
}

// Validate checks if the GenerateSidxInMediaSegments value is valid.
func (g GenerateSidxInMediaSegmentsFlag) Validate() error {
	return nil
}

// TempDirFlag represents the directory for temporary files.
type TempDirFlag string

// Parse returns the string representation of TempDir for use in the command line flags.
func (t TempDirFlag) Parse() string {
	return fmt.Sprintf("--temp_dir=%s", t)
}

// Validate checks if the TempDir value is valid.
func (t TempDirFlag) Validate() error {
	if t == "" {
		return fmt.Errorf("temp_dir must not be empty")
	}
	return nil
}

// Mp4IncludePsshInStreamFlag represents the flag to include PSSH in the encrypted stream (MP4 only).
type Mp4IncludePsshInStreamFlag struct{}

// Parse returns the string representation of Mp4IncludePsshInStream for use in the command line flags.
func (m Mp4IncludePsshInStreamFlag) Parse() string {
	return "--mp4_include_pssh_in_stream"
}

// Validate checks if the Mp4IncludePsshInStream value is valid.
func (m Mp4IncludePsshInStreamFlag) Validate() error {
	return nil
}

// TransportStreamTimestampOffsetMsFlag represents the offset for transport stream timestamps in milliseconds.
type TransportStreamTimestampOffsetMsFlag int

// Parse returns the string representation of TransportStreamTimestampOffsetMs for use in the command line flags.
func (m TransportStreamTimestampOffsetMsFlag) Parse() string {
	return fmt.Sprintf("--transport_stream_timestamp_offset_ms=%d", m)
}

// Validate checks if the TransportStreamTimestampOffsetMs value is valid.
func (m TransportStreamTimestampOffsetMsFlag) Validate() error {
	return nil
}

// DefaultTextZeroBiasMsFlag represents the bias threshold for text stream starting time.
type DefaultTextZeroBiasMsFlag int

// Parse returns the string representation of DefaultTextZeroBiasMs for use in the command line flags.
func (d DefaultTextZeroBiasMsFlag) Parse() string {
	return fmt.Sprintf("--default_text_zero_bias_ms=%d", d)
}

// Validate checks if the DefaultTextZeroBiasMs value is valid.
func (d DefaultTextZeroBiasMsFlag) Validate() error {
	return nil
}

// StartSegmentNumberFlag represents the start number in DASH SegmentTemplate and HLS segment name.
type StartSegmentNumberFlag int

// Parse returns the string representation of StartSegmentNumber for use in the command line flags.
func (s StartSegmentNumberFlag) Parse() string {
	return fmt.Sprintf("--start_segment_number=%d", s)
}

// Validate checks if the StartSegmentNumber value is valid.
func (s StartSegmentNumberFlag) Validate() error {
	return nil
}

// UseDoviSupplementalCodecsFlag represents the flag to signal DolbyVision using the modern supplemental codecs approach.
type UseDoviSupplementalCodecsFlag struct{}

// Parse returns the string representation of UseDoviSupplementalCodecs for use in the command line flags.
func (f UseDoviSupplementalCodecsFlag) Parse() string {
	return "--use_dovi_supplemental_codecs"
}

// Validate checks if the UseDoviSupplementalCodecs value is valid.
func (f UseDoviSupplementalCodecsFlag) Validate() error {
	return nil
}
