package packlit

import (
	"encoding/json"
	"testing"
)

func TestBuildFromJSON_Minimal(t *testing.T) {
	req := PackagerOptions{
		Binary: "/usr/bin/packager",
		Streams: []StreamDescriptor{{
			Input:  "/in.mp4",
			Stream: "video",
			Output: "/out.mp4",
		}},
		MPD: &MpdFlags{MpdOutput: "/manifest.mpd"},
	}

	sp, err := BuildFromJSON(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cmd, err := sp.PreviewCommand()
	if err != nil {
		t.Fatalf("preview error: %v", err)
	}
	if cmd == "" {
		t.Fatalf("empty command preview")
	}

	args, err := sp.BuildAndValidate()
	if err != nil {
		t.Fatalf("build error: %v", err)
	}
	found := false
	for _, a := range args {
		if a == "--mpd_output=/manifest.mpd" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected mpd_output flag in args: %v", args)
	}
}

func TestBuildFromJSON_Full(t *testing.T) {
	payload := `{
		"binary": "/usr/bin/packager",
		"streams": [
			{
				"input": "/in.mp4",
				"stream": "video",
				"output": "/out.mp4",
				"segment_template": "/seg_$Time$.m4s",
				"init_segment": "/init.mp4"
			}
		],
		"packager": {
			"quiet": true,
			"vmodule": "http_file=1"
		},
		"muxer": {
			"segment_duration": 4,
			"segment_sap_aligned": true
		},
		"mpd": {
			"mpd_output": "/manifest.mpd",
			"base_urls": ["https://cdn.example/"],
			"generate_static_live_mpd": true,
			"allow_approximate_segment_timeline": true,
			"dump_stream_info": true
		},
		"manifest": {
			"default_language": "eng"
		},
		"http": {
			"user_agent": "packlit/1.0"
		},
		"hls": {
			"hls_master_playlist_output": "/master.m3u8"
		},
		"encryption": {
			"protection_scheme": "cbcs",
			"protection_systems": ["Widevine"],
			"raw_key": {
				"enable_raw_key_encryption": true,
				"keys": [
					{
						"label": "video",
						"key_id": "00",
						"key": "11"
					}
				]
			},
			"widevine": {
				"enable_widevine_encryption": true,
				"key_server_url": "https://license",
				"signer": "abc"
			},
			"playready": {
				"enable_playready_encryption": true,
				"playready_server_url": "https://pr"
			}
		}
	}
	`

	var req PackagerOptions
	if err := json.Unmarshal([]byte(payload), &req); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	sp, err := BuildFromJSON(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	args, err := sp.BuildAndValidate()
	if err != nil {
		t.Fatalf("build error: %v", err)
	}

	expect := []string{
		"--mpd_output=/manifest.mpd",
		"--base_urls=https://cdn.example/",
		"--generate_static_live_mpd",
		"--allow_approximate_segment_timeline",
		"--dump_stream_info",
		"--segment_duration=4",
		"--quiet",
		"--hls_master_playlist_output=/master.m3u8",
		"--protection_scheme=cbcs",
	}
	for _, e := range expect {
		found := false
		for _, a := range args {
			if a == e {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected flag %q not found in args: %v", e, args)
		}
	}
}
