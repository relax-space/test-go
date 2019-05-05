package base_test

import (
	"test-go/base"
	"testing"
)

//simple test
func TestAddOne(t *testing.T) {
	var (
		in       = 1
		expected = 2
	)
	actual := base.AddOne(1)
	if actual != expected {
		t.Errorf("AddOne(%d) = %d; expected %d", in, actual, expected)
	}
}

//Table-Driven Test
func TestAddTable(t *testing.T) {
	tests := []struct {
		in  int
		exp int
	}{
		{1, 2},
		{-1, 1},
	}

	for _, tt := range tests {
		actual := base.AddOne(tt.in)
		if actual != tt.exp {
			t.Errorf("AddOne(%v) = %v; expected %d", tt.in, actual, tt.exp)
		}
	}
}
