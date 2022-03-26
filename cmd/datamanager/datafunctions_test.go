package datamanager_test

import (
	"math"
	"testing"
	"transductive-experimental-design/cmd/datamanager"
)

func TestEuclideanDistance_Basic(t *testing.T) {
	x := datamanager.Matrix{N: 3, M: 1, Matrix: [][]float64{{1, 2, 3}}}
	y := datamanager.Matrix{N: 3, M: 1, Matrix: [][]float64{{2, 3, 4}}}

	value, _ := datamanager.EuclideanDistance(x, y)
	expected := math.Sqrt(3)

	if value != expected {
		t.Errorf("Expected: %f ; Got: %f", expected, value)
	}
}

func TestEuclideanDistance_Error(t *testing.T) {
	x := datamanager.Matrix{N: 3, M: 1, Matrix: [][]float64{{1, 2, 3}}}
	y := datamanager.Matrix{N: 4, M: 1, Matrix: [][]float64{{2, 3, 4, 5}}}

	_, error := datamanager.EuclideanDistance(x, y)

	if error == nil {
		t.Errorf("Expected Error")
	}
}

func TestEuclideanNorm_Basic(t *testing.T) {
	x := datamanager.Matrix{N: 3, M: 1, Matrix: [][]float64{{1, 2, 3}}}

	value, _ := datamanager.EuclideanNorm(x)
	expected := math.Sqrt(14)

	if value != expected {
		t.Errorf("Expected: %f ; Got: %f", expected, value)
	}
}

func TestEuclideanNorm_Error(t *testing.T) {
	x := datamanager.Matrix{N: 3, M: 2, Matrix: [][]float64{{1, 2, 3}, {1, 2, 3}}}

	_, error := datamanager.EuclideanNorm(x)

	if error == nil {
		t.Errorf("Expected Error")
	}
}
