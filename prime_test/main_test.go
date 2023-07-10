package main

import (
	"testing"
)

// Tests for the isPrime function, using table tests
func TestIsPrime(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{1, false},
		{15, false},
		{7, true},
	}

	for _, test := range tests {
		got, _ := isPrime(test.n)
		if got != test.want {
			t.Errorf("isPrime(%d) = %v, want %v", test.n, got, test.want)
		}
	}
}
