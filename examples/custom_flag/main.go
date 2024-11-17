package main

import (
	"fmt"
	"log"

	packlit "github.com/m4urici0gm/packlit/pkg"
)

type MyCustomFlag struct{}

// Compile-time check for interface.
var _ packlit.ShakaParser = (*MyCustomFlag)(nil)

// Implement fmt.Stringer interface
func (m MyCustomFlag) Parse() string {
	return fmt.Sprintf("--my-custom-flag")
}

// Implement gaka.OptionParser interface
func (m MyCustomFlag) Validate() error {
	// Custom validation...
	return nil
}

// Wrap your custom descriptor into a func.
func WithMyCustomFlag() packlit.ShakaFlagFn {
	return func(so *packlit.ShakaFlags) {
		so.Add(MyCustomFlag{})
	}
}

func main() {
	// Equivalent to
	// $ packager in=some_content.mp4,stream=video,output=video.mp4 \
	//          --my-custom-flag
	flags := packlit.NewFlags(WithMyCustomFlag())
	runner := packlit.NewBuilder().
		WithStream(
			packlit.NewStreamBuilder().
				WithInput("some_content.mp4").
				WithStream(packlit.StreamTypeVideo).
				WithOutput("video.mp4").
				Build(),
		).
		WithFlag(flags).
		Build()

	if err := runner.Run(); err != nil {
		log.Fatalf("error running shaka-packager: %v", err)
	}
}