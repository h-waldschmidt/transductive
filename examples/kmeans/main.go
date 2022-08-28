package main

import (
	"fmt"
	"log"
	"transductive-experimental-design/pkg/kmeans"
	"transductive-experimental-design/pkg/plt"
)

func main() {
	distribution := plt.CreateNormalDistribution(0, 0.1, 10)
	anotherDistribution := plt.CreateNormalDistribution(4, 0.1, 10)
	distribution = append(distribution, anotherDistribution...)
	another := plt.CreateNormalDistribution(8, 0.1, 10)
	distribution = append(distribution, another...)
	another = plt.CreateNormalDistribution(6, 0.1, 10)
	distribution = append(distribution, another...)

	matrix := plt.ConvertCoordinatesToMatrix(distribution)

	test, err := kmeans.Calculate(matrix, 40)
	if err != nil {
		log.Fatal(err)
	}
	distributionTest := plt.ConvertMatrixToCoordinateSlice(&test.Centroids)

	err = plt.PlotSelectedPoints(distribution, distributionTest, "plots/test_kmeans.png")
	if err != nil {
		log.Fatal(err)
	}

	inertia := test.Inertia()
	fmt.Printf("%v", inertia)
}
