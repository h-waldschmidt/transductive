package lialg

import (
	"log"
	"math"
)

type Eigen struct {
	Values  []float64
	Vectors []Matrix
}

// Calculate Eigen uses the QR Algorithm to calculate the eigenvalues (Schur Factorization)
// and the Eigenvectors by solving (A - eigenValue * I) * x = 0 for each eigenvalue
func (matrix *Matrix) CalculateEigen() Eigen {
	var eigen Eigen
	if matrix.N != matrix.M {
		log.Fatalf("matrix has to be quadratic")
	}

	// copy matrix data
	ak := *NewMatrix(matrix.N, matrix.N)
	for i := 0; i < matrix.N; i++ {
		for j := 0; j < matrix.M; j++ {
			ak.Matrix[i][j] = matrix.Matrix[i][j]
		}
	}
	qq := *NewMatrix(matrix.M, matrix.M)
	qq.AddDiagonal(1)

	// same tolerance used as numpy
	tolerance := 1e-08

	for i := 0; i < 500; i++ {
		previous := ak
		s := ak.Matrix[ak.N-1][ak.N-1]
		ak.SubDiagonal(s)
		q, r := ak.QrDecomposition()
		ak = MatrixMultiplication(r, q)
		ak.AddDiagonal(s)
		qq = MatrixMultiplication(qq, q)

		equal := CompDiagonalClose(ak, previous, tolerance)
		if equal {
			break
		}
	}

	eigen.Vectors = make([]Matrix, ak.N)

	// solve (A - eigenValue * I) * x = 0 for each eigenvalue
	for i := 0; i < matrix.N; i++ {
		var eigenVector Matrix

		// test all columns for eigen vectors
		// basically sets one each variable as one and then tests
		for k := 0; k < matrix.N; k++ {

			// copy matrix data
			curMatrix := *NewMatrix(matrix.N, matrix.N)
			for l := 0; l < matrix.N; l++ {
				for j := 0; j < matrix.M; j++ {
					curMatrix.Matrix[l][j] = matrix.Matrix[l][j]
				}
			}

			// subtract current eigen value
			curMatrix.SubDiagonal(ak.Matrix[i][i])

			zero := *NewMatrix(1, curMatrix.M)

			// remove first column of r
			for j := 0; j < zero.M; j++ {
				zero.Matrix[0][j] = -curMatrix.Matrix[k][j]
			}
			curMatrix.Matrix = append(curMatrix.Matrix[:k], curMatrix.Matrix[k+1:]...)
			curMatrix.N--

			// solve linear system
			q, r := curMatrix.QrDecomposition()
			q = q.TransposeMatrix()
			zero = MatrixMultiplication(q, zero)
			eigenVector = BackwardsSubstitution(&r, &zero)

			eigenVector.Matrix[0] = append(eigenVector.Matrix[0], 0)
			copy(eigenVector.Matrix[0][k+1:], eigenVector.Matrix[0][k:])
			eigenVector.Matrix[0][k] = 1
			eigenVector.M++

			// weird bug with huge vectors and cosine similarity, that says that vectors are similar
			// just skip those vectors, because probably not right
			// TODO: find fix for this
			if EuclideanNorm(eigenVector.Matrix[0]) > 1e12 {
				continue
			}

			cache := MatrixMultiplication(*matrix, eigenVector)
			cosineSimilarity := CosineSimilarity(eigenVector, cache)
			if math.Abs(cosineSimilarity)-1 < tolerance {
				eigen.Vectors[i] = eigenVector
				break
			}
		}
	}

	eigen.Values = make([]float64, ak.N)
	for i := 0; i < matrix.N; i++ {
		eigen.Values[i] = ak.Matrix[i][i]
	}
	return eigen
}

func BackwardsSubstitution(matrix *Matrix, vector *Matrix) Matrix {
	if matrix.M != vector.M || vector.N != 1 || matrix.N > matrix.M {
		log.Fatal("Either Dimensions do not match or the given vector has not only one dimension")
	}

	solution := NewMatrix(1, matrix.N)
	for i := matrix.N - 1; i >= 0; i-- {
		solution.Matrix[0][i] = vector.Matrix[0][i]
		for j := i + 1; j < matrix.N; j++ {
			solution.Matrix[0][i] -= (matrix.Matrix[j][i] * solution.Matrix[0][j])
		}
		solution.Matrix[0][i] /= matrix.Matrix[i][i]
	}

	return *solution
}

// calculating the QR-Decomposition using the Householder Transformation
//
// Explanation can be found here: https://en.wikipedia.org/wiki/QR_decomposition#Using_Householder_reflections
func (matrix *Matrix) QrDecomposition() (Matrix, Matrix) {
	//initialize q,r matrix and x,e vectors
	var q Matrix
	r := *matrix

	for i := 0; i < matrix.N-1; i++ {
		x := *NewMatrix(1, matrix.M-i)
		e := *NewMatrix(1, matrix.M-i)

		for j := 0; j < e.M; j++ {
			x.Matrix[0][j] = r.Matrix[i][j+i]
		}
		e.Matrix[0][0] = 1

		alpha := EuclideanNorm(x.Matrix[0])
		if x.Matrix[0][0] >= 0 {
			alpha *= -1
		}

		// e should be ith vector of identity matrix
		for j := 0; j < e.M; j++ {
			e.Matrix[0][j] = x.Matrix[0][j] + alpha*e.Matrix[0][j]
		}
		norm := EuclideanNorm(e.Matrix[0])
		cache := e.MatrixScalarMultiplication(1 / norm)
		qMin := cache.houseHolderTransformation(i)

		qT := qMin.calculateQT(i)
		if i == 0 {
			q = qT
			r = MatrixMultiplication(qT, *matrix)
		} else {
			q = MatrixMultiplication(qT, q)
			r = MatrixMultiplication(qT, r)
		}
	}
	return q.TransposeMatrix(), r
}

// Householder Transformation
//
// Explanation can be found here: https://de.wikipedia.org/wiki/Householdertransformation
func (vector *Matrix) houseHolderTransformation(k int) Matrix {
	if vector.N != 1 {
		log.Fatal("operation can only be performed on vector")
	}

	matrix := NewMatrix(vector.M, vector.M)
	for i := 0; i < matrix.M; i++ {
		for j := 0; j < matrix.N; j++ {
			matrix.Matrix[j][i] = -2 * vector.Matrix[0][j] * vector.Matrix[0][i]
			if i == j {
				matrix.Matrix[j][i] += 1
			}
		}
	}
	return *matrix
}

// helper function for QR-Decomposition
func (matrix *Matrix) calculateQT(k int) Matrix {
	if matrix.N != matrix.M {
		log.Fatal("given matrix is not quadratic")
	}
	qT := NewMatrix(matrix.N+k, matrix.M+k)
	for i := 0; i < qT.N; i++ {
		for j := 0; j < qT.M; j++ {
			if i < k || j < k {
				if i == j {
					qT.Matrix[i][j] = 1
				} else {
					qT.Matrix[i][j] = 0
				}
			} else {
				qT.Matrix[i][j] = matrix.Matrix[i-k][j-k]
			}
		}
	}
	return *qT
}
