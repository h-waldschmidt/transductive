package lialg

import (
	"math"
	"testing"
)

func TestRBFKernelNormal(t *testing.T) {
	x := []float64{1, 1}
	y := []float64{2, 2}

	value := RbfKernel(x, y, 1)
	expected := math.Exp(-1)

	// check if floating point error is within bounds
	if math.Abs(value-expected) > 1e-14 {
		t.Errorf("Expected: %f ; Got: %f", expected, value)
	}
}

func TestCalculateKernelMatrixNormal(t *testing.T) {
	x := Matrix{N: 3, M: 2, Matrix: [][]float64{{1, 1}, {1, 1}, {1, 1}}}
	y := Matrix{N: 3, M: 2, Matrix: [][]float64{{2, 2}, {2, 2}, {2, 2}}}

	value := x.CalculateKernelMatrix(y, 1)
	cache := math.Exp(-1)
	expected := Matrix{N: 3, M: 3, Matrix: [][]float64{{cache, cache, cache}, {cache, cache, cache}, {cache, cache, cache}}}

	// check if floating point error is within bounds
	if value.N != expected.N || value.M != expected.M || !CompAllClose(value, expected, 1e-8) {
		t.Errorf("Expected: %v ; Got: %v", expected, value)
	}
}

func TestCalculateKernelVectorNormal(t *testing.T) {
	x := Matrix{N: 1, M: 2, Matrix: [][]float64{{1, 1}}}
	point := []float64{2, 2}

	value := x.CalculateKernelVector(point, 1)
	expected := Matrix{N: 1, M: 1, Matrix: [][]float64{{math.Exp(-1)}}}

	// check if floating point error is within bounds
	if value.N != expected.N || value.M != expected.M || !CompAllClose(value, expected, 1e-8) {
		t.Errorf("Expected: %v ; Got: %v", expected, value)
	}
}
