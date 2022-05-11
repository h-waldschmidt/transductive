package qpsolver

import (
	"log"
	"transductive-experimental-design/cmd/datamanager"

	"github.com/badgerodon/lalg"
	"github.com/badgerodon/quadprog"
)

// Use the following library https://github.com/badgerodon/quadprog as current solution
// might write my own byndings to https://doc.cgal.org/latest/QP_solver/index.html
func SolveQP(D datamanager.Matrix, d datamanager.Matrix, A1 datamanager.Matrix, b1 datamanager.Matrix, A2 datamanager.Matrix, b2 datamanager.Matrix) datamanager.Matrix {
	// convert Matrices from this package into lalg Matrices and Vectors
	lD := convertDMatrixToLMatrix(D)
	ld := convertDMatrixToLVector(d)
	lA1 := convertDMatrixToLMatrix(A1)
	lb1 := convertDMatrixToLVector(b1)
	lA2 := convertDMatrixToLMatrix(A2)
	lb2 := convertDMatrixToLVector(b2)

	lAns, err := quadprog.Solve(lD, ld, lA1, lb1, lA2, lb2)
	if err != nil {
		log.Fatal(err)
	}

	dAns := convertLVectorToDMatrix(lAns)

	return dAns
}

// convert datamanager.Matrix tpye to lalg.Matrix type
func convertDMatrixToLMatrix(dMatrix datamanager.Matrix) lalg.Matrix {

}

// convert datamanager.Matrix (with N=1) to lalg.Vector type
func convertDMatrixToLVector(dVector datamanager.Matrix) lalg.Vector {

}

// convert lalg.Matrix type to datamanager.Matrix type
func convertLMatrixToDMatrix(lMatrix lalg.Matrix) datamanager.Matrix {

}

// convert lalg.Vector type to datamanager.Matrix (with N=1) type
func convertLVectorToDMatrix(lVector lalg.Vector) datamanager.Matrix {

}
