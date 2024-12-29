// Copyright 2024 Mauricio Barbosa. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

package packlit

import (
	"context"
	"log"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockCommandRunner struct {
	mock.Mock
}

func (m *MockCommandRunner) CommandContext(ctx context.Context, name string, args ...string) *exec.Cmd {
	argsSlice := m.Called(ctx, name, args)
	return argsSlice.Get(0).(*exec.Cmd)
}

func newLogger() *log.Logger {
	return log.New(os.Stdout, "[packlit]", log.LstdFlags)
}

func TestShakaExecutorRun(t *testing.T) {
	// given
	ctx := context.Background()

	mockRunner := &MockCommandRunner{}
	mockRunner.On("CommandContext", ctx, "mock_binary", []string{}).
		Return(exec.Command("echo", "hello world"))

	shaka := NewShakaPackager("mock_binary")
	executor := &ShakaExecutor{
		logger:        newLogger(),
		shaka:         shaka,
		commandRunner: mockRunner,
	}

	// when
	err := executor.RunWithContext(ctx)

	// then
	require.NoError(t, err)
	mockRunner.AssertCalled(t, "CommandContext", ctx, "mock_binary", []string{})
}

func TestShakaExecutorRunError(t *testing.T) {
	// given
	ctx := context.Background()

	mockRunner := &MockCommandRunner{}
	mockRunner.On("CommandContext", ctx, "shaka-packager", []string{}).
		Return(exec.Command("false"))

	shaka := NewShakaPackager("shaka-packager")
	executor := &ShakaExecutor{
		logger:        newLogger(),
		shaka:         shaka,
		commandRunner: mockRunner,
	}

	// when
	err := executor.RunWithContext(ctx)

	// then
	require.Error(t, err)
	require.Contains(t, err.Error(), "exit status 1")

	mockRunner.AssertCalled(t, "CommandContext", ctx, mock.Anything, mock.Anything)
}

func TestShakaExecutorRunAsync(t *testing.T) {
	ctx := context.Background()

	mockRunner := &MockCommandRunner{}
	mockRunner.On("CommandContext", ctx, "mock_binary", []string{""}).
		Return(exec.Command("echo", "hello world"))

	shaka := NewShakaPackager("mock_binary")
	executor := &ShakaExecutor{
		logger:        newLogger(),
		shaka:         shaka,
		commandRunner: mockRunner,
	}

	// when
	chn, err := executor.RunAsync(ctx)

	// then
	require.NoError(t, err)
	require.NoError(t, <-chn)
	mockRunner.AssertCalled(t, "CommandContext", ctx, "mock_binary", []string{""})
}

func TestShakaExecutorErrorAsync(t *testing.T) {
	ctx := context.Background()

	mockRunner := &MockCommandRunner{}
	mockRunner.On("CommandContext", ctx, "shaka-packager", []string{""}).
		Return(exec.Command("false"))

	shaka := NewShakaPackager("shaka-packager")
	executor := &ShakaExecutor{
		logger:        newLogger(),
		shaka:         shaka,
		commandRunner: mockRunner,
	}

	// when
	chn, err := executor.RunAsync(ctx)

	// then
	require.NoError(t, err)

	err = <-chn
	require.Error(t, err)
	require.Contains(t, err.Error(), "exit status 1")

	mockRunner.AssertCalled(t, "CommandContext", ctx, mock.Anything, mock.Anything)
}