package lialg

import "log"

type Eigen struct {
	Values  []float64
	Vectors []Matrix
}

// CalculateEigen uses the QR Algorithm to calculate the eigenvalues and eigenvectors
//
// Explanation can be found here: https://towardsdatascience.com/eigenvalues-and-eigenvectors-89483fb56d56
//
// Algorithm produces only correct eigenvectors for normal matrices (https://en.wikipedia.org/wiki/Normal_matrix)
func (matrix *Matrix) CalculateEigen() Eigen {
	var eigen Eigen
	if matrix.N != matrix.M {
		log.Fatalf("matrix has to be quadratic")
	}

	q, _ := matrix.QrDecomposition()
	e := MatrixMultiplication(q.TransposeMatrix(), *matrix)
	e = MatrixMultiplication(e, q)
	u := q
	for i := 0; i < 500; i++ {
		previous := e
		q, _ = e.QrDecomposition()
		e = MatrixMultiplication(q.TransposeMatrix(), e)
		e = MatrixMultiplication(e, q)
		u = MatrixMultiplication(u, q)
		// same tolerance used as numpy
		tolerance := 1e-08

		equal := CompDiagonalClose(e, previous, tolerance)
		if equal {
			break
		}
	}
	eigen.Vectors = make([]Matrix, u.N)

	for i := 0; i < u.N; i++ {
		cache := NewMatrix(1, u.M)
		cache.Matrix[0] = u.Matrix[i]
		eigen.Vectors[i] = *cache
	}
	eigen.Values = make([]float64, e.N)
	for i := 0; i < e.N; i++ {
		eigen.Values[i] = e.Matrix[i][i]
	}
	return eigen
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
		e := *NewMatrix(1, matrix.N-i)

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
