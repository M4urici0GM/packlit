package gaka

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildStreamSelector(t *testing.T) {
	// given
	descriptor := &StreamOptions{
		options: []StreamOption{},
	}

	WithInput("some_content.mp4")(descriptor)
	WithStream("video")(descriptor)
	WithOutput("video.mp4")(descriptor)

	// when
	result, err := buildStreamDescriptor(descriptor)

	// then
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, "input=some_content.mp4,stream_selector=video,output=video.mp4", result)
}

func TestInputFormat(t *testing.T) {
	// given
	input := InputSelector("file.mp4")

	// when
	str := input.String()

	// then
	assert.Equal(t, "input=file.mp4", str)
}

func TestInputValidate(t *testing.T) {
	// given
	input := InputSelector("")

	// when
	result := input.Validate()

	// then
	assert.Error(t, result, "input empty is not allowed.")
}

func TestStreamSelectorFormat(t *testing.T) {
	// given
	selector := StreamSelector("audio")

	// when
	result := selector.String()

	// then
	assert.Equal(t, "stream_selector=audio", result)
}

func TestStreamSelectorValidate(t *testing.T) {
	// given
	allowed := []string{"audio", "video", "1", "0"}

	// wen
	for _, allowed := range allowed {
		selector := StreamSelector(allowed)

		result := selector.Validate()

		// then
		assert.NoError(t, result)
	}
}

func TestStreamSelectorValidateError(t *testing.T) {
	// given
	list := []string{"random", "-1"}

	// when
	for _, notAllowed := range list {
		selector := StreamSelector(notAllowed)
		result := selector.Validate()

		// then
		assert.Error(t, result)
	}
}
