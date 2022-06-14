package main

import (
	"log"
	"transductive-experimental-design/cmd/plt"
	"transductive-experimental-design/cmd/transductive"
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
	test := transductive.AlternatingOptimization(matrix, 4, 1, 1)
	distributionTest := plt.ConvertMatrixToCoordinateSlice(&test)

	err := plt.PlotSelectedPoints(distribution, distributionTest, "../../plots/test_alternating.png")
	if err != nil {
		log.Fatal(err)
	}
}
