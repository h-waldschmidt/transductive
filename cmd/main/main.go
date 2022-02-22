package main

import (
	"../transductive"
)

func main() {
	distribution := transductive.CreateNormalDistribution(0, 1, 50)

	transductive.PlotDistribution(distribution, "../../plots/test.png")
}
