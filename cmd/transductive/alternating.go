package transductive

import (
	"log"
	"transductive-experimental-design/cmd/datamanager"
)

func AlternatingOptimization(points datamanager.Matrix, numOfSelectedPoints int, lambda float64, sigma float64) datamanager.Matrix {

	// create K = V * V^T matrix with V being the point Matrix
	points_T := points.TransposeMatrix()
	k, err := datamanager.MatrixMultiplication(points, points_T)
	if err != nil {
		log.Fatal(err)
	}

	// eigen components of K matrix
	eigen, err := k.CalculateEigen()
	if err != nil {
		log.Fatal(err)
	}

	// create K*K matrix
	kk, err := datamanager.MatrixMultiplication(k, k)
	if err != nil {
		log.Fatal(err)
	}

	// create all the (K*K + lambda* eigen_value_i)^-1 matrices
	// those are needed to find all the alpha_i
	kk_slice := make([]datamanager.Matrix, len(eigen.Values))
	for i := 0; i < len(eigen.Values); i++ {
		kk_slice[i] = kk
		for j := 0; j < kk_slice[i].M; j++ {
			kk_slice[i].Matrix[j][j] += lambda * eigen.Values[i]
		}
		kk_slice[i], err = kk_slice[i].Inverse()
		if err != nil {
			log.Fatal(err)
		}
	}
	//repeat until no major improvement
	// for testing purposes I'm running the algorithm a fixed time
	// find optimal alpha
	// find optimal beta
	// normalize Beta Matrix

	// extract selected Points from Beta Matrix,
	// by selecting the numOfSelectedPoints biggest points
}

func findAlpha() {}

func findBeta() {}

func normalizeBetaMatrix() {}
