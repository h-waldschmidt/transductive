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

	matrix := NewMatrix(pointsY.N, pointsX.N)

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

	// initializing the vector(Matrix with N=1)
	vector := NewMatrix(1, pointsX.N)

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

// also known as 2-Norm
func EuclideanNorm(x []float64) float64 {

	var norm float64
	for i := 0; i < len(x); i++ {
		norm += math.Pow(x[i], 2)
	}

	return math.Sqrt(norm)
}

// also known as 1-Norm
func SumNorm(x []float64) float64 {
	var norm float64
	for i := 0; i < len(x); i++ {
		norm += x[i]
	}

	return norm
}

// convert slice to diagonal Matrix
func SliceToMatrix(x []float64) Matrix {
	// not using the constructor for efficiency
	matrix := Matrix{len(x), len(x), make([][]float64, len(x))}
	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i] = make([]float64, matrix.M)
		matrix.Matrix[i][i] = x[i]
	}

	return matrix
}

// convert diagonal Matrix to slice
func (matrix Matrix) MatrixToSlice() ([]float64, error) {

	// n and m have to be the same dimensions
	if matrix.N != matrix.M {
		return nil, fmt.Errorf("dimensions of Matrix are not same")
	}

	slice := make([]float64, matrix.N)
	for i := 0; i < matrix.N; i++ {
		slice[i] = matrix.Matrix[i][i]
	}

	return slice, nil
}

func MatrixMultiplication(matrix1 Matrix, matrix2 Matrix) (Matrix, error) {

	// The inner dimensions need to be the same
	if matrix1.N != matrix2.M {
		return Matrix{0, 0, [][]float64{}}, fmt.Errorf("could not multiply the matrices")
	}

	matrix := NewMatrix(matrix1.M, matrix2.N)

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
	transpose := NewMatrix(matrix.M, matrix.N)

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

	matrix := NewMatrix(matrix1.N, matrix1.M)

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

	matrix := NewMatrix(matrix1.N, matrix1.M)

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
// Explanation can be found here: https://de.wikipedia.org/wiki/QR-Algorithmus#Einfache_QR-Iteration
func (matrix Matrix) CalculateEigen() (Eigen, error) {
	var eigen Eigen
	if matrix.N != matrix.M {
		return eigen, fmt.Errorf("given matrix is not quadratic")
	}

	q, r := matrix.qrDecomposition()
	a_i := matrix
	q_i := q
	var previous Matrix
	var err error
	// QR-Algorithm
	for i := 0; i < 500; i++ {
		previous = a_i
		a_i, err = MatrixMultiplication(r, q)
		if err != nil {
			log.Fatal(err)
		}
		q, r = a_i.qrDecomposition()

		q_i, err = MatrixMultiplication(q_i, q)
		if err != nil {
			log.Fatal(err)
		}

		// same tolerance used as numpy
		tolerance := 1e-08

		equal, err := compAllClose(a_i, previous, tolerance)
		if err != nil {
			log.Fatal(err)
		}
		if equal {
			break
		}
	}

	// convert a_i and q_i into eigen datastructure
	eigen.Vectors = make([]Matrix, q_i.N)
	cache := NewMatrix(1, q_i.M)
	for i := 0; i < q_i.N; i++ {
		cache.Matrix[0] = q_i.Matrix[i]
		eigen.Vectors[i] = cache
	}
	eigen.Values = make([]float64, a_i.N)
	for i := 0; i < a_i.N; i++ {
		eigen.Values[i] = a_i.Matrix[i][i]
	}

	return eigen, nil
}

// calculating the QR-Decomposition using the Householder Transformation
// Explanation can be found here: https://en.wikipedia.org/wiki/QR_decomposition#Using_Householder_reflections
func (matrix Matrix) qrDecomposition() (Matrix, Matrix) {
	//initialize all needed variables
	var q, r Matrix
	x := NewMatrix(1, matrix.M)
	e := NewMatrix(1, matrix.N)

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

// Householder Transformation
// Explanation can be found here: https://de.wikipedia.org/wiki/Householdertransformation
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

// helper function for QR-Composition
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

// calculates the multiplicative Inverse of the matrix,
// using the Gauss Jordan Algorithm
//
// to keep the complexity in check, this function can only
// be performed on symmetric matrices
//
// implementation based on this article: https://www.codesansar.com/numerical-methods/python-program-inverse-matrix-using-gauss-jordan.htm
func (matrix Matrix) Inverse() (Matrix, error) {
	var inverse Matrix
	if matrix.N != matrix.M {
		return inverse, fmt.Errorf("given matrix is not quadratic")
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
	inverse = Matrix{matrix.N, matrix.M, make([][]float64, matrix.N)}
	for i := 0; i < inverse.N; i++ {
		inverse.Matrix[i] = matrix.Matrix[matrix.N+i]
	}

	return inverse, nil
}

// since the inverse of a diagonal matrix can easily be computed
// by inverting each entry, this function can be used for efficiency
func (matrix Matrix) InverseDiagonal() (Matrix, error) {
	var inverse Matrix
	if matrix.N != matrix.M {
		return inverse, fmt.Errorf("dimensions of Matrix are not same")
	}

	inverse = NewMatrix(matrix.N, matrix.M)
	// not using the constructor for efficiency
	inverse = Matrix{matrix.N, matrix.M, make([][]float64, matrix.N)}
	for i := 0; i < inverse.N; i++ {
		inverse.Matrix[i] = make([]float64, inverse.M)
		inverse.Matrix[i][i] = 1 / matrix.Matrix[i][i]
	}

	return inverse, nil
}
