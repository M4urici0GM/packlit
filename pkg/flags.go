// Copyright 2024 Mauricio Barbosa. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

package packlit

import (
	"fmt"
	"net/url"
	"strings"
)

var (
	_ ShakaParser = (*MpdOutput)(nil)
	_             = (*AllowApproximateSegmentTimeline)(nil)
	_             = (*GenerateStaticLiveMpd)(nil)
	_             = (*DumpStreamInfo)(nil)
	_             = (*BaseUrls)(nil)
)

func buildFlags(flags *ShakaFlags) ([]string, error) {
	built := make([]string, 0)
	for _, flag := range flags.Flags {
		if err := flag.Validate(); err != nil {
			return nil, err
		}

		built = append(built, flag.Parse())

		//// This is to make sure, flags must be separated with two elements
		//// instead of one string with the space between.
		//flagValue := flag.Parse()
		//groups := strings.Split(flagValue, " ")
		//if len(groups) < 2 {
		//	built = append(built, flagValue)
		//	continue
		//}
		//
		//built = append(built, []string{groups[0], groups[1]}...)
	}

	return built, nil
}

// MpdOutput MPD output file name.
// Equivalent to --mpd_output=file.mpd
type MpdOutput string

func (m MpdOutput) Parse() string {
	return fmt.Sprintf("--mpd_output=%s", string(m))
}

func (m MpdOutput) Validate() error {
	return nil
}

// AllowApproximateSegmentTimeline adds --allow_approximate_segment_timeline flag
type AllowApproximateSegmentTimeline struct{}

func (a AllowApproximateSegmentTimeline) Validate() error {
	return nil
}

func (a AllowApproximateSegmentTimeline) Parse() string {
	return "--allow_approximate_segment_timeline"
}

// GenerateStaticLiveMpd If enabled, generates static mpd.
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

// DumpStreamInfo Shaka Packager can be used to inspect the content of a media file and dump basic stream.
type DumpStreamInfo struct{}

func (g DumpStreamInfo) Parse() string {
	return "--dump_stream_info"
}

func (g DumpStreamInfo) Validate() error {
	return nil
}

// Comma separated BaseURLs for the MPD:
// <url>[,<url>]….
//
// The values will be added as <BaseURL> element(s) immediately under the <MPD> element.
type BaseUrls []string

func (b BaseUrls) Parse() string {
	return fmt.Sprintf("--base_urls=%s", strings.Join(b, ","))
}

func (b BaseUrls) Validate() error {
	for _, baseUrl := range b {
		if _, err := url.ParseRequestURI(baseUrl); err != nil {
			return err
		}
	}

	return nil
}
