package datamanager_test

import (
	"math"
	"testing"
	"transductive-experimental-design/cmd/datamanager"
)

func TestEuclideanDistance_Basic(t *testing.T) {
	x := datamanager.Matrix{N: 1, M: 3, Matrix: [][]float64{{1, 2, 3}}}
	y := datamanager.Matrix{N: 1, M: 3, Matrix: [][]float64{{2, 3, 4}}}

	value, err := datamanager.EuclideanDistance(x.Matrix[0], y.Matrix[0])
	if err != nil {
		t.Errorf("%v", err)
	}
	expected := math.Sqrt(3)

	if value != expected {
		t.Errorf("Expected: %f ; Got: %f", expected, value)
	}
}

func TestEuclideanDistance_Error(t *testing.T) {
	x := datamanager.Matrix{N: 1, M: 3, Matrix: [][]float64{{1, 2, 3}}}
	y := datamanager.Matrix{N: 1, M: 4, Matrix: [][]float64{{2, 3, 4, 5}}}

	_, error := datamanager.EuclideanDistance(x.Matrix[0], y.Matrix[0])

	if error == nil {
		t.Errorf("Expected Error")
	}
}

func TestEuclideanNorm_Basic(t *testing.T) {
	x := datamanager.Matrix{N: 1, M: 3, Matrix: [][]float64{{1, 2, 3}}}

	value := datamanager.EuclideanNorm(x.Matrix[0])
	expected := math.Sqrt(14)

	if value != expected {
		t.Errorf("Expected: %f ; Got: %f", expected, value)
	}
}
