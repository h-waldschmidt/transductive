package lialg

import (
	"log"
	"math"
)

// using the popular rbfKernel (https://en.wikipedia.org/wiki/Radial_basis_function_kernel)
//
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

// Applies the Kernel function to every component of the matrix
func (pointsX *Matrix) CalculateKernelMatrix(pointsY Matrix, sigma float64) Matrix {
	//x and y need to have the same dimensions
	if pointsX.M != pointsY.M {
		log.Fatal("could not use RBFKernel, points do not have the same dimensions")
	}

	matrix := NewMatrix(pointsY.N, pointsX.N)
	for i := 0; i < matrix.N; i++ {
		for j := 0; j < matrix.M; j++ {
			matrix.Matrix[i][j] = RbfKernel(pointsX.Matrix[i], pointsY.Matrix[j], sigma)
		}
	}
	return *matrix
}

// Applies the Kernel function on every component of the vector
func (pointsX *Matrix) CalculateKernelVector(point []float64, sigma float64) Matrix {
	//x and y need to have the same dimensions
	if pointsX.M != len(point) {
		log.Fatal("could not use RBFKernel, points do not have the same dimensions")
	}

	vector := NewMatrix(1, pointsX.N)
	for i := 0; i < pointsX.N; i++ {
		vector.Matrix[0][i] = RbfKernel(pointsX.Matrix[i], point, sigma)
	}

	return *vector
}
