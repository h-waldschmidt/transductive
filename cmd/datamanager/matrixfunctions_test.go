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

	if !cmp.Equal(value, expected) {
		t.Errorf("Expected: %v ; Got: %v", expected, value)
	}
}

func TestEigenBasic(t *testing.T) {
	a := datamanager.Matrix{3, 3, [][]float64{{3, 2, -2}, {-1, 0, 2}, {0, 0, -1}}}

	value := a.CalculateEigen()

	expected_first_vector := datamanager.Matrix{1, 3, [][]float64{{1, 2, 1}}}
	expected_second_vector := datamanager.Matrix{1, 3, [][]float64{{1, 1, 0}}}
	expected_third_vector := datamanager.Matrix{1, 3, [][]float64{{0, 0, 1}}}
	expected := datamanager.Eigen{[]float64{1, 2, -1}, []datamanager.Matrix{expected_first_vector, expected_second_vector, expected_third_vector}}
	if !cmp.Equal(value, expected) {
		t.Errorf("Expected: %v ; Got: %v", expected, value)
	}
}

func TestQRDecompositionBasic(t *testing.T) {
	a := datamanager.Matrix{3, 3, [][]float64{{12, 6, -4}, {-51, 167, 24}, {4, -68, -41}}}

	q, r := a.QrDecomposition()

	value := datamanager.MatrixMultiplication(q, r)
	if a.N != value.N || a.M != value.M || !datamanager.CompAllClose(a, value, 1e-12) {
		t.Errorf("Expected: %v ; Got: %v", a, value)
	}
}
