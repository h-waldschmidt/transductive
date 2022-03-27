package main

import (
	"transductive-experimental-design/cmd/datamanager"
	"transductive-experimental-design/cmd/transductive"
)

func main() {
	distribution := datamanager.CreateNormalDistribution(0, 0.1, 10)
	anotherDistribution := datamanager.CreateNormalDistribution(4, 0.1, 10)
	distribution = append(distribution, anotherDistribution...)
	another := datamanager.CreateNormalDistribution(8, 0.1, 10)

	distribution = append(distribution, another...)
	matrix := datamanager.ConvertCoordinatesToMatrix(distribution)

	test := transductive.SequentialOptimization(matrix, 3, 1.8, 1.5)
	distribution_test, _ := datamanager.ConvertMatrixToCoordinateSlice(test)
	datamanager.PlotSelectedPoints(distribution, distribution_test, "../../plots/test_sequential.png")
}
