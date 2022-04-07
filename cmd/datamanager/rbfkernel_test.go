package datamanager_test

import (
	"math"
	"testing"
	"transductive-experimental-design/cmd/datamanager"
)

func TestRBFKernel_Basic(t *testing.T) {
	x := []float64{1, 1}
	y := []float64{2, 2}

	value, err := datamanager.RbfKernel(x, y, 1)
	if err != nil {
		t.Errorf("%v", err)
	}
	expected := math.Exp(-1)

	if value != expected {
		t.Errorf("Expected: %f ; Got: %f", expected, value)
	}
}
