package main

import (
	"transductive-experimental-design/cmd/datamanager"
	"transductive-experimental-design/cmd/transductive"
)

func main() {
	distribution := datamanager.CreateNormalDistribution(0, 0.1, 2)
	anotherDistribution := datamanager.CreateNormalDistribution(4, 0.1, 2)
	distribution = append(distribution, anotherDistribution...)

	test := transductive.SequentialOptimization(distribution, 2, 1.8, 1.5)
	datamanager.PlotSelectedPoints(distribution, test, "../../plots/test_sequential.png")
}
