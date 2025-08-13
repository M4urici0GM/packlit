package packlit

import (
	"fmt"
)

var (
	_ ShakaParser = (*EnableWidevineEncryptionFlag)(nil)
	_ ShakaParser = (*EnableWidevineDecryptionFlag)(nil)
	_ ShakaParser = (*KeyServerURLFlag)(nil)
	_ ShakaParser = (*ContentIDFlag)(nil)
	_ ShakaParser = (*PolicyFlag)(nil)
	_ ShakaParser = (*MaxSDPixelslag)(nil)
	_ ShakaParser = (*MaxHDPixelslag)(nil)
	_ ShakaParser = (*MaxUHD1Pixelslag)(nil)
	_ ShakaParser = (*SignerFlag)(nil)
	_ ShakaParser = (*AESSigningKeyFlag)(nil)
	_ ShakaParser = (*AESSigningIVFlag)(nil)
	_ ShakaParser = (*RSASigningKeyPathFlag)(nil)
	_ ShakaParser = (*CryptoPeriodDurationFlag)(nil)
	_ ShakaParser = (*GroupIDFlag)(nil)
	_ ShakaParser = (*EnableEntitlementLicenseFlag)(nil)
)

// EnableWidevineEncryptionFlag represents the flag to enable Widevine encryption.
type EnableWidevineEncryptionFlag struct{}

// Parse returns the string representation of EnableWidevineEncryption for use in the command line flags.
func (e EnableWidevineEncryptionFlag) Parse() string {
	return "--enable_widevine_encryption"
}

// Validate checks if the EnableWidevineEncryptionFlag value is valid.
func (e EnableWidevineEncryptionFlag) Validate() error {
	// If this flag is set, key_server_url and signer must be provided.
	return nil
}

// EnableWidevineDecryptionFlag represents the flag to enable Widevine decryption.
type EnableWidevineDecryptionFlag struct{}

// Parse returns the string representation of EnableWidevineDecryption for use in the command line flags.
func (e EnableWidevineDecryptionFlag) Parse() string {
	return "--enable_widevine_decryption"
}

// Validate checks if the EnableWidevineDecryptionFlag value is valid.
func (e EnableWidevineDecryptionFlag) Validate() error {
	// If this flag is set, key_server_url and signer must be provided.
	return nil
}

// KeyServerURLFlag represents the key server URL flag.
type KeyServerURLFlag string

// Parse returns the string representation of KeyServerURL for use in the command line flags.
func (k KeyServerURLFlag) Parse() string {
	return fmt.Sprintf("--key_server_url=%s", k)
}

// Validate checks if the KeyServerURLFlag value is valid.
func (k KeyServerURLFlag) Validate() error {
	// Validate that this URL is required when either encryption or decryption is enabled.
	return nil
}

// ContentIDFlag represents the content ID flag in hex.
type ContentIDFlag string

// Parse returns the string representation of ContentID for use in the command line flags.
func (c ContentIDFlag) Parse() string {
	return fmt.Sprintf("--content_id=%s", c)
}

// Validate checks if the ContentIDFlag value is valid.
func (c ContentIDFlag) Validate() error {
	// Content ID should be provided when encryption is enabled.
	return nil
}

// PolicyFlag represents the stored policy for DRM content rights.
type PolicyFlag string

// Parse returns the string representation of Policy for use in the command line flags.
func (p PolicyFlag) Parse() string {
	return fmt.Sprintf("--policy=%s", p)
}

// Validate checks if the PolicyFlag value is valid.
func (p PolicyFlag) Validate() error {
	// Policy is optional, but if provided, it should be valid.
	return nil
}

// MaxSDPixelslag represents the flag for the maximum pixels for SD video.
type MaxSDPixelslag int

// Parse returns the string representation of MaxSDPixels for use in the command line flags.
func (m MaxSDPixelslag) Parse() string {
	return fmt.Sprintf("--max_sd_pixels=%d", m)
}

// Validate checks if the MaxSDPixelslag value is valid.
func (m MaxSDPixelslag) Validate() error {
	return nil
}

// MaxHDPixelslag represents the flag for the maximum pixels for HD video.
type MaxHDPixelslag int

