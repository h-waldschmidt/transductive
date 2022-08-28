package kmeans

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"sync"
	"time"
	"transductive-experimental-design/internal/lialg"
)

// Clusters saves all the central points (centroids) as a matrix
// Each point of the points Matrix is assigned to a matrix in the assignments slice
// index of point in points matrix represents index in assignments slice
// assignments slice saves the index of the cluster it is assigned to
type Clusters struct {
	Points        *lialg.Matrix
	NumOfClusters int
	ClusterPoints [][]int
	Centroids     lialg.Matrix
	Assignments   []int
}

// cluster the data by using the basic k means clustering algorithm
// the centroids in the first rounds are randomly initialized
func Calculate(points lialg.Matrix, numOfClusters int) (Clusters, error) {
	clusters := Clusters{
		&points,
		numOfClusters,
		make([][]int, numOfClusters),
		*lialg.NewMatrix(numOfClusters, points.M),
		make([]int, points.N),
	}

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
				d = math.Min(d, lialg.EuclideanDistance(point, centroid))
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
					distance := lialg.EuclideanDistance(
						points.Matrix[j],
						clusters.Centroids.Matrix[k],
					)

					if distance < minDistance {
						minDistance = distance
						bestClusterIndex = k
					}
				}
				clusters.Assignments[j] = bestClusterIndex
			}(j)
		}
		wg.Wait()
		clusters.updateClusters()
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

// calculates the silhouette coefficient of the cluster
// see: https://en.wikipedia.org/wiki/Silhouette_(clustering)
func (clusters *Clusters) SilhouetteCoefficient() float64 {
	var coefficient float64
	for i := range clusters.Points.Matrix {
		if i == 10 {
			fmt.Printf("")
		}
		coefficient += clusters.silhouette(i)
	}
	return coefficient / float64(len(clusters.Points.Matrix))
}

func (clusters *Clusters) silhouette(index int) float64 {
	var intraDistance float64
	clusterIndex := clusters.Assignments[index]
	if len(clusters.ClusterPoints[clusterIndex]) > 1 {
		for _, val := range clusters.ClusterPoints[clusterIndex] {

			intraDistance += lialg.EuclideanDistance(
				clusters.Points.Matrix[val],
				clusters.Points.Matrix[index],
			)
		}
		intraDistance /= float64(len(clusters.ClusterPoints[clusterIndex]) - 1)
	}

	var nearestClusterDistance float64
	for _, val := range clusters.ClusterPoints[clusterIndex] {
		d := math.Inf(1)
		for i := range clusters.Centroids.Matrix {
			if i != clusterIndex {
				var cache float64
				for _, ind := range clusters.ClusterPoints[i] {
					cache += lialg.EuclideanDistance(clusters.Points.Matrix[val], clusters.Points.Matrix[ind])
				}
				cache /= float64(len(clusters.ClusterPoints[i]))
				d = math.Min(d, cache)
			}
		}
		nearestClusterDistance += d
	}
	nearestClusterDistance /= float64(len(clusters.Centroids.Matrix) - 1)
	silhouette := (nearestClusterDistance - intraDistance) / math.Max(nearestClusterDistance, intraDistance)
	return silhouette
}

// helper function that updates the centroids by calculating
// the average of the items in the cluster
func (clusters *Clusters) updateClusters() {
	clusters.Centroids = *lialg.NewMatrix(clusters.Centroids.N, clusters.Centroids.M)

	var err error
	clusterNumOfItems := make([]int, clusters.Centroids.N)

	// reset cluster points
	clusters.ClusterPoints = make([][]int, clusters.NumOfClusters)

	for i, cluster := range clusters.Assignments {
		clusterNumOfItems[cluster]++

		clusters.ClusterPoints[cluster] = append(clusters.ClusterPoints[cluster], i)

		clusters.Centroids.Matrix[cluster], err = sliceAddition(
			clusters.Centroids.Matrix[cluster],
			clusters.Points.Matrix[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < clusters.Centroids.N; i++ {
		if clusterNumOfItems[i] != 0 {
			clusters.Centroids.Matrix[i] = sliceMultiplication(
				clusters.Centroids.Matrix[i],
				1/float64(clusterNumOfItems[i]),
			)
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
