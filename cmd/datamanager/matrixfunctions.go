package datamanager

import (
	"log"
	"math"
)

// Applies the Kernel function to every component of the matrix
func (pointsX *Matrix) CalculateKernelMatrix(pointsY Matrix, sigma float64) Matrix {

	//x and y need to have the same dimensions
	if pointsX.M != pointsY.M {
		log.Fatal("could not use RBFKernel, points do not have the same dimensions")
	}

	matrix := NewMatrix(pointsY.N, pointsX.N)

	// calculating all the values
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

	// initializing the vector(Matrix with N=1)
	vector := NewMatrix(1, pointsX.N)

	// calculating all the values
	for i := 0; i < pointsX.N; i++ {
		vector.Matrix[0][i] = RbfKernel(pointsX.Matrix[i], point, sigma)
	}

	return *vector
}

// Calculates the Euclidian Distance between two vectors
func EuclideanDistance(x, y []float64) float64 {

	//x and y need to be vectors and have the same dimensions
	if len(x) != len(y) {
		log.Fatal("dimensions of the points do not match")
	}

	var distance float64
	for i := 0; i < len(x); i++ {
		distance += math.Pow(x[i]-y[i], 2)
	}

	return math.Sqrt(distance)
}

// Calculates Euclidian Norm of Vector
// also known as 2-Norm
func EuclideanNorm(x []float64) float64 {

	var norm float64
	for i := 0; i < len(x); i++ {
		norm += math.Pow(x[i], 2)
	}

	return math.Sqrt(norm)
}

// Calculates Sum Norm of Vector
// also known as 1-Norm
func SumNorm(x []float64) float64 {
	var norm float64
	for i := 0; i < len(x); i++ {
		norm += x[i]
	}

	return norm
}

// convert slice to diagonal Matrix
func SliceToDiagonalMatrix(x []float64) Matrix {
	// not using the constructor for efficiency
	matrix := Matrix{len(x), len(x), make([][]float64, len(x))}
	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i] = make([]float64, matrix.M)
		matrix.Matrix[i][i] = x[i]
	}

	return matrix
}

// convert diagonal Matrix to slice
func (matrix *Matrix) DiagonalMatrixToSlice() []float64 {
	if matrix.N != matrix.M {
		log.Fatal("matrix has to be quadratic")
	}

	slice := make([]float64, matrix.N)
	for i := 0; i < matrix.N; i++ {
		slice[i] = matrix.Matrix[i][i]
	}

	return slice
}

// converts a Matrix with N=1 to a Diagonal Matrix
func (matrix *Matrix) VectorToDiagonalMatrix() Matrix {
	if matrix.N != 1 {
		log.Fatal("matrix needs to only consist of 1 column")
	}

	ans := NewMatrix(matrix.M, matrix.M)
	for i := 0; i < matrix.N; i++ {
		ans.Matrix[i][i] = matrix.Matrix[0][i]
	}
	return *ans
}

// Converts a Diagonal Matrix to a Matrix with N=1
func (matrix *Matrix) DiagonalMatrixToVector() Matrix {

	ans := NewMatrix(1, matrix.M)
	for i := 0; i < matrix.N; i++ {
		ans.Matrix[0][i] = matrix.Matrix[i][i]
	}
	return *ans
}

// creates the n x n identity matrix
func CreateIdentity(n int) Matrix {
	identity := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		identity.Matrix[i][i] = 1
	}
	return *identity
}

func CreateAllOnesVector(n int) Matrix {
	allOnes := NewMatrix(1, n)
	for i := 0; i < n; i++ {
		allOnes.Matrix[0][i] = 1
	}
	return *allOnes
}

// Basic Matrix Multiplication
func MatrixMultiplication(matrix1, matrix2 Matrix) Matrix {

	// The inner dimensions need to be the same
	if matrix1.N != matrix2.M {
		log.Fatal("inner dimensions of matrices do not match")
	}

	matrix := NewMatrix(matrix2.N, matrix1.M)

	// need to test if this is the cache efficient version of matrix multiplication
	for i := 0; i < matrix1.M; i++ {
		for j := 0; j < matrix2.N; j++ {
			for k := 0; k < matrix1.N; k++ {
				matrix.Matrix[j][i] += matrix1.Matrix[k][i] * matrix2.Matrix[j][k]
			}
		}
	}

	return *matrix
}

// Transpose the given matrix
func (matrix *Matrix) TransposeMatrix() Matrix {
	transpose := NewMatrix(matrix.M, matrix.N)

	for i := 0; i < transpose.N; i++ {
		for j := 0; j < transpose.M; j++ {
			transpose.Matrix[i][j] = matrix.Matrix[j][i]
		}
	}

	return *transpose
}

