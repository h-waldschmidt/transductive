package lialg

import (
	"log"
	"math"
	"sync"
)

// N represents the number of columns (so that a vector can be in one array)
//
// M represents the number rows (e.g. dimension of a vector)
type Matrix struct {
	N, M   int
	Matrix [][]float64
}

// constructor for Matrix
//
// makes it easier to create the Matrix array
func NewMatrix(n, m int) *Matrix {
	if n < 0 || m < 0 {
		log.Fatal("dimensions of matrix must be greater than 0")
	}

	matrix := &Matrix{n, m, make([][]float64, n)}
	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i] = make([]float64, matrix.M)
	}
	return matrix
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
//
// also known as 2-Norm
func EuclideanNorm(x []float64) float64 {
	var norm float64
	for i := 0; i < len(x); i++ {
		norm += math.Pow(x[i], 2)
	}
	return math.Sqrt(norm)
}

// Calculates Sum Norm of Vector
//
// also known as 1-Norm
func SumNorm(x []float64) float64 {
	var norm float64
	for i := 0; i < len(x); i++ {
		norm += math.Abs(x[i])
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
	for i := 0; i < matrix.M; i++ {
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
	var wg sync.WaitGroup

	// need to test if this is the cache efficient version of matrix multiplication
	for i := 0; i < matrix1.M; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < matrix2.N; j++ {
				for k := 0; k < matrix1.N; k++ {
					matrix.Matrix[j][i] += matrix1.Matrix[k][i] * matrix2.Matrix[j][k]
				}
			}
		}(i)
	}
	wg.Wait()
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

// adds a component to each diagonal element of matrix
//
// ensure that matrix is quadratic
func (matrix *Matrix) AddDiagonal(a float64) {
	if matrix.N != matrix.M {
		log.Fatalf("matrix has to be quadratic")
	}

	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i][i] += a
	}
}

// subtracts a component to each diagonal element of matrix
//
// ensure that matrix is quadratic
func (matrix *Matrix) SubDiagonal(a float64) {
	matrix.AddDiagonal(-a)
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

// tests if two matrices are the same within the given tolerance
//
// similar to this numpy function: https://numpy.org/doc/stable/reference/generated/numpy.allclose.html
func CompAllClose(matrix1, matrix2 Matrix, tolerance float64) bool {
	if matrix1.N != matrix2.N || matrix1.M != matrix2.M {
		log.Fatal("matrices do not have the same dimensions")
	}

	if tolerance < 0 {
		log.Fatal("tolerance must be positive")
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

// tests if the diagonal of two quadratic matrices is almost same within given tolerance
func CompDiagonalClose(matrix1, matrix2 Matrix, tolerance float64) bool {
	if matrix1.N != matrix2.N || matrix1.M != matrix2.M {
		log.Fatal("matrices do not have the same dimensions")
	}
	if matrix1.N != matrix2.N {
		log.Fatal("matrices must be quadratic")
	}
	if tolerance < 0 {
		log.Fatal("tolerance must be positive")
	}

	var difference float64
	for i := 0; i < matrix1.N; i++ {
		difference = math.Abs(matrix1.Matrix[i][i] - matrix2.Matrix[i][i])
		if difference > tolerance {
			return false
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
//
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

		value := math.Round(matrix.Matrix[i][i]*1e8) / 1e8
		if value != 0 {
			inverse.Matrix[i][i] = 1 / value
		} else {
			inverse.Matrix[i][i] = 0
		}
	}
	return inverse
}
