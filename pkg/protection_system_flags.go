package packlit

import (
	"fmt"
	"strings"
)

var (
	_ ShakaParser = (*ProtectionSystemsFlag)(nil)
)

// ProtectionSystemsFlag represents the flag for specifying the protection systems.
type ProtectionSystemsFlag []string

// Parse returns the string representation of ProtectionSystemsFlag for use in the command line flags.
func (p ProtectionSystemsFlag) Parse() string {
	return fmt.Sprintf("--protection_systems=%s", strings.Join(p, ","))
}

// Validate checks if the protection systems are valid.
func (p ProtectionSystemsFlag) Validate() error {
	return nil
}
