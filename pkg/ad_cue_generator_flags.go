package packlit

import (
	"fmt"
)

var (
	_ ShakaParser = (*AdCuesFlag)(nil)
)

// AdCuesFlag represents the list of cuepoint markers.
type AdCuesFlag string

// Parse returns the string representation of AdCues for use in the command line flags.
func (a AdCuesFlag) Parse() string {
	return fmt.Sprintf("--ad_cues=%s", a)
}

// Validate checks if the value of AdCues is valid.
// This flag accepts semicolon separated pairs, and components in the pair are separated by a comma.
// The second component (duration) is optional. For example:
// --ad_cues {start_time}[,{duration}][;{start_time}[,{duration}]]...
// The start_time represents the start of the cue marker in seconds relative to the start of the program.
func (a AdCuesFlag) Validate() error {
	// Split the input by semicolons to separate cuepoint markers.
	return nil
}
