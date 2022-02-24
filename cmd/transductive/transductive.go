package transductive

import (
	"fmt"
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

func calculateKernelMatrix(pointsX []Coordinate, pointsY []Coordinate, variance float64) Matrix {
	// initializing the matrix
	matrix := Matrix{len(pointsX), len(pointsY), make([][]float64, len(pointsX))}
	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i] = make([]float64, len(pointsY))
	}

	// calculating all the values
	for i := 0; i < len(pointsX); i++ {
		for j := 0; j < len(pointsY); j++ {
			matrix.Matrix[i][j] = rbfKernel(pointsX[i], pointsY[j], variance)
		}
	}

	return matrix
}

func calculateKernelVector(pointsX []Coordinate, point Coordinate, variance float64) Vector {
	// initializing the vector
	vector := Vector{len(pointsX), make([]float64, len(pointsX))}

	// calculating all the values
	for i := 0; i < len(pointsX); i++ {
		vector.Vector[i] = rbfKernel(pointsX[i], point, variance)

	}

	return vector
}

func euclideanDistance(x Vector, y Vector) (float64, error) {

	//the size of the vectors need to be the same
	if len(x.Vector) != len(y.Vector) {
		return 0, fmt.Errorf("could not calculate euclidean Distance")
	}

	var distance float64
	for i := 0; i < len(x.Vector); i++ {
		distance += math.Pow(x.Vector[i]-y.Vector[i], 2)
	}

	return math.Sqrt(distance), nil
}

func euclideanNorm(x Vector) float64 {

	var norm float64
	for i := 0; i < len(x.Vector); i++ {
		norm += math.Pow(x.Vector[i], 2)
	}

	return math.Sqrt(norm)
}

func matrixMultiplication(matrix1 Matrix, matrix2 Matrix) (Matrix, error) {

	// The inner dimensions need to be the same
	if matrix1.M != matrix2.N {
		return Matrix{0, 0, [][]float64{}}, fmt.Errorf("could not multiply the matrices")
	}

	// initialize the matrix
	matrix := Matrix{matrix1.N, matrix2.M, make([][]float64, matrix1.N)}
	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i] = make([]float64, matrix2.M)
	}

	// need to test if this is the cache efficient version of matrix multiplication
	for i := 0; i < matrix1.N; i++ {
		for j := 0; j < matrix2.M; j++ {
			var sum float64
			for k := 0; k < matrix1.M; k++ {
				sum += matrix1.Matrix[i][k] * matrix2.Matrix[k][j]
			}
			matrix.Matrix[i][j] = sum
		}
	}

	return matrix, nil
}
