// Copyright 2024 Mauricio Barbosa. All rights reserved.
// Use of this source code is governed by a LGPL
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"

	packlit "github.com/m4urici0gm/packlit/pkg"
)

type MyCustomDescritor string

// Compile-time check for interface.
var _ packlit.ShakaParser = (*MyCustomDescritor)(nil)

// Implement fmt.Stringer interface
func (m MyCustomDescritor) Parse() string {
	return fmt.Sprintf("my_custom_descriptor=%s", string(m))
}

// Implement gaka.StreamOption interface
func (m MyCustomDescritor) Validate() error {
	// Custom validation...
	return nil
}

// Wrap your custom descriptor into a func.
func WithCustomDescritpor(val string) packlit.StreamDescriptorFn {
	return func(so *packlit.StreamOptions) {
		so.Add(MyCustomDescritor(val))
	}
}

func main() {
	// Equivalent to
	// $ packager in=some_content.mp4,my_custom_descriptor=some-custom-value

	runner := packlit.NewBuilder().
		WithStream(packlit.NewStreamDescriptor(packlit.WithInput("input.mp4"), WithCustomDescritpor("some-custom-value"))).
		Build()

	if err := runner.Run(); err != nil {
		log.Fatalf("error running shaka-packager: %v", err)
	}
}
