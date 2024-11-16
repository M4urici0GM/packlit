package gaka

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildDescriptors(t *testing.T) {
	// given
	opts := []*StreamOptions{
		{
			[]StreamOption{
				InputSelector("file.mp4"),
				StreamSelector("audio"),
				OutputSelector("audio.mp4"),
			},
		},
		{
			[]StreamOption{
				InputSelector("file.mp4"),
				StreamSelector("video"),
				OutputSelector("video.mp4"),
			},
		},
	}

	// when
	result, err := buildStreamDescriptors(opts...)

	// then
	assert.NoError(t, err)
	assert.Equal(t, "input=file.mp4,stream_selector=audio,output=audio.mp4 input=file.mp4,stream_selector=video,output=video.mp4", result)
}

func TestBuildDescriptorWithInvalid(t *testing.T) {
	// given
	opts := []*StreamOptions{
		{
			[]StreamOption{
				InputSelector("file.mp4"),
				StreamSelector("invalid-stream-selector"),
				OutputSelector("audio.mp4"),
			},
		},
	}

	// when
	result, err := buildStreamDescriptors(opts...)

	// then
	assert.Error(t, err)
	assert.Empty(t, result)
}

func TestBuildFlags(t *testing.T) {
	// given
	flags := &ShakaOptions{
		flags: []ShakaFlag{
			MpdOutput("file.mpd"),
			GenerateStaticLiveMpd{},
		},
	}

	// when
	result, err := buildFlags(flags)

	// then
	assert.NoError(t, err)
	assert.Equal(t, "--mpd_output=file.mpd --generate_static_live_mpd", result)
}
