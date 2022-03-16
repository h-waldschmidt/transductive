package datamanager

import (
	"math"
	"testing"
)

func TestRBFKernel_Basic(t *testing.T) {
	x := Coordinate{X1: 1, X2: 1}
	y := Coordinate{X1: 2, X2: 2}

	value := RbfKernel(x, y, 1)
	expected := math.Exp(-1)
	if value != expected {
		t.Errorf("Expected: %f ; Got: %f", expected, value)
	}
}
