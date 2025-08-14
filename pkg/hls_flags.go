package packlit

import (
	"fmt"
)

var (
	_ ShakaParser = (*HLSMasterPlaylistOutputFlag)(nil)
	_ ShakaParser = (*HLSBaseURLFlag)(nil)
	_ ShakaParser = (*HLSKeyURIFlag)(nil)
	_ ShakaParser = (*HLSPlaylistTypeFlag)(nil)
	_ ShakaParser = (*HLSMediaSequenceNumberFlag)(nil)
	_ ShakaParser = (*HLSStartTimeOffsetFlag)(nil)
	_ ShakaParser = (*HLSCreateSessionKeysFlag)(nil)
)

// HLSMasterPlaylistOutputFlag represents the output path for the master playlist for HLS.
type HLSMasterPlaylistOutputFlag string

// Parse returns the string representation of HLSMasterPlaylistOutput for use in the command line flags.
func (h HLSMasterPlaylistOutputFlag) Parse() string {
	return fmt.Sprintf("--hls_master_playlist_output=%s", h)
}

// Validate checks if the value of HLSMasterPlaylistOutput is valid.
// This flag must be used to output HLS.
func (h HLSMasterPlaylistOutputFlag) Validate() error {
	if len(h) == 0 {
		return fmt.Errorf("hls_master_playlist_output is required")
	}
	return nil
}

// HLSBaseURLFlag represents the base URL for the Media Playlists and media files listed in the playlists.
type HLSBaseURLFlag string

// Parse returns the string representation of HLSBaseURL for use in the command line flags.
func (h HLSBaseURLFlag) Parse() string {
	return fmt.Sprintf("--hls_base_url=%s", h)
}

// Validate checks if the value of HLSBaseURL is valid.
// The base URL is the prefix for the media files.
func (h HLSBaseURLFlag) Validate() error {
	if len(h) == 0 {
		return fmt.Errorf("hls_base_url is required")
	}
	return nil
}

// HLSKeyURIFlag represents the key URI for the 'identity' and 'com.apple.streamingkeydelivery' key formats.
type HLSKeyURIFlag string

// Parse returns the string representation of HLSKeyURI for use in the command line flags.
func (h HLSKeyURIFlag) Parse() string {
	return fmt.Sprintf("--hls_key_uri=%s", h)
}

// Validate checks if the value of HLSKeyURI is valid.
// The key URI is required if the playlist is encrypted and uses the specified key formats.
func (h HLSKeyURIFlag) Validate() error {
	if len(h) == 0 {
		return fmt.Errorf("hls_key_uri is required for encrypted playlists with 'identity' and 'com.apple.streamingkeydelivery' formats")
	}
	return nil
}

// HLSPlaylistTypeFlag represents the type of the playlist: VOD, EVENT, or LIVE.
type HLSPlaylistTypeFlag string

// Parse returns the string representation of HLSPlaylistType for use in the command line flags.
func (h HLSPlaylistTypeFlag) Parse() string {
	return fmt.Sprintf("--hls_playlist_type=%s", h)
}

// Validate checks if the value of HLSPlaylistType is valid.
func (h HLSPlaylistTypeFlag) Validate() error {
	return nil
}

// HLSMediaSequenceNumberFlag represents the initial EXT-X-MEDIA-SEQUENCE value.
type HLSMediaSequenceNumberFlag int

// Parse returns the string representation of HLSMediaSequenceNumber for use in the command line flags.
func (h HLSMediaSequenceNumberFlag) Parse() string {
	return fmt.Sprintf("--hls_media_sequence_number=%d", h)
}

// Validate checks if the value of HLSMediaSequenceNumber is valid.
func (h HLSMediaSequenceNumberFlag) Validate() error {
	return nil
}

// HLSStartTimeOffsetFlag represents the floating-point number for EXT-X-START time offset.
type HLSStartTimeOffsetFlag float32

// Parse returns the string representation of HLSStartTimeOffset for use in the command line flags.
func (h HLSStartTimeOffsetFlag) Parse() string {
	return fmt.Sprintf("--hls_start_time_offset=%v", h)
}

// Validate checks if the value of HLSStartTimeOffset is valid.
func (h HLSStartTimeOffsetFlag) Validate() error {
	return nil
}

// HLSCreateSessionKeysFlag represents the flag to enable creation of session keys for Offline HLS playback.
type HLSCreateSessionKeysFlag struct{}

// Parse returns the string representation of HLSCreateSessionKeys for use in the command line flags.
func (h HLSCreateSessionKeysFlag) Parse() string {
	return "--create_session_keys"
}

// Validate checks if the value of HLSCreateSessionKeys is valid.
func (h HLSCreateSessionKeysFlag) Validate() error {
	return nil // No specific validation logic for this flag
}
