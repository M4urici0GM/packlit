package packlit

import "fmt"

var (
	_ ShakaParser = (*UserAgentFlag)(nil)
	_ ShakaParser = (*CaFileFlag)(nil)
	_ ShakaParser = (*ClientCertFileFlag)(nil)
	_ ShakaParser = (*ClientCertPrivateKeyFileFlag)(nil)
	_ ShakaParser = (*ClientCertPrivateKeyPasswordFlag)(nil)
	_ ShakaParser = (*DisablePeerVerificationFlag)(nil)
	_ ShakaParser = (*IgnoreHttpOutputFailuresFlag)(nil)
	_ ShakaParser = (*IoCacheSizeFlag)(nil)
)

// UserAgentFlag represents the flag for setting a custom User-Agent string for HTTP requests.
type UserAgentFlag string

func (u UserAgentFlag) Parse() string {
	return fmt.Sprintf("--user_agent=%s", u)
}

func (u UserAgentFlag) Validate() error {
	return nil
}

// CaFileFlag represents the flag for setting the Certificate Authority file for the server cert.
type CaFileFlag string

func (c CaFileFlag) Parse() string {
	return fmt.Sprintf("--ca_file=%s", c)
}

func (c CaFileFlag) Validate() error {
	return nil
}

// ClientCertFileFlag represents the flag for setting the client certificate file path.
type ClientCertFileFlag string

func (c ClientCertFileFlag) Parse() string {
	return fmt.Sprintf("--client_cert_file=%s", c)
}

func (c ClientCertFileFlag) Validate() error {
	return nil
}

// ClientCertPrivateKeyFileFlag represents the flag for setting the Private Key file path for the client cert.
type ClientCertPrivateKeyFileFlag string

func (c ClientCertPrivateKeyFileFlag) Parse() string {
	return fmt.Sprintf("--client_cert_private_key_file=%s", c)
}

func (c ClientCertPrivateKeyFileFlag) Validate() error {
	return nil
}

// ClientCertPrivateKeyPasswordFlag represents the flag for setting the password to the private key file.
type ClientCertPrivateKeyPasswordFlag string

func (c ClientCertPrivateKeyPasswordFlag) Parse() string {
	return fmt.Sprintf("--client_cert_private_key_password=%s", c)
}

func (c ClientCertPrivateKeyPasswordFlag) Validate() error {
	return nil
}

// DisablePeerVerificationFlag represents the flag for disabling peer verification for HTTP connections.
type DisablePeerVerificationFlag struct{}

func (d DisablePeerVerificationFlag) Parse() string {
	return "--disable_peer_verification"
}

func (d DisablePeerVerificationFlag) Validate() error {
	return nil
}

// IgnoreHttpOutputFailuresFlag represents the flag for ignoring HTTP output failures.
type IgnoreHttpOutputFailuresFlag struct{}

func (i IgnoreHttpOutputFailuresFlag) Parse() string {
	return "--ignore_http_output_failures"
}

func (i IgnoreHttpOutputFailuresFlag) Validate() error {
	return nil
}

// IoCacheSizeFlag represents the flag for setting the IO cache size.
type IoCacheSizeFlag uint

func (i IoCacheSizeFlag) Parse() string {
	return fmt.Sprintf("--io_cache_size=%d", i)
}

func (i IoCacheSizeFlag) Validate() error {
	return nil
}