// add the matrices component wise
func MatrixAddition(matrix1, matrix2 Matrix) Matrix {

	// the dimensions of the matrices have to match
	if matrix1.N != matrix2.N || matrix1.M != matrix2.M {
		log.Fatal("dimensions of matrices do not match")
	}

	matrix := NewMatrix(matrix1.N, matrix1.M)

	for i := 0; i < matrix.N; i++ {
		for j := 0; j < matrix.M; j++ {
			matrix.Matrix[i][j] = matrix1.Matrix[i][j] + matrix2.Matrix[i][j]
		}
	}

	return *matrix
}

// subtract the matrices component wise
func MatrixSubtraction(matrix1, matrix2 Matrix) Matrix {
	// the dimensions of the matrices have to match
	if matrix1.N != matrix2.N || matrix1.M != matrix2.M {
		log.Fatal("dimensions of matrices do not match")
	}

	matrix := NewMatrix(matrix1.N, matrix1.M)

	for i := 0; i < matrix.N; i++ {
		for j := 0; j < matrix.M; j++ {
			matrix.Matrix[i][j] = matrix1.Matrix[i][j] - matrix2.Matrix[i][j]
		}
	}

	return *matrix
}

// Multiply all elements of the matrix with the scalar
func (matrix *Matrix) MatrixScalarMultiplication(scalar float64) Matrix {
	answer := *matrix
	for i := 0; i < answer.N; i++ {
		for j := 0; j < answer.M; j++ {
			answer.Matrix[i][j] *= scalar
		}
	}
	return answer
}

// Multiply each element of matrix1 with the coresponding element of matrix2
func ComponentWiseMultiplication(matrix1, matrix2 Matrix) Matrix {
	// matrix1 and matrix2 need to have same dimensions
	if matrix1.N != matrix2.N || matrix1.M != matrix2.M {
		log.Fatal("dimensions of the matrices are not the same")
	}

	for i := 0; i < matrix1.N; i++ {
		for j := 0; j < matrix1.M; j++ {
			matrix1.Matrix[i][j] = matrix1.Matrix[i][j] * matrix2.Matrix[i][j]
		}
	}

	return matrix1
}

// CalculateEigen uses the QR Algorithm to calculate the eigenvalues and eigenvectors
// Explanation can be found here: https://de.wikipedia.org/wiki/QR-Algorithmus#Einfache_QR-Iteration
func (matrix *Matrix) CalculateEigen() Eigen {
	var eigen Eigen
	if matrix.N != matrix.M {
		log.Fatalf("matrix has to be quadratic")
	}

	q, r := matrix.QrDecomposition()
	a_i := *matrix
	q_i := q
	var previous Matrix

	// QR-Algorithm
	for i := 0; i < 500; i++ {
		previous = a_i
		a_i = MatrixMultiplication(r, q)
		q, r = a_i.QrDecomposition()

		q_i = MatrixMultiplication(q_i, q)

		// same tolerance used as numpy
		tolerance := 1e-08

		equal := compAllClose(a_i, previous, tolerance)
		if equal {
			break
		}
	}

	// convert a_i and q_i into eigen datastructure
	eigen.Vectors = make([]Matrix, q_i.N)
	cache := NewMatrix(1, q_i.M)
	for i := 0; i < q_i.N; i++ {
		cache.Matrix[0] = q_i.Matrix[i]
		eigen.Vectors[i] = *cache
	}
	eigen.Values = make([]float64, a_i.N)
	for i := 0; i < a_i.N; i++ {
		eigen.Values[i] = a_i.Matrix[i][i]
	}

	return eigen
}

// calculating the QR-Decomposition using the Householder Transformation
// Explanation can be found here: https://en.wikipedia.org/wiki/QR_decomposition#Using_Householder_reflections
func (matrix *Matrix) QrDecomposition() (Matrix, Matrix) {
	//initialize q,r matrix and x,e vectors
	var q Matrix
	r := *matrix
	x := *NewMatrix(1, matrix.M)
	e := *NewMatrix(1, matrix.N)

	for i := 0; i < matrix.N; i++ {
		x.Matrix[0] = r.Matrix[i]

		alpha := EuclideanNorm(x.Matrix[0])
		// TODO
		if x.Matrix[0][i] > 0 {
			alpha *= -1
		}

		// e should be ith vector of identity matrix
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
		cache := e.MatrixScalarMultiplication(1 / norm)
		q_min := cache.houseHolderTransformation(i)

		q_t := q_min.calculateQ_T(i)
		if i == 0 {
			q = q_t
			r = MatrixMultiplication(q_t, *matrix)
		} else {
			q = MatrixMultiplication(q_t, q)

			r = MatrixMultiplication(q_t, r)
		}
	}
	return q.TransposeMatrix(), r
}

