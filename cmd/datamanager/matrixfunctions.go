package datamanager

import (
	"fmt"
	"log"
	"math"
)

func CalculateKernelMatrix(pointsX Matrix, pointsY Matrix, sigma float64) (Matrix, error) {

	//x and y need to have the same dimensions
	if pointsX.M != pointsY.M {
		return Matrix{0, 0, make([][]float64, 0)}, fmt.Errorf("could not use RBFKernel")
	}

	// initializing the matrix
	matrix := Matrix{pointsY.N, pointsX.N, make([][]float64, pointsY.N)}
	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i] = make([]float64, pointsX.N)
	}

	// calculating all the values
	for i := 0; i < matrix.N; i++ {
		for j := 0; j < matrix.M; j++ {
			var err error
			matrix.Matrix[i][j], err = RbfKernel(pointsX.Matrix[i], pointsY.Matrix[j], sigma)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return matrix, nil
}

func CalculateKernelVector(pointsX Matrix, point []float64, sigma float64) (Matrix, error) {

	//x and y need to have the same dimensions
	if pointsX.M != len(point) {
		return Matrix{0, 0, make([][]float64, 0)}, fmt.Errorf("could not use RBFKernel")
	}

	// initializing the vector(Matrix with M=1)
	vector := Matrix{1, pointsX.N, make([][]float64, 1)}
	vector.Matrix[0] = make([]float64, pointsX.N)

	// calculating all the values
	for i := 0; i < pointsX.N; i++ {
		var err error
		vector.Matrix[0][i], err = RbfKernel(pointsX.Matrix[i], point, sigma)
		if err != nil {
			log.Fatal(err)
		}
	}

	return vector, nil
}

func EuclideanDistance(x []float64, y []float64) (float64, error) {

	//x and y need to be vectors and have the same dimensions
	if len(x) != len(y) {
		return 0, fmt.Errorf("could not calculate euclidean Distance")
	}

	var distance float64
	for i := 0; i < len(x); i++ {
		distance += math.Pow(x[i]-y[i], 2)
	}

	return math.Sqrt(distance), nil
}

func EuclideanNorm(x []float64) float64 {

	var norm float64
	for i := 0; i < len(x); i++ {
		norm += math.Pow(x[i], 2)
	}

	return math.Sqrt(norm)
}

func MatrixMultiplication(matrix1 Matrix, matrix2 Matrix) (Matrix, error) {

	// The inner dimensions need to be the same
	if matrix1.N != matrix2.M {
		return Matrix{0, 0, [][]float64{}}, fmt.Errorf("could not multiply the matrices")
	}

	// initialize the matrix
	matrix := Matrix{matrix1.M, matrix2.N, make([][]float64, matrix1.M)}
	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i] = make([]float64, matrix.M)
	}

	// need to test if this is the cache efficient version of matrix multiplication
	for i := 0; i < matrix1.M; i++ {
		for j := 0; j < matrix2.N; j++ {
			var sum float64
			for k := 0; k < matrix1.N; k++ {
				sum += matrix1.Matrix[k][i] * matrix2.Matrix[j][k]
			}
			matrix.Matrix[i][j] = sum
		}
	}

	return matrix, nil
}

func TransposeMatrix(matrix Matrix) Matrix {
	//initialize the transpose matrix
	transpose := Matrix{matrix.M, matrix.N, make([][]float64, matrix.M)}
	for i := 0; i < transpose.N; i++ {
		transpose.Matrix[i] = make([]float64, transpose.M)
	}

	for i := 0; i < transpose.N; i++ {
		for j := 0; j < transpose.M; j++ {
			transpose.Matrix[i][j] = matrix.Matrix[j][i]
		}
	}

	return transpose
}

func MatrixAddition(matrix1 Matrix, matrix2 Matrix) (Matrix, error) {

	// the dimensions of the matrices have to match
	if matrix1.N != matrix2.N || matrix1.M != matrix2.M {
		return Matrix{0, 0, [][]float64{}}, fmt.Errorf("could not add the matrices")
	}

	//initialize the matrix
	matrix := Matrix{matrix1.N, matrix1.M, make([][]float64, matrix1.N)}
	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i] = make([]float64, matrix.M)
	}

	for i := 0; i < matrix.N; i++ {
		for j := 0; j < matrix.M; j++ {
			matrix.Matrix[i][j] = matrix1.Matrix[i][j] + matrix2.Matrix[i][j]
		}
	}

	return matrix, nil
}

func MatrixSubtraction(matrix1 Matrix, matrix2 Matrix) (Matrix, error) {
	// the dimensions of the matrices have to match
	if matrix1.N != matrix2.N || matrix1.M != matrix2.M {
		return Matrix{0, 0, [][]float64{}}, fmt.Errorf("could not add the matrices")
	}

	//initialize the matrix
	matrix := Matrix{matrix1.N, matrix1.M, make([][]float64, matrix1.N)}
	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i] = make([]float64, matrix.M)
	}

	for i := 0; i < matrix.N; i++ {
		for j := 0; j < matrix.M; j++ {
			matrix.Matrix[i][j] = matrix1.Matrix[i][j] - matrix2.Matrix[i][j]
		}
	}

	return matrix, nil
}

func MatrixScalarMultiplication(matrix Matrix, scalar float64) Matrix {

	for i := 0; i < matrix.N; i++ {
		for j := 0; j < matrix.M; j++ {
			matrix.Matrix[i][j] *= scalar
		}
	}
	return matrix
}
