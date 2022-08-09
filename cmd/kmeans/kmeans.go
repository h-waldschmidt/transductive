package kmeans

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"sync"
	"time"
	"transductive-experimental-design/cmd/lialg"

	"golang.org/x/exp/constraints"
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

// cluster the data by using the basic k means clustering algorithm
// the centroids in the first rounds are randomly initialized
func Calculate(points lialg.Matrix, numOfClusters int) (Clusters, error) {
	clusters := Clusters{&points, *lialg.NewMatrix(numOfClusters, points.M), make([]int, points.N)}

	if points.N < numOfClusters {
		return clusters, fmt.Errorf("there should be at least %d points", numOfClusters)
	}

	// initialize clusters.Assignments with -1
	for i := range clusters.Assignments {
		clusters.Assignments[i] = -1
	}

	// k means++ initialization
	rand.Seed(time.Now().UnixNano())
	rand := rand.Intn(len(points.Matrix))
	clusters.Centroids.Matrix[0] = points.Matrix[rand]
	for i := 1; i < numOfClusters; i++ {

		distances := make([]float64, len(points.Matrix))
		for j, point := range points.Matrix {

			d := math.Inf(1)
			for _, centroid := range clusters.Centroids.Matrix {
				d = min(d, lialg.EuclideanDistance(point, centroid))
			}
			distances[j] = d
		}
		clusters.Centroids.Matrix[i] = points.Matrix[maxIndexSlice(distances)]
	}

	maxIterations := 50
	for i := 0; i < maxIterations; i++ {
		var wg sync.WaitGroup
		for j := range points.Matrix {
			wg.Add(1)
			go func(j int) {
				defer wg.Done()

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
			}(j)
		}
		wg.Wait()
		clusters.updateCentroids()
	}
	return clusters, nil
}

// calculates the inertia value of the cluster
// see: https://en.wikipedia.org/wiki/K-means_clustering#Global_optimization_and_metaheuristics
func (clusters *Clusters) Inertia() float64 {
	var inertia float64
	for i, point := range clusters.Points.Matrix {
		index := clusters.Assignments[i]
		inertia += lialg.EuclideanDistance(point, clusters.Centroids.Matrix[index])
	}
	return inertia
}

// helper function that updates the centroids by calculating
// the average of the items in the cluster
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

	for i := 0; i < clusters.Centroids.N; i++ {
		if clusterNumOfItems[i] != 0 {
			clusters.Centroids.Matrix[i] = sliceMultiplication(clusters.Centroids.Matrix[i], 1/float64(clusterNumOfItems[i]))
		}
	}
}

// adds the coresponding items of each slice (see example below)
// slices need to have the same length
// example: a = [0,1,2,3], b = [1,2,3,4] a+b = [1,3,5,7]
func sliceAddition(a, b []float64) ([]float64, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("slices have to be same length")
	}

	ans := make([]float64, len(a))
	for i := 0; i < len(ans); i++ {
		ans[i] = a[i] + b[i]
	}
	return ans, nil
}

// multiplies every item of a slice with a factor
func sliceMultiplication(a []float64, factor float64) []float64 {
	ans := make([]float64, len(a))
	for i := 0; i < len(ans); i++ {
		ans[i] = factor * a[i]
	}
	return ans
}

// return smaller value of two given values
func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// return larger value of two given values
func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func maxIndexSlice(a []float64) int {
	maxIndex := -1
	maxValue := math.Inf(-1)
	for i, value := range a {
		if value > maxValue {
			maxIndex = i
			maxValue = value
		}
	}
	return maxIndex
}
