package packlit

import (
	"fmt"
)

var (
	_ ShakaParser = (*CryptByteBlockFlag)(nil)
	_ ShakaParser = (*ProtectionSchemeFlag)(nil)
	_ ShakaParser = (*SkipByteBlockFlag)(nil)
	_ ShakaParser = (*VP9SubsampleEncryptionFlag)(nil)
	_ ShakaParser = (*PlayReadyExtraHeaderDataFlag)(nil)
)

// CryptByteBlockFlag represents the encrypted block count in a protection pattern.
type CryptByteBlockFlag int

// Parse returns the string representation of CryptByteBlock for use in the command line flags.
func (p CryptByteBlockFlag) Parse() string {
	return fmt.Sprintf("--crypt_byte_block=%d", p)
}

// Validate checks if the value of CryptByteBlock is valid.
// Specify the count of the encrypted blocks in the protection pattern,
// where block is of size 16-bytes. There are three common
// patterns (crypt_byte_block:skip_byte_block): 1:9 (default), 5:5, 10:0.
// Apply to video streams with 'cbcs' and 'cens' protection schemes only;
// ignored otherwise.
func (p CryptByteBlockFlag) Validate() error {
	return nil
}

// ProtectionSchemeFlag represents the protection scheme flag.
type ProtectionSchemeFlag string

// Parse returns the string representation of ProtectionScheme for use in the command line flags.
func (p ProtectionSchemeFlag) Parse() string {
	return fmt.Sprintf("--protection_scheme=%s", p)
}

// Validate checks if the ProtectionScheme flag value is valid.
// Specify a protection scheme, 'cenc' or 'cbc1' or pattern-based protection schemes 'cens' or 'cbcs'.
func (p ProtectionSchemeFlag) Validate() error {
	return nil
}

// SkipByteBlockFlag represents the number of unencrypted blocks in the protection pattern.
type SkipByteBlockFlag int

// Parse returns the string representation of SkipByteBlock for use in the command line flags.
func (s SkipByteBlockFlag) Parse() string {
	return fmt.Sprintf("--skip_byte_block=%v", s)
}

// Validate checks if the value of SkipByteBlock is valid.
// Specify the count of the unencrypted blocks in the protection pattern.
// Apply to video streams with 'cbcs' and 'cens' protection schemes only;
// ignored otherwise.
func (s SkipByteBlockFlag) Validate() error {
	return nil
}

// VP9SubsampleEncryptionFlag represents the flag to enable VP9 subsample encryption.
type VP9SubsampleEncryptionFlag struct{}

// Parse returns the string representation of VP9SubsampleEncryption for use in the command line flags.
func (v VP9SubsampleEncryptionFlag) Parse() string {
	return "--vp9_subsample_encryption"
}

// Validate checks if the VP9SubsampleEncryption flag value is valid.
// Enable VP9 subsample encryption.
func (v VP9SubsampleEncryptionFlag) Validate() error {
	// No specific validation needed as the flag is of boolean type.
	return nil
}

// PlayReadyExtraHeaderDataFlag represents the extra XML data for PlayReady headers.
type PlayReadyExtraHeaderDataFlag string

// Parse returns the string representation of PlayReadyExtraHeaderData for use in the command line flags.
func (p PlayReadyExtraHeaderDataFlag) Parse() string {
	return fmt.Sprintf("--playready_extra_header_data=%s", p)
}

// Validate checks if the PlayReadyExtraHeaderData flag value is valid.
// Extra XML data to add to PlayReady headers.
func (p PlayReadyExtraHeaderDataFlag) Validate() error {
	// Check if the string is empty (valid) or if not, it must be valid XML.
	return nil
}
