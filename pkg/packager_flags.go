package packlit

import "fmt"

var (
	_ ShakaParser = (*LicensesFlag)(nil)
	_ ShakaParser = (*QuietFlag)(nil)
	_ ShakaParser = (*UseFakeClockForMuxerFlag)(nil)
	_ ShakaParser = (*TestPackagerVersionFlag)(nil)
	_ ShakaParser = (*SingleThreadedFlag)(nil)
	_ ShakaParser = (*VModuleFlag)(nil)
	_ ShakaParser = (*VersionFlag)(nil)
)

// LicensesFlag represents the flag for dumping licenses information.
type LicensesFlag struct{}

func (g LicensesFlag) Parse() string {
	return "--licenses"
}

func (g LicensesFlag) Validate() error {
	return nil
}

// QuietFlag represents the flag for suppressing LOG(INFO) output.
type QuietFlag struct{}

func (g QuietFlag) Parse() string {
	return "--quiet"
}

func (g QuietFlag) Validate() error {
	return nil
}

// UseFakeClockForMuxerFlag represents the flag for using a fake clock for muxing.
type UseFakeClockForMuxerFlag struct{}

func (g UseFakeClockForMuxerFlag) Parse() string {
	return "--use_fake_clock_for_muxer"
}

func (g UseFakeClockForMuxerFlag) Validate() error {
	return nil
}

// TestPackagerVersionFlag represents the flag for specifying the packager version for testing.
type TestPackagerVersionFlag struct{}

func (g TestPackagerVersionFlag) Parse() string {
	return "--test_packager_version"
}

func (g TestPackagerVersionFlag) Validate() error {
	return nil
}

// SingleThreadedFlag represents the flag for enabling single-threaded content generation.
type SingleThreadedFlag struct{}

func (g SingleThreadedFlag) Parse() string {
	return "--single_threaded"
}

func (g SingleThreadedFlag) Validate() error {
	return nil
}

// VModuleFlag adds verbose logging through --v or --vmodule command line flags.
// Equivalent to --vmodule=http_file=1
type VModuleFlag string

func (v VModuleFlag) Parse() string {
	return fmt.Sprintf("--vmodule=%s", v)
}

func (v VModuleFlag) Validate() error {
	return nil
}

// VersionFlag adds --version for checking packager version
type VersionFlag struct{}

func (v VersionFlag) Parse() string {
	return "--version"
}

func (v VersionFlag) Validate() error {
	return nil
}
