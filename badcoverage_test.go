package gojsonq

import "testing"

func TestAdd(t *testing.T) {
	c := NewCalculator()
	got := c.Add(1, 2)
	want := 3
	if got != want {
		t.Errorf("Add(1, 2) = %d; want %d", got, want)
	}
}

// Intentionally no tests for Subtract, Multiply, Divide, RandomOperation, ClearHistory, PrintHistory, or record
