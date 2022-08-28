package kmeans

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/h-waldschmidt/transductive/internal/lialg"
	"github.com/h-waldschmidt/transductive/pkg/plt"
)

func TestClusteringError(t *testing.T) {
	x := lialg.Matrix{N: 1, M: 2, Matrix: [][]float64{{1, 2}}}
	_, err := Calculate(x, 2)

	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestClustering(t *testing.T) {
	// just generate some points with normal distributions
	distribution := plt.CreateNormalDistribution(0, 0.1, 10)
	anotherDistribution := plt.CreateNormalDistribution(4, 0.1, 10)
	distribution = append(distribution, anotherDistribution...)

	matrix := plt.ConvertCoordinatesToMatrix(distribution)

	clustering, err := Calculate(matrix, 2)

	if err != nil {
		t.Errorf("Expected: No error Got: %v", err)
	}

	// check if data of the clustering makes sense

	if !cmp.Equal(*clustering.Points, matrix) {
		t.Errorf(
			"Points inside clustering are not same. Got: %v, Expected: %v",
			*&clustering.Points,
			matrix,
		)
	}

	if clustering.NumOfClusters != 2 {
		t.Errorf("Wrong number of clusters. Expected: %v Got: %v", 2, clustering.NumOfClusters)
	}

	if len(clustering.ClusterPoints) != 2 {
		t.Errorf("Unexpected ClusterPoints: %v", clustering.ClusterPoints)
	}

	if clustering.Centroids.N != 2 || clustering.Centroids.M != 2 ||
		len(clustering.Centroids.Matrix) != 2 || len(clustering.Centroids.Matrix[0]) != 2 ||
		len(clustering.Centroids.Matrix[1]) != 2 {
		t.Errorf("Centroids have unexpected dimensions: %v", clustering.Centroids)
	}

	if len(clustering.Assignments) != 20 {
		t.Errorf("Unexpected assignments: %v", clustering.Assignments)
	}

	// test Inertia and SilhouetteCoefficient
	inertia := clustering.Inertia()
	if inertia == math.NaN() || inertia == 0 || inertia == math.Inf(1) || inertia == math.Inf(-1) {
		t.Errorf("Unexpected Inertia value: %f", inertia)
	}

	silhouette := clustering.SilhouetteCoefficient()
	if silhouette == math.NaN() || silhouette < -1 || silhouette > 1 {
		t.Errorf("Unexpected SilhouetteCoefficient: %f", silhouette)
	}
}
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
