package packlit

import (
	"fmt"
)

var (
	_ ShakaParser = (*PlayReadyEncryptionFlag)(nil)
	_ ShakaParser = (*PlayReadyServerURLFlag)(nil)
	_ ShakaParser = (*ProgramIdentifierFlag)(nil)
)

// PlayReadyEncryptionFlag represents the flag to enable PlayReady encryption.
type PlayReadyEncryptionFlag struct{}

// Parse returns the string representation of PlayReadyEncryption for use in the command line flags.
func (p PlayReadyEncryptionFlag) Parse() string {
	return "--enable_playready_encryption"
}

// Validate checks if the PlayReadyEncryption value is valid.
func (p PlayReadyEncryptionFlag) Validate() error {
	return nil
}

// PlayReadyServerURLFlag represents the PlayReady packaging server URL.
type PlayReadyServerURLFlag string

// Parse returns the string representation of PlayReadyServerURL for use in the command line flags.
func (p PlayReadyServerURLFlag) Parse() string {
	return fmt.Sprintf("--playready_server_url=%s", p)
}

func (p PlayReadyServerURLFlag) Validate() error {
	return nil
}

// ProgramIdentifierFlag represents the program identifier for packaging requests.
type ProgramIdentifierFlag string

// Parse returns the string representation of ProgramIdentifier for use in the command line flags.
func (p ProgramIdentifierFlag) Parse() string {
	return fmt.Sprintf("--program_identifier=%s", p)
}

func (p ProgramIdentifierFlag) Validate() error {
	return nil
}
