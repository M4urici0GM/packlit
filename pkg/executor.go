// Copyright 2024 Mauricio Barbosa. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

package packlit

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"os/exec"
)

type CommandRunner interface {
	CommandContext(context.Context, string, ...string) *exec.Cmd
}

type NativeCommandRunner struct{}

func (n *NativeCommandRunner) CommandContext(ctx context.Context, name string, args ...string) *exec.Cmd {
	return exec.CommandContext(ctx, name, args...)
}

type ShakaExecutor struct {
	shaka         *ShakaPackager
	logger        *log.Logger
	commandRunner CommandRunner
}

// Creates new executor.
// If logger is nil, it will create a new one with prefix [packlit]
func NewExecutor(shaka *ShakaPackager, logger ...*log.Logger) *ShakaExecutor {
	return &ShakaExecutor{
		shaka:         shaka,
		logger:        getLogger(logger...),
		commandRunner: &NativeCommandRunner{},
	}
}

func (e *ShakaExecutor) RunWithContext(ctx context.Context) error {
	args, err := e.shaka.BuildAndValidate()
	if err != nil {
		return err
	}

	return e.runShaka(ctx, e.shaka.Binary, args)
}

func (e *ShakaExecutor) Run() error {
	return e.RunWithContext(context.Background())
}

func (e *ShakaExecutor) RunAsync(ctx context.Context) (<-chan error, error) {
	args, flags, err := e.shaka.Args()
	if err != nil {
		return nil, err
	}

	errorChn := make(chan error)

	go func() {
		defer close(errorChn)

		select {
		case <-ctx.Done():
			errorChn <- ctx.Err()
			return
		default:
            args = append(args, flags)
			err := e.runShaka(ctx, e.shaka.Binary, args)
			if err != nil {
				errorChn <- err
			}
		}
	}()

	return errorChn, nil
}

func readLog(reader io.Reader, fn func(string)) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		msg := scanner.Text()
		fn(msg)
	}

	return scanner.Err()
}

func (e *ShakaExecutor) runShaka(ctx context.Context, binaryPath string, args []string) error {
	cmd := e.commandRunner.CommandContext(ctx, binaryPath, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		log.Fatalf("failed to start command: %v", err)
	}

	go func() {
		_ = readLog(stdout, func(msg string) {
			log.Printf("[debug] %s", msg)
		})
	}()

	go func() {
		_ = readLog(stderr, func(msg string) {
			log.Printf("[stderr] %s", msg)
		})
	}()

	return cmd.Wait()
}

func getLogger(logger ...*log.Logger) *log.Logger {
	if len(logger) != 0 && logger[0] != nil {
		return logger[0]
	}

	return log.New(os.Stdout, "[packlit] ", log.LstdFlags)
}
