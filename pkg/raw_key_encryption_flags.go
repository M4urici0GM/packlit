package packlit

import (
	"fmt"
	"strings"
)

var (
	_ ShakaParser = (*EnableRawKeyEncryptionFlag)(nil)
	_ ShakaParser = (*EnableRawKeyDecryptionFlag)(nil)
	_ ShakaParser = (*KeysFlag)(nil)
	_ ShakaParser = (*IvFlag)(nil)
	_ ShakaParser = (*PsshFlag)(nil)
)

// EnableRawKeyEncryptionFlag represents the flag to enable raw key encryption.
type EnableRawKeyEncryptionFlag struct{}

// Parse returns the string representation of EnableRawKeyEncryption for use in the command line flags.
func (e EnableRawKeyEncryptionFlag) Parse() string {
	return "--enable_raw_key_encryption"
}

// Validate checks if the EnableRawKeyEncryptionFlag value is valid.
func (e EnableRawKeyEncryptionFlag) Validate() error {
	// No specific validation logic for this boolean flag.
	return nil
}

// EnableRawKeyDecryptionFlag represents the flag to enable raw key decryption.
type EnableRawKeyDecryptionFlag struct{}

// Parse returns the string representation of EnableRawKeyDecryption for use in the command line flags.
func (e EnableRawKeyDecryptionFlag) Parse() string {
	return "--enable_raw_key_decryption"
}

// Validate checks if the EnableRawKeyDecryptionFlag value is valid.
func (e EnableRawKeyDecryptionFlag) Validate() error {
	// No specific validation logic for this boolean flag.
	return nil
}

// KeysFlag represents the flag for key information in the command line.
type KeysFlag []string

// Parse returns the string representation of KeysFlag for use in the command line flags.
func (k KeysFlag) Parse() string {
	return fmt.Sprintf("--keys=%s", strings.Join(k, ","))
}

// Validate checks if the KeysFlag value is valid.
func (k KeysFlag) Validate() error {
	return nil
}

// IvFlag represents the IV in hex string format (for testing).
type IvFlag string

// Parse returns the string representation of IvFlag for use in the command line flags.
func (i IvFlag) Parse() string {
	return fmt.Sprintf("--iv=%s", i)
}

// Validate checks if the IvFlag value is valid.
func (i IvFlag) Validate() error {
	// IV flag is optional, no specific validation needed.
	return nil
}

// PsshFlag represents the PSSH box(es) in hex string format.
type PsshFlag string

// Parse returns the string representation of PsshFlag for use in the command line flags.
func (p PsshFlag) Parse() string {
	return fmt.Sprintf("--pssh=%s", p)
}

// Validate checks if the PsshFlag value is valid.
func (p PsshFlag) Validate() error {
	// PSSH flag is optional, no specific validation needed.
	return nil
}
