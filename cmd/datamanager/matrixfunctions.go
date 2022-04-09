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

func (matrix Matrix) TransposeMatrix() Matrix {
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

func (matrix Matrix) MatrixScalarMultiplication(scalar float64) Matrix {

	for i := 0; i < matrix.N; i++ {
		for j := 0; j < matrix.M; j++ {
			matrix.Matrix[i][j] *= scalar
		}
	}
	return matrix
}

// CalculateEigen uses the QR Algorithm to calculate the eigenvalues and eigenvectors
//Guide can be found here: https://de.wikipedia.org/wiki/QR-Algorithmus#Einfache_QR-Iteration
func (matrix Matrix) CalculateEigen() (Eigen, error) {
	var eigen Eigen
	if matrix.N != matrix.M {
		return eigen, fmt.Errorf("given matrix is not quadratic")
	}

	q, r := matrix.qrDecomposition()
	x := matrix
	var previous Matrix
	var err error
	for i := 0; i < 500; i++ {
		previous = q
		x, err = MatrixMultiplication(r, q)
		if err != nil {
			log.Fatal(err)
		}
		q, r = x.qrDecomposition()
	}
	return eigen, nil
}

// calculating the QR-Decomposition using the Householder Transformation
// Guide can be found here: https://en.wikipedia.org/wiki/QR_decomposition#Using_Householder_reflections
func (matrix Matrix) qrDecomposition() (Matrix, Matrix) {
	//initialize all needed variables
	var q, r Matrix
	x := Matrix{1, matrix.M, make([][]float64, 1)}
	e := Matrix{1, matrix.N, make([][]float64, 1)}

	for i := 0; i < matrix.N; i++ {
		x.Matrix[0] = matrix.Matrix[i]
		e.Matrix[0] = matrix.Matrix[i]

		alpha := EuclideanNorm(x.Matrix[0])
		if matrix.Matrix[i][i] > 0 {
			alpha *= -1
		}

		for j := 0; j < e.M; j++ {
			if i == j {
				e.Matrix[0][j] = 1
			} else {
				e.Matrix[0][j] = 0
			}
		}

		for j := 0; j < e.M; j++ {
			e.Matrix[0][j] = x.Matrix[0][j] + alpha*e.Matrix[0][j]
		}
		norm := EuclideanNorm(e.Matrix[0])
		e = e.MatrixScalarMultiplication(norm)
		q_min, err := e.houseHolderTransformation()
		if err != nil {
			log.Fatal(err)
		}

		q_t, err := q_min.calculateQ_T(i)
		if err != nil {
			log.Fatal(err)
		}
		if i == 0 {
			q = q_t
			r, err = MatrixMultiplication(q_t, matrix)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			q, err = MatrixMultiplication(q_t, q)
			if err != nil {
				log.Fatal(err)
			}

			r, err = MatrixMultiplication(q_t, r)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return q, r
}

func (vector Matrix) houseHolderTransformation() (Matrix, error) {
	var matrix Matrix
	if vector.N != 1 {
		return matrix, fmt.Errorf("operation can only be performed on vector (matrix.N == 1)")
	}
	var err error
	vector_t := vector.TransposeMatrix()
	matrix, err = MatrixMultiplication(vector, vector_t)
	if err != nil {
		log.Fatal(err)
	}

	matrix = matrix.MatrixScalarMultiplication(2)
	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i][i]++
	}
	return matrix, nil
}

func (matrix Matrix) calculateQ_T(k int) (Matrix, error) {
	var q_t Matrix
	if matrix.N != matrix.M {
		return q_t, fmt.Errorf("given matrix is not quadratic")
	}
	q_t = matrix
	for i := 0; i < matrix.N; i++ {
		for j := 0; j < matrix.M; j++ {
			if i < k || j < k {
				if i == j {
					q_t.Matrix[i][j] = 1
				} else {
					q_t.Matrix[i][j] = 0
				}
			} else {
				q_t.Matrix[i][j] = matrix.Matrix[i-k][j-k]
			}
		}
	}
	return q_t, nil
}

// tests if two matrices are the same within the given tolerance
// similar to this numpy function: https://numpy.org/doc/stable/reference/generated/numpy.allclose.html
func compAllClose(matrix1 Matrix, matrix2 Matrix, tolerance float64) (bool, error) {
	if matrix1.N != matrix2.N || matrix1.M != matrix2.M {
		return false, fmt.Errorf("can't compare matrices, because they don't have the same dimensions")
	}

	var difference float64
	for i := 0; i < matrix1.N; i++ {
		for j := 0; j < matrix1.M; j++ {
			difference = math.Abs(matrix1.Matrix[i][j] - matrix2.Matrix[i][j])
			if difference > tolerance {
				return false, nil
			}
		}
	}
	return true, nil
}
