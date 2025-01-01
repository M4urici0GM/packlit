// Copyright 2024 Mauricio Barbosa. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

package packlit

import (
    "bufio"
    "context"
    "fmt"
    "io"
    "os"
    "path"
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
    assert.Equal(t, []string{"--mpd_output=file.mpd", "--generate_static_live_mpd"}, result)
}

func TestShouldTransmuxCorrectly(t *testing.T) {
    // given
    type AssetTrack struct {
        Id   string
        Type string
    }

    // Build file path
    absPath, err := filepath.Abs(filepath.Join("testdata", "proof.mp4"))
    require.NoError(t, err)

    t.Log("absPath:", absPath)

    // Check for file
    r, err := os.Open(absPath)
    require.NoError(t, err)

    outputPath, err := filepath.Abs(filepath.Join("testdata", strings.ReplaceAll(uuid.NewString(), "-", "")))
    require.NoError(t, err)

    err = os.MkdirAll(outputPath, 0755)
    require.NoError(t, err)
   
    defer func(path string, t *testing.T) {
        err := os.RemoveAll(path)
        if err != nil {
            t.Logf("failed to remove output directory: %s", path)
        }
    }(outputPath, t)

    // many tracks as you can have.
    assets := []AssetTrack{
        {uuid.New().String(), "video"},
    }

    shaka := NewBuilder()
    mpdOutput := path.Join("/var/media", "h264.mpd")

    for _, assetTrack := range assets {
        streamType := StreamType(assetTrack.Type)

        basePath := path.Join("/var/media", assetTrack.Id)
        initSegment := path.Join(basePath, "init.mp4")
        segmentTemplate := path.Join(basePath, "$Time$.m4s")

        shaka.WithStream(
            NewStreamBuilder().
                WithInput("/var/input/proof.mp4").
                WithStream(streamType).
                AddOption(WithInitSegment(initSegment)).
                AddOption(WithSegmentTemplate(segmentTemplate)).
                Build(),
        )
    }

    t.Log("What is shaka? {}", shaka)

    shakaFlags := NewFlags(
        WithStaticLiveMpd(),
        WithSegmentDuration("4"),
        WithMpdOutput(mpdOutput),
        WithApproximateSegmentTimeline())

    runner := shaka.WithFlag(shakaFlags).Build()

    // Build descriptors
    args, flags, err := runner.Args()
    require.NoError(t, err)

    // Build CMD
    var containerCmd []string
    containerCmd = append(containerCmd, "/usr/bin/packager")
    containerCmd = append(containerCmd, args...)
    containerCmd = append(containerCmd, flags...)

    t.Log("container cmd is:", containerCmd)

    ctx := context.Background()
    c, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
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
                        BindOptions: &mount.BindOptions{
                            CreateMountpoint: true,
                        },
                    },
                }
            },
            WaitingFor: wait.ForLog("Packaging completed successfully."),
        },
        Started: true,
    })

    if log, err := c.Logs(ctx); err == nil {
        reader := bufio.NewReader(log)
        messageBuf := make([]byte, reader.Size())
        _, _ = io.ReadFull(log, messageBuf)

        t.Log("Reading logs:", string(messageBuf))
    }

    require.NoError(t, err)
}

func cleanupFiles(t *testing.T, path string) {
    err := os.RemoveAll(path)
    require.NoError(t, err)
}

func verifyFileExistence(t *testing.T, path string) {
    _, err := os.Stat(path)
    require.NoError(t, err)
}
