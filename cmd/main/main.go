package main

import (
	transductive "transductive-experimental-design/cmd/transductive/datamanager"
)

func main() {
	distribution := transductive.CreateNormalDistribution(0, 1, 50)

	transductive.PlotDistribution(distribution, "../../plots/test.png")
}
