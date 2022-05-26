package transductive

import (
	"math"
	"sort"
	"transductive-experimental-design/cmd/datamanager"
	"transductive-experimental-design/cmd/qpsolver"
)

type ValueCoordinateTuple struct {
	value      float64
	coordinate []float64
}

// TODO: create global variables
func AlternatingOptimization(points datamanager.Matrix, numOfSelectedPoints int, lambda, sigma float64) datamanager.Matrix {

	// create K = V * V^T matrix with V being the point Matrix
	k := points.CalculateKernelMatrix(points, sigma)

	// eigen components of K matrix
	eigen := k.CalculateEigen()

	// create K*K matrix
	kk := datamanager.MatrixMultiplication(k, k)

	// create all the (K*K + lambda* eigen_value_i)^-1 matrices
	// those are needed to find all the alpha_i
	kk_slice := make([]datamanager.Matrix, len(eigen.Values))
	for i := 0; i < len(eigen.Values); i++ {
		kk_slice[i] = kk
		for j := 0; j < kk_slice[i].M; j++ {
			kk_slice[i].Matrix[j][j] += lambda * eigen.Values[i]
		}
		kk_slice[i] = kk_slice[i].Inverse()
	}

	// TODO: try to init beta with different methods and Values
	// initialize beta slice
	beta := datamanager.NewMatrix(1, points.N)
	for i := 0; i < points.N; i++ {
		beta.Matrix[0][i] = 0.1
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
		betaDiagonal := beta.VectorToDiagonalMatrix()
		findAlpha(*alphaMatrix, betaDiagonal, k, kk_slice, eigen.Vectors)

		// find optimal beta
		cache := findBeta(*beta, *alphaMatrix, kk, k, eigen, lambda, sigma)
		// normalize Beta Matrix
		cache = datamanager.ComponentWiseMultiplication(*beta, cache)
		beta = &cache
	}

	// create slice with value coordinate tuples
	tuples := make([]ValueCoordinateTuple, points.N)
	for i := 0; i < points.N; i++ {
		tuples[i] = ValueCoordinateTuple{beta.Matrix[0][i], points.Matrix[i]}
	}

	// sort it in descending order
	sort.Slice(tuples, func(i, j int) bool {
		return tuples[i].value > tuples[j].value
	})

	// extract the numOfSelectedPoints first values
	ans := datamanager.NewMatrix(numOfSelectedPoints, points.M)
	for i := 0; i < numOfSelectedPoints; i++ {
		ans.Matrix[i] = tuples[i].coordinate
	}
	return *ans
}

func findAlpha(alphaMatrix, betaDiagonal, kMatrix datamanager.Matrix, kkMatrices, eigenVectors []datamanager.Matrix) datamanager.Matrix {
	newAlphaMatrix := alphaMatrix
	for i := 0; i < alphaMatrix.N; i++ {
		cache := betaDiagonal.InverseDiagonal()

		cache = datamanager.MatrixMultiplication(cache, kkMatrices[i])

		cache = datamanager.MatrixMultiplication(cache, kMatrix)

		cache = datamanager.MatrixMultiplication(cache, eigenVectors[i])

		newAlphaMatrix.Matrix[i] = cache.Matrix[0]
	}

	return newAlphaMatrix
}

func findBeta(beta, alphaMatrix, kk, k datamanager.Matrix, eigen datamanager.Eigen, lambda, sigma float64) datamanager.Matrix {
	// prepare H matrix and f vector for qpsolver
	h := datamanager.NewMatrix(beta.M, beta.M)
	f := datamanager.NewMatrix(beta.M, 1)
	identity := datamanager.CreateIdentity(beta.M)

	for i := 0; i < len(eigen.Values); i++ {
		alphaDiagonal := eigen.Vectors[i].VectorToDiagonalMatrix()

		cacheH := identity.MatrixScalarMultiplication(lambda * eigen.Values[i])
		cacheH = datamanager.MatrixAddition(cacheH, kk)
		cacheH = datamanager.MatrixMultiplication(alphaDiagonal, cacheH)
		cacheH = datamanager.MatrixMultiplication(cacheH, alphaDiagonal)

		cacheH = datamanager.MatrixAddition(cacheH, *h)
		h = &cacheH

		cacheF := eigen.Vectors[i].TransposeMatrix()
		cacheF = cacheF.MatrixScalarMultiplication(math.Sqrt(eigen.Values[i]))
		cacheF = datamanager.MatrixMultiplication(cacheF, k)
		cacheF = datamanager.MatrixMultiplication(cacheF, alphaDiagonal)
		cacheF = datamanager.MatrixAddition(cacheF, *f)
		f = &cacheF
	}
	cacheF := f.MatrixScalarMultiplication(2)
	cache := datamanager.CreateAllOnesVector(beta.M)
	cache = cache.TransposeMatrix()
	cache = cache.MatrixScalarMultiplication(sigma)
	cache = datamanager.MatrixSubtraction(cache, cacheF)
	cache = cache.TransposeMatrix()
	f = &cache

	return qpsolver.Solve(*h, *f)
}
