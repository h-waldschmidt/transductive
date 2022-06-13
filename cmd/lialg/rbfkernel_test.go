package lialg_test

import (
	"math"
	"testing"
	"transductive-experimental-design/cmd/lialg"
)

func TestRBFKernelBasic(t *testing.T) {
	x := []float64{1, 1}
	y := []float64{2, 2}

	value := lialg.RbfKernel(x, y, 1)
	expected := math.Exp(-1)

	// check if floating point error is within bounds
	if math.Abs(value-expected) > 1e-14 {
		t.Errorf("Expected: %f ; Got: %f", expected, value)
	}
}
