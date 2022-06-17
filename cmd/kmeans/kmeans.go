package kmeans

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
	"transductive-experimental-design/cmd/lialg"
)

// Clusters saves all the central points (centroids) as a matrix
// Each point of the points Matrix is assigned to a matrix in the assignments slice
// index of point in points matrix represents index in assignments slice
// assignments slice saves the index of the cluster it is assigned to
type Clusters struct {
	Points      *lialg.Matrix
	Centroids   lialg.Matrix
	Assignments []int
}

func Calculate(points lialg.Matrix, numOfClusters int) (Clusters, error) {
	clusters := Clusters{&points, *lialg.NewMatrix(numOfClusters, points.M), make([]int, points.N)}

	if points.N < numOfClusters {
		return clusters, fmt.Errorf("there should be at least %d points", numOfClusters)
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numOfClusters; i++ {
		clusters.Centroids.Matrix[i] = points.Matrix[rand.Intn(numOfClusters+1)]
	}

	maxIterations := 50
	for i := 0; i < maxIterations; i++ {
		for j := 0; j < points.N; j++ {
			minDistance := math.Inf(1)
			bestClusterIndex := -1
			for k := 0; k < clusters.Centroids.N; k++ {
				distance := lialg.EuclideanDistance(points.Matrix[j], clusters.Centroids.Matrix[k])

				if distance < minDistance {
					minDistance = distance
					bestClusterIndex = k
				}
			}
			clusters.Assignments[j] = bestClusterIndex
		}
		clusters.updateCentroids()
	}
	return clusters, nil
}

func (clusters *Clusters) updateCentroids() {
	clusters.Centroids = *lialg.NewMatrix(clusters.Centroids.N, clusters.Centroids.M)

	var err error
	clusterNumOfItems := make([]int, clusters.Centroids.N)

	for i, cluster := range clusters.Assignments {
		clusterNumOfItems[cluster]++

		clusters.Centroids.Matrix[cluster], err = sliceAddition(clusters.Centroids.Matrix[cluster], clusters.Points.Matrix[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < clusters.Centroids.M; i++ {
		if clusterNumOfItems[i] != 0 {
			clusters.Centroids.Matrix[i] = sliceMultiplication(clusters.Centroids.Matrix[i], 1/float64(clusterNumOfItems[i]))
		}
	}
}

func sliceAddition(a, b []float64) ([]float64, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("slices have to be same length")
	}

	ans := a
	for i := 0; i < len(ans); i++ {
		ans[i] += b[i]
	}
	return ans, nil
}

func sliceMultiplication(a []float64, factor float64) []float64 {
	ans := a
	for i := 0; i < len(ans); i++ {
		ans[i] *= factor
	}
	return ans
}
