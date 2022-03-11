package main

import (
	"transductive-experimental-design/cmd/datamanager"
)

func main() {
	distribution := datamanager.CreateNormalDistribution(0, 1, 50)

	datamanager.PlotDistribution(distribution, "../../plots/test.png")
}
