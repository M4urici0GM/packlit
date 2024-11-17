// Copyright 2024 Mauricio Barbosa. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

package packlit

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type ShakaExecutor struct {
	shaka  *ShakaPackager
	logger *log.Logger
}

func NewExecutor(shaka *ShakaPackager, logger *log.Logger) *ShakaExecutor {
	if logger == nil {
		logger = log.New(os.Stdout, "[packlit] ", log.LstdFlags)
	}

	return &ShakaExecutor{shaka, logger}
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

func (e *ShakaExecutor) RunAsync(ctx context.Context) (chan<- error, error) {
	args, err := e.shaka.BuildAndValidate()
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
	cmd := exec.CommandContext(ctx, binaryPath, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	go func() {
		err := readLog(stdout, func(msg string) {
			log.Printf("[debug] %s", msg)
		})

		if err != nil {
			log.Println(fmt.Errorf("error reading stdout: %v", err))
		}
	}()

	go func() {
		err := readLog(stderr, func(msg string) {
			log.Printf("[stderr] %s", msg)
		})

		if err != nil {
			log.Println(fmt.Errorf("error reading stderr: %v", err))
		}
	}()

	return cmd.Run()
}