package lialg

import (
	"log"

	"github.com/badgerodon/lalg"
	"github.com/badgerodon/quadprog"
)

// Use the following library https://github.com/badgerodon/quadprog as current solution
//
// might write my own byndings to https://doc.cgal.org/latest/QP_solver/index.html
func QPSolve(D, d Matrix) Matrix {
	// convert Matrices from this package into lalg Matrices and Vectors
	lD := convertDMatrixToLMatrix(D)
	ld := convertDMatrixToLVector(d)

	// equality constraints for beta
	lA1 := lalg.NewMatrix(1, D.N)
	for i := 0; i < D.N; i++ {
		lA1.Set(0, i, 1.0)
	}
	lb1 := lalg.NewVector(1)
	lb1[0] = 1

	lA2 := lalg.NewIdentity(D.N)
	lb2 := lalg.NewVector(D.N)

	lAns, err := quadprog.Solve(lD, ld, lA1, lb1, lA2, lb2)
	if err != nil {
		log.Fatal(err)
	}

	dAns := convertLVectorToDMatrix(lAns)

	return dAns
}

// convert internal Matrix type to lalg Matrix type
func convertDMatrixToLMatrix(dMatrix Matrix) lalg.Matrix {
	lMatrix := lalg.NewMatrix(dMatrix.M, dMatrix.N)

	for i := 0; i < lMatrix.Rows; i++ {
		for j := 0; j < lMatrix.Cols; j++ {
			lMatrix.Set(i, j, dMatrix.Matrix[j][i])
		}
	}
	return lMatrix
}

// convert internal Matrix type (with N=1) to lalg Vector type
func convertDMatrixToLVector(dVector Matrix) lalg.Vector {
	if dVector.N != 1 {
		log.Fatal("N of dVector has to be 1")
	}

	lVector := lalg.NewVector(dVector.M)
	for i := 0; i < len(lVector); i++ {
		lVector[i] = dVector.Matrix[0][i]
	}
	return lVector
}

// convert lalg Matrix type to internal Matrix type
func convertLMatrixToDMatrix(lMatrix lalg.Matrix) Matrix {
	dMatrix := NewMatrix(lMatrix.Cols, lMatrix.Rows)

	for i := 0; i < dMatrix.N; i++ {
		for j := 0; j < dMatrix.M; j++ {
			dMatrix.Matrix[i][j] = lMatrix.Get(j, i)
		}
	}
	return *dMatrix
}

// convert lalg Vector type to internal Matrix (with N=1) type
func convertLVectorToDMatrix(lVector lalg.Vector) Matrix {
	dVector := NewMatrix(1, len(lVector))
	dVector.Matrix[0] = lVector
	return *dVector
}
