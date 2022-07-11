package kmeans

import (
	"testing"
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

	for i := 0; i < 5; i++ {
		if ans[i] != a[i]+b[i] {
			t.Error("unexpected value")
		}
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
