package gojsonq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	expected := "Hello, SonarQube!"
	actual := greeting()
	assert.Equal(t, actual, expected, "Greeting should match expected value")
}
