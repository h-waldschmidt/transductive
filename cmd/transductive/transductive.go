package transductive

import (
	"math"
)

func Transductive() {

}

// using the popular rbfKernel (https://en.wikipedia.org/wiki/Radial_basis_function_kernel)
// is necessary for the kernel regression
func rbfKernel(x1 Coordinate, x2 Coordinate, variance float64) float64 {
	value := -math.Pow(x1.X1-x2.X1, 2) + math.Pow(x1.X2-x2.X2, 2)
	value = math.Sqrt(value)
	value /= 2 * variance
	value = math.Exp(value)

	return value
}
