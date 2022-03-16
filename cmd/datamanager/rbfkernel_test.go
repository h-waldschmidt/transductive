package datamanager_test

import (
	"math"
	"testing"
	"transductive-experimental-design/cmd/datamanager"
)

func TestRBFKernel_Basic(t *testing.T) {
	x := datamanager.Coordinate{X1: 1, X2: 1}
	y := datamanager.Coordinate{X1: 2, X2: 2}

	value := datamanager.RbfKernel(x, y, 1)
	expected := math.Exp(-1)

	if value != expected {
		t.Errorf("Expected: %f ; Got: %f", expected, value)
	}
}
