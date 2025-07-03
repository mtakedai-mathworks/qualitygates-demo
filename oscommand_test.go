package dangerous

import (
	"strings"
	"testing"
)

func TestRunCommand(t *testing.T) {
	out, err := RunCommand("hello")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !strings.Contains(out, "hello") {
		t.Errorf("Expected output to contain 'hello', got: %s", out)
	}
}

func TestSafeRunCommand(t *testing.T) {
	out, err := SafeRunCommand("world")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !strings.Contains(out, "world") {
		t.Errorf("Expected output to contain 'world', got: %s", out)
	}
}
