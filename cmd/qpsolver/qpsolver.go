package qpsolver

import (
	"log"
	"transductive-experimental-design/cmd/datamanager"

	"github.com/badgerodon/lalg"
	"github.com/badgerodon/quadprog"
)

// Use the following library https://github.com/badgerodon/quadprog as current solution
// might write my own byndings to https://doc.cgal.org/latest/QP_solver/index.html
func Solve(D, d datamanager.Matrix) datamanager.Matrix {
	// convert Matrices from this package into lalg Matrices and Vectors
	lD := convertDMatrixToLMatrix(D)
	ld := convertDMatrixToLVector(d)
	lA1 := lalg.NewIdentity(0)
	lb1 := lalg.NewVector(0)
	lA2 := lalg.NewIdentity(D.N)
	lb2 := lalg.NewVector(D.N)

	lAns, err := quadprog.Solve(lD, ld, lA1, lb1, lA2, lb2)
	if err != nil {
		log.Fatal(err)
	}

	dAns := convertLVectorToDMatrix(lAns)

	return dAns
}

// convert datamanager.Matrix tpye to lalg.Matrix type
func convertDMatrixToLMatrix(dMatrix datamanager.Matrix) lalg.Matrix {
	lMatrix := lalg.NewMatrix(dMatrix.M, dMatrix.N)

	for i := 0; i < lMatrix.Rows; i++ {
		for j := 0; j < lMatrix.Cols; j++ {
			lMatrix.Set(i, j, dMatrix.Matrix[j][i])
		}
	}

	return lMatrix
}

// convert datamanager.Matrix (with N=1) to lalg.Vector type
func convertDMatrixToLVector(dVector datamanager.Matrix) lalg.Vector {
	if dVector.N != 1 {
		log.Fatal("N of dVector has to be 1")
	}

	lVector := lalg.NewVector(dVector.M)

	for i := 0; i < len(lVector); i++ {
		lVector[i] = dVector.Matrix[0][i]
	}

	return lVector
}

// convert lalg.Matrix type to datamanager.Matrix type
func convertLMatrixToDMatrix(lMatrix lalg.Matrix) datamanager.Matrix {
	dMatrix := datamanager.NewMatrix(lMatrix.Cols, lMatrix.Rows)

	for i := 0; i < dMatrix.N; i++ {
		for j := 0; j < dMatrix.M; j++ {
			dMatrix.Matrix[i][j] = lMatrix.Get(j, i)
		}
	}

	return *dMatrix
}

// convert lalg.Vector type to datamanager.Matrix (with N=1) type
func convertLVectorToDMatrix(lVector lalg.Vector) datamanager.Matrix {
	dVector := datamanager.NewMatrix(1, len(lVector))
	dVector.Matrix[0] = lVector

	return *dVector
}
