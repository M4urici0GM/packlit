package packlit 

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestBuildDescriptors(t *testing.T) {
	// given
	opts := []*StreamOptions{
		{
			[]ShakaParser{
				InputSelector("file.mp4"),
				StreamSelector("audio"),
				OutputSelector("audio.mp4"),
			},
		},
		{
			[]ShakaParser{
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
	assert.Equal(t, []string{"input=file.mp4,stream_selector=audio,output=audio.mp4", "input=file.mp4,stream_selector=video,output=video.mp4"}, result)
}

func TestBuildDescriptorWithInvalid(t *testing.T) {
	// given
	opts := []*StreamOptions{
		{
			[]ShakaParser{
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
	flags := &ShakaFlags{
		Flags: []ShakaParser{
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

func TestShouldTransmuxCorrectly(t *testing.T) {
	// given

	// given
	opts := []*StreamOptions{
		{
			[]ShakaParser{
				InputSelector("/var/input/proof.mp4"),
				StreamSelector("audio"),
				OutputSelector("/var/media/audio.mp4"),
			},
		},
		{
			[]ShakaParser{
				InputSelector("/var/input/proof.mp4"),
				StreamSelector("video"),
				OutputSelector("/var/media/video.mp4"),
			},
		},
	}

	runner := &ShakaRunner{
		StreamOptions: opts,
		Flags:         &ShakaFlags{},
	}

	// Build descriptors
	args, _, err := runner.Args()
	require.NoError(t, err)

	// Build file path
	absPath, err := filepath.Abs(filepath.Join("testdata", "proof.mp4"))
	require.NoError(t, err)

	// Check for file
	r, err := os.Open(absPath)
	require.NoError(t, err)

	// Build CMD
	containerCmd := []string{"/usr/bin/packager"}
	containerCmd = append(containerCmd, args...)

	outputPath, err := filepath.Abs(filepath.Join("testdata", strings.ReplaceAll(uuid.NewString(), "-", "")))
	require.NoError(t, err)

	err = os.MkdirAll(outputPath, 0755)
	require.NoError(t, err)

	ctx := context.Background()
	_, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: "google/shaka-packager",
			Cmd:   containerCmd,
			Files: []testcontainers.ContainerFile{
				{
					Reader:            r,
					HostFilePath:      absPath,
					ContainerFilePath: "/var/input/proof.mp4",
					FileMode:          0o700,
				},
			},
			HostConfigModifier: func(hc *container.HostConfig) {
				hc.AutoRemove = false
				hc.Mounts = []mount.Mount{
					{
						Type:   mount.TypeBind,
						Source: fmt.Sprintf("%s/", outputPath),
						Target: "/var/media/",
					},
				}
			},
			WaitingFor: wait.ForLog("Packaging completed successfully."),
		},
		Started: true,
	})

	require.NoError(t, err)

	// then
	for _, n := range []string{"video.mp4", "audio.mp4"} {
		verifyFileExistence(t, filepath.Join(outputPath, n))
	}

	cleanupFiles(t, outputPath)
}

func cleanupFiles(t *testing.T, path string) {
	err := os.RemoveAll(path)
	require.NoError(t, err)
}

func verifyFileExistence(t *testing.T, path string) {
	_, err := os.Stat(path)
	require.NoError(t, err)
}