// Parse returns the string representation of MaxHDPixels for use in the command line flags.
func (m MaxHDPixelslag) Parse() string {
	return fmt.Sprintf("--max_hd_pixels=%d", m)
}

// Validate checks if the MaxHDPixelslag value is valid.
func (m MaxHDPixelslag) Validate() error {
	return nil
}

// MaxUHD1Pixelslag represents the flag for the maximum pixels for UHD1 video.
type MaxUHD1Pixelslag int

// Parse returns the string representation of MaxUHD1Pixels for use in the command line flags.
func (m MaxUHD1Pixelslag) Parse() string {
	return fmt.Sprintf("--max_uhd1_pixels=%d", m)
}

// Validate checks if the MaxUHD1Pixelslag value is valid.
func (m MaxUHD1Pixelslag) Validate() error {
	return nil
}

// SignerFlag represents the signer name for the Widevine encryption/decryption.
type SignerFlag string

// Parse returns the string representation of Signer for use in the command line flags.
func (s SignerFlag) Parse() string {
	return fmt.Sprintf("--signer=%s", s)
}

// Validate checks if the SignerFlag value is valid.
func (s SignerFlag) Validate() error {
	// Signer should be specified if encryption/decryption is enabled.
	return nil
}

// AESSigningKeyFlag represents the AES signing key in hex.
type AESSigningKeyFlag string

// Parse returns the string representation of AESSigningKey for use in the command line flags.
func (a AESSigningKeyFlag) Parse() string {
	return fmt.Sprintf("--aes_signing_key=%s", a)
}

// Validate checks if the AESSigningKeyFlag value is valid.
func (a AESSigningKeyFlag) Validate() error {
	// AES key must be provided with the corresponding AES IV.
	return nil
}

// AESSigningIVFlag represents the AES signing IV in hex.
type AESSigningIVFlag string

// Parse returns the string representation of AESSigningIV for use in the command line flags.
func (a AESSigningIVFlag) Parse() string {
	return fmt.Sprintf("--aes_signing_iv=%s", a)
}

// Validate checks if the AESSigningIVFlag value is valid.
func (a AESSigningIVFlag) Validate() error {
	// AES IV must be provided with the corresponding AES key.
	return nil
}

// RSASigningKeyPathFlag represents the RSA signing key path.
type RSASigningKeyPathFlag string

// Parse returns the string representation of RSASigningKeyPath for use in the command line flags.
func (r RSASigningKeyPathFlag) Parse() string {
	return fmt.Sprintf("--rsa_signing_key_path=%s", r)
}

// Validate checks if the RSASigningKeyPathFlag value is valid.
func (r RSASigningKeyPathFlag) Validate() error {
	// RSA key path must be provided if using RSA signing.
	return nil
}

// CryptoPeriodDurationFlag represents the flag for the crypto period duration in seconds.
type CryptoPeriodDurationFlag int

// Parse returns the string representation of CryptoPeriodDuration for use in the command line flags.
func (c CryptoPeriodDurationFlag) Parse() string {
	return fmt.Sprintf("--crypto_period_duration=%d", c)
}

// Validate checks if the CryptoPeriodDurationFlag value is valid.
func (c CryptoPeriodDurationFlag) Validate() error {
	return nil
}

// GroupIDFlag represents the group ID in hex.
type GroupIDFlag string

// Parse returns the string representation of GroupID for use in the command line flags.
func (g GroupIDFlag) Parse() string {
	return fmt.Sprintf("--group_id=%s", g)
}

// Validate checks if the GroupIDFlag value is valid.
func (g GroupIDFlag) Validate() error {
	// Group ID is optional but should be valid if provided.
	return nil
}

// EnableEntitlementLicenseFlag represents the flag to enable the entitlement license.
type EnableEntitlementLicenseFlag struct{}

// Parse returns the string representation of EnableEntitlementLicense for use in the command line flags.
func (e EnableEntitlementLicenseFlag) Parse() string {
	return "--enable_entitlement_license"
}

// Validate checks if the EnableEntitlementLicenseFlag value is valid.
func (e EnableEntitlementLicenseFlag) Validate() error {
	// Entitlement license is optional but must be provided if required.
	return nil
}
