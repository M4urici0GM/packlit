// Copyright 2024 Mauricio Barbosa. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

package packlit

import (
	"fmt"
	"strings"
)

func buildFlags(flags *ShakaFlags) (string, error) {
	built := make([]string, 0)
	for _, flag := range flags.Flags {
		if err := flag.Validate(); err != nil {
			return "", err
		}

		built = append(built, flag.Parse())
	}

	return strings.Join(built, " "), nil
}

// MPD output file name.
// Equivalent to --mpd_output=file.mpd
type MpdOutput string

func (m MpdOutput) Parse() string {
	return fmt.Sprintf("--mpd_output=%s", string(m))
}

func (m MpdOutput) Validate() error {
	return nil
}

// If enabled, generates static mpd.
// If segment_template is specified in stream descriptors, shaka-packager generates dynamic mpd by default;
// if this flag is enabled, shaka-packager generates static mpd instead.
// Note that if segment_template is not specified, shaka-packager always generates static mpd regardless of the value of this flag.
type GenerateStaticLiveMpd struct{}

func (g GenerateStaticLiveMpd) Parse() string {
	return "--generate_static_live_mpd"
}

func (g GenerateStaticLiveMpd) Validate() error {
	return nil
}

// Shaka Packager can be used to inspect the content of a media file and dump basic stream.
type DumpStreamInfo struct{}

func (g DumpStreamInfo) Parse() string {
	return "--dump_stream_info"
}

func (g DumpStreamInfo) Validate() error {
	return nil
}
