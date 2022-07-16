package lialg

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMatrixMultiplicationNormal(t *testing.T) {
	a := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}
	b := Matrix{2, 3, [][]float64{{1, 0, 4}, {2, 1, 0}}}

	value := MatrixMultiplication(a, b)
	expected := Matrix{2, 2, [][]float64{{7, 9}, {8, 2}}}

	if !cmp.Equal(value, expected) {
		t.Errorf("Expected: %v ; Got: %v", expected, value)
	}
}

func TestEuclideanDistanceNormal(t *testing.T) {
	x := Matrix{N: 1, M: 3, Matrix: [][]float64{{1, 2, 3}}}
	y := Matrix{N: 1, M: 3, Matrix: [][]float64{{2, 3, 4}}}

	value := EuclideanDistance(x.Matrix[0], y.Matrix[0])
	expected := math.Sqrt(3)

	if value != expected {
		t.Errorf("Expected: %f ; Got: %f", expected, value)
	}
}

func TestEuclideanNormNormal(t *testing.T) {
	x := Matrix{N: 1, M: 3, Matrix: [][]float64{{1, 2, 3}}}

	value := EuclideanNorm(x.Matrix[0])
	expected := math.Sqrt(14)

	if value != expected {
		t.Errorf("Expected: %f ; Got: %f", expected, value)
	}
}

func TestSumNormNormal(t *testing.T) {
	x := Matrix{N: 1, M: 3, Matrix: [][]float64{{1, 2, -3}}}

	value := SumNorm(x.Matrix[0])
	expected := 6.0

	if value != expected {
		t.Errorf("Expected: %f ; Got: %f", expected, value)
	}
}

func TestSliceToDiagonalMatrixNormal(t *testing.T) {
	x := []float64{1, 2, 3}

	value := SliceToDiagonalMatrix(x)
	expected := Matrix{N: 3, M: 3, Matrix: [][]float64{{1, 0, 0}, {0, 2, 0}, {0, 0, 3}}}

	if !cmp.Equal(value, expected) {
		t.Errorf("Expected: %v ; Got: %v", expected, value)
	}
}

func TestDiagonalMatrixToSliceNormal(t *testing.T) {
	x := Matrix{N: 3, M: 3, Matrix: [][]float64{{1, 0, 0}, {0, 2, 0}, {0, 0, 3}}}

	value := x.DiagonalMatrixToSlice()
	expected := []float64{1, 2, 3}

	if !cmp.Equal(value, expected) {
		t.Errorf("Expected: %v ; Got: %v", expected, value)
	}
}

func TestVectorToDiagonalMatrixNormal(t *testing.T) {
	x := Matrix{N: 1, M: 3, Matrix: [][]float64{{1, 2, 3}}}

	value := x.VectorToDiagonalMatrix()
	expected := Matrix{N: 3, M: 3, Matrix: [][]float64{{1, 0, 0}, {0, 2, 0}, {0, 0, 3}}}

	if !cmp.Equal(value, expected) {
		t.Errorf("Expected: %v ; Got: %v", expected, value)
	}
}

func TestDiagonalMatrixToVectorNormal(t *testing.T) {
	x := Matrix{N: 3, M: 3, Matrix: [][]float64{{1, 0, 0}, {0, 2, 0}, {0, 0, 3}}}

	value := x.DiagonalMatrixToVector()
	expected := Matrix{N: 1, M: 3, Matrix: [][]float64{{1, 2, 3}}}

	if !cmp.Equal(value, expected) {
		t.Errorf("Expected: %v ; Got: %v", expected, value)
	}
}
