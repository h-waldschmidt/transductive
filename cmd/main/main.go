package main

import (
	"transductive-experimental-design/cmd/transductive"
)

func main() {
	distribution := transductive.CreateNormalDistribution(0, 1, 50)

	transductive.PlotDistribution(distribution, "../../plots/test.png")
}
