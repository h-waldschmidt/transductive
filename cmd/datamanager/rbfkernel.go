package datamanager

import (
	"log"
	"math"
)

// using the popular rbfKernel (https://en.wikipedia.org/wiki/Radial_basis_function_kernel)
// is necessary for the kernel regression and many parts of the algorithms
func RbfKernel(x1 []float64, x2 []float64, sigma float64) float64 {

	//x and y need to have the same dimensions
	if len(x1) != len(x2) {
		log.Fatal("could not use RBFKernel, points do not have the same dimensions")
	}

	result := EuclideanDistance(x1, x2)

	result = math.Pow(result, 2) / (2 * sigma)
	result = math.Exp(-result)

	return result
}
