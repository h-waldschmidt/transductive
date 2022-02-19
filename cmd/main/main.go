package main

import (
	transductive "transductive-experimental-design/cmd/transductive/datamanager"
)

func main() {
	transductive.CreateNormalDistribution(0, 1, 50)
}
