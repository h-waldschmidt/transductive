package lialg

import (
	"testing"
)

func TestEigenBasic(t *testing.T) {
	a := Matrix{3, 3, [][]float64{{3, 2, -2}, {-1, 0, 2}, {0, 0, -1}}}
	//a := Matrix{4, 4, [][]float64{{26, 40, 51, 54}, {40, 67, 62, 83}, {41, 62, 95, 70}, {54, 83, 70, 126}}}
	value := a.CalculateEigen()

	expectedFirstVector := Matrix{1, 3, [][]float64{{1, 2, 1}}}
	expectedSecondVector := Matrix{1, 3, [][]float64{{1, 1, 0}}}
	expectedThirdVector := Matrix{1, 3, [][]float64{{0, 0, 1}}}
	expected := Eigen{[]float64{1, 2, -1}, []Matrix{expectedFirstVector, expectedSecondVector, expectedThirdVector}}
	if !CompAllClose(expectedFirstVector, value.Vectors[1], 1e-4) || !CompAllClose(expectedSecondVector, value.Vectors[0], 1e-4) || !CompAllClose(expectedThirdVector, value.Vectors[2], 1e-4) {
		t.Errorf("Expected: %v ; Got: %v", expected, value)
	}
}

func TestQRDecompositionBasic(t *testing.T) {
	a := Matrix{3, 3, [][]float64{{12, 6, -4}, {-51, 167, 24}, {4, -68, -41}}}

	q, r := a.QrDecomposition()

	value := MatrixMultiplication(q, r)
	if a.N != value.N || a.M != value.M || !CompAllClose(a, value, 1e-12) {
		t.Errorf("Expected: %v ; Got: %v", a, value)
	}
}
