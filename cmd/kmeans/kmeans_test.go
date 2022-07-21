package kmeans

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSliceAdditionNormal(t *testing.T) {
	a := []float64{1, 2, 3, 4, 5}
	b := []float64{1, 2, 3, 4, 5}
	ans, err := sliceAddition(a, b)

	if len(ans) != 5 {
		t.Error("answer has unexpected length")
	}

	if err != nil {
		t.Errorf("didn't expect error: %v", err)
	}

	expected := []float64{2, 4, 6, 8, 10}

	if !cmp.Equal(expected, ans) {
		t.Error("unexpected value")
	}
}

func TestSliceAdditionError(t *testing.T) {
	a := []float64{1, 2, 3, 4, 5}
	b := []float64{1, 2, 3, 4}
	ans, err := sliceAddition(a, b)

	if ans != nil || err == nil {
		t.Error("Expected error")
	}
}

func TestSliceMultiplicationNormal(t *testing.T) {
	a := []float64{1, 2, 3, 4, 5}

	t.Run("Zero Slice", func(t *testing.T) {
		ans := sliceMultiplication(a, 0)
		expected := make([]float64, 5)

		if !cmp.Equal(expected, ans) {
			t.Error("expected zero slice")
		}
	})

	t.Run("Same", func(t *testing.T) {
		ans := sliceMultiplication(a, 1)
		expected := []float64{1, 2, 3, 4, 5}

		if !cmp.Equal(expected, ans) {
			t.Error("expected same slice")
		}
	})

	t.Run("Normal", func(t *testing.T) {
		ans := sliceMultiplication(a, 2)
		expected := []float64{2, 4, 6, 8, 10}

		if !cmp.Equal(expected, ans) {
			t.Error("unexpected value")
		}
	})
}
