package transductive

import (
	"log"
	"transductive-experimental-design/cmd/datamanager"
)

// TODO: create global variables

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

	// TODO: try to init beta with different methods and Values
	// initialize beta slice
	beta := make([]float64, points.N)
	for i := 0; i < points.N; i++ {
		beta[i] = 0.1
	}

	// TODO: try to init alphaMatrix with different methods and Values
	// initialize alpha matrix
	alphaMatrix := datamanager.NewMatrix(len(eigen.Values), points.N)
	for i := 0; i < alphaMatrix.N; i++ {
		for j := 0; j < alphaMatrix.M; j++ {
			alphaMatrix.Matrix[i][j] = 0.1
		}
	}

	//repeat until no major improvement
	for i := 0; i < 50; i++ {
		// for testing purposes I'm running the algorithm with fixed rounds

		// find optimal alpha
		betaDiagonal := datamanager.SliceToMatrix(beta)
		findAlpha(alphaMatrix, betaDiagonal, k, kk_slice, eigen.Vectors)

		// find optimal beta
		// normalize Beta Matrix
	}

	// extract selected Points from Beta Matrix,
	// by selecting the numOfSelectedPoints biggest points
}

func findAlpha(alphaMatrix datamanager.Matrix, betaDiagonal datamanager.Matrix, kMatrix datamanager.Matrix, kkMatrices []datamanager.Matrix, eigenVectors []datamanager.Matrix) datamanager.Matrix {
	newAlphaMatrix := alphaMatrix
	for i := 0; i < alphaMatrix.N; i++ {
		cache, err := betaDiagonal.InverseDiagonal()
		if err != nil {
			log.Fatal(err)
		}

		cache, err = datamanager.MatrixMultiplication(cache, kkMatrices[i])
		if err != nil {
			log.Fatal(err)
		}

		cache, err = datamanager.MatrixMultiplication(cache, kMatrix)
		if err != nil {
			log.Fatal(err)
		}

		cache, err = datamanager.MatrixMultiplication(cache, eigenVectors[i])
		if err != nil {
			log.Fatal(err)
		}

		newAlphaMatrix.Matrix[i] = cache.Matrix[0]
	}

	return newAlphaMatrix
}

func findBeta() {}

func normalizeBetaMatrix() {}
