package datamanager_test

import (
	"testing"
	"transductive-experimental-design/cmd/datamanager"

	"github.com/google/go-cmp/cmp"
)

func TestMatrixMultiplicationBasic(t *testing.T) {
	a := datamanager.Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}
	b := datamanager.Matrix{2, 3, [][]float64{{1, 0, 4}, {2, 1, 0}}}

	value := datamanager.MatrixMultiplication(a, b)
	expected := datamanager.Matrix{2, 2, [][]float64{{7, 9}, {8, 2}}}

	// check if floating point error is within bounds
	if !cmp.Equal(value, expected) {
		t.Errorf("Expected: %v ; Got: %v", expected, value)
	}
}
