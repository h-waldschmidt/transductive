package plt

import (
	"testing"
	"transductive-experimental-design/cmd/lialg"

	"github.com/google/go-cmp/cmp"
)

func TestConvertSliceToCoordinateNormal(t *testing.T) {
	point := []float64{1, 1}

	result := ConvertSliceToCoordinate(point)
	expected := Coordinate{1, 1}

	if !cmp.Equal(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestConvertMatrixToCoordinateSliceNormal(t *testing.T) {
	matrix := lialg.Matrix{N: 3, M: 2, Matrix: [][]float64{{3, 1}, {2, 0}, {1, 2}}}

	result := ConvertMatrixToCoordinateSlice(&matrix)
	expected := []Coordinate{{3, 1}, {2, 0}, {1, 2}}

	if !cmp.Equal(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestConvertCoordinatesToMatrixNormal(t *testing.T) {
	input := []Coordinate{{3, 1}, {2, 0}, {1, 2}}

	result := ConvertCoordinatesToMatrix(input)
	expected := lialg.Matrix{N: 3, M: 2, Matrix: [][]float64{{3, 1}, {2, 0}, {1, 2}}}

	if !cmp.Equal(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}
