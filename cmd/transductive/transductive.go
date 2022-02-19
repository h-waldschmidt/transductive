package transductive

import "math"

func Transductive() {

}

func rbfKernel(x1 [2]float64, x2 [2]float64, variance float64) float64 {
	value := -math.Pow(x1[0]-x2[0], 2) + math.Pow(x1[1]-x2[1], 2)
	value = math.Sqrt(value)
	value /= 2 * variance
	value = math.Exp(value)

	return value
}
