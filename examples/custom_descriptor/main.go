package main

import (
	"fmt"
	"log"

	packlit "github.com/m4urici0gm/packlit/pkg"
)

type MyCustomDescritor string

// Implement fmt.Stringer interface
func (m MyCustomDescritor) String() string {
	return fmt.Sprintf("my_custom_descriptor=%s", string(m))
}

// Implement gaka.StreamOption interface
func (m MyCustomDescritor) Validate() error {
	// Custom validation...
	return nil
}

// Wrap your custom descriptor into a func.
func WithCustomDescritpor(val string) packlit.StreamOpt {
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
