package packlit

import "fmt"

// MPD output file name.
// Equivalent to --mpd_output=file.mpd
type MpdOutput string

func (m MpdOutput) String() string {
	return fmt.Sprintf("--mpd_output=%s", string(m))
}

func (m MpdOutput) Validate() error {
	return nil
}

func WithMpdOutput(output string) func(*ShakaFlags) {
	return func(so *ShakaFlags) {
		so.Add(MpdOutput(output))
	}
}

// If enabled, generates static mpd.
// If segment_template is specified in stream descriptors, shaka-packager generates dynamic mpd by default;
// if this flag is enabled, shaka-packager generates static mpd instead.
// Note that if segment_template is not specified, shaka-packager always generates static mpd regardless of the value of this flag.
type GenerateStaticLiveMpd struct{}

func (g GenerateStaticLiveMpd) String() string {
	return "--generate_static_live_mpd"
}

func (g GenerateStaticLiveMpd) Validate() error {
	return nil
}

type DumpStreamInfo struct{}

func (g DumpStreamInfo) String() string {
	return "--dump_stream_info"
}

func (g DumpStreamInfo) Validate() error {
	return nil
}

func WithDumpStream() ShakaFlagFn {
	return func(so *ShakaFlags) {
		so.Add(DumpStreamInfo{})
	}
}