// Householder Transformation
// Explanation can be found here: https://de.wikipedia.org/wiki/Householdertransformation
func (vector *Matrix) houseHolderTransformation(k int) Matrix {
	if vector.N != 1 {
		log.Fatal("operation can only be performed on vector")
	}

	matrix := NewMatrix(vector.M-k, vector.M-k)
	for i := 0; i < matrix.N; i++ {
		for j := 0; j < matrix.M; j++ {
			matrix.Matrix[i][j] = -2 * vector.Matrix[0][i] * vector.Matrix[0][j]
			if i == j {
				matrix.Matrix[i][j]++
			}
		}
	}
	/**
	vector_t := vector.TransposeMatrix()
	matrix := MatrixMultiplication(*vector, vector_t)

	matrix.MatrixScalarMultiplication(-2)
	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i][i]++
	}
	*/
	return *matrix
}

// helper function for QR-Decomposition
func (matrix *Matrix) calculateQ_T(k int) Matrix {
	if matrix.N != matrix.M {
		log.Fatal("given matrix is not quadratic")
	}
	q_t := NewMatrix(matrix.N+k, matrix.M+k)
	for i := 0; i < q_t.N; i++ {
		for j := 0; j < q_t.M; j++ {
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
	return *q_t
}

// tests if two matrices are the same within the given tolerance
// similar to this numpy function: https://numpy.org/doc/stable/reference/generated/numpy.allclose.html
func compAllClose(matrix1, matrix2 Matrix, tolerance float64) bool {
	if matrix1.N != matrix2.N || matrix1.M != matrix2.M {
		log.Fatal("matrices do not have the same dimensions")
	}

	var difference float64
	for i := 0; i < matrix1.N; i++ {
		for j := 0; j < matrix1.M; j++ {
			difference = math.Abs(matrix1.Matrix[i][j] - matrix2.Matrix[i][j])
			if difference > tolerance {
				return false
			}
		}
	}
	return true
}

// calculates the multiplicative Inverse of the matrix,
// using the Gauss Jordan Algorithm
//
// to keep the complexity in check, this function can only
// be performed on symmetric matrices
//
// implementation based on this article: https://www.codesansar.com/numerical-methods/python-program-inverse-matrix-using-gauss-jordan.htm
func (matrix Matrix) Inverse() Matrix {
	if matrix.N != matrix.M {
		log.Fatal("given matrix is not quadratic")
	}

	// augment the identity matrix of Order n
	for i := 0; i < matrix.N; i++ {
		matrix.Matrix = append(matrix.Matrix, make([]float64, matrix.M))
		matrix.Matrix[matrix.N+i][i] = 1
	}

	// calculate the inverse
	for i := 0; i < matrix.N; i++ {
		for j := 0; j < matrix.M; j++ {
			if i != j {
				ratio := matrix.Matrix[i][j] / matrix.Matrix[i][i]

				for k := 0; k < matrix.N*2; k++ {
					matrix.Matrix[k][j] = matrix.Matrix[k][j] - ratio*matrix.Matrix[k][i]
				}
			}
		}
	}
	for i := 0; i < matrix.M; i++ {
		divisor := matrix.Matrix[i][i]
		for j := 0; j < matrix.N*2; j++ {
			matrix.Matrix[j][i] = matrix.Matrix[j][i] / divisor
		}
	}

	// extract the inverse and return it
	// not using the constructor for efficiency
	inverse := Matrix{matrix.N, matrix.M, make([][]float64, matrix.N)}
	for i := 0; i < inverse.N; i++ {
		inverse.Matrix[i] = matrix.Matrix[matrix.N+i]
	}

	return inverse
}

// Calculate the Inverse of a DiagonalMatrix
// since the inverse of a diagonal matrix can easily be computed
// by inverting each entry, this function can be used for efficiency
func (matrix *Matrix) InverseDiagonal() Matrix {
	if matrix.N != matrix.M {
		log.Fatal("dimensions of matrices do not match")
	}

	// not using the constructor for efficiency
	inverse := Matrix{matrix.N, matrix.M, make([][]float64, matrix.N)}
	for i := 0; i < inverse.N; i++ {
		inverse.Matrix[i] = make([]float64, inverse.M)
		inverse.Matrix[i][i] = 1 / matrix.Matrix[i][i]
	}

	return inverse
}
