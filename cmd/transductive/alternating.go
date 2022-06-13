package transductive

import (
	"math"
	"sort"
	"transductive-experimental-design/cmd/lialg"
)

type ValueCoordinateTuple struct {
	value      float64
	coordinate []float64
}

// TODO: create global variables
func AlternatingOptimization(points lialg.Matrix, numOfSelectedPoints int, lambda, sigma float64) lialg.Matrix {

	// create K = V * V^T matrix with V being the point Matrix
	k := lialg.MatrixMultiplication(points, points.TransposeMatrix())
	//k := points.CalculateKernelMatrix(points, sigma)

	// eigen components of K matrix
	eigen := k.CalculateEigen()

	// create K*K matrix
	kk := lialg.MatrixMultiplication(k, k)

	// create all the (K*K + lambda* eigen_value_i)^-1 matrices
	// those are needed to find all the alpha_i
	kk_slice := make([]lialg.Matrix, len(eigen.Values))
	for i := 0; i < len(eigen.Values); i++ {
		kk_slice[i] = kk
		for j := 0; j < kk_slice[i].M; j++ {
			kk_slice[i].Matrix[j][j] += lambda * eigen.Values[i]
		}
		kk_slice[i] = kk_slice[i].Inverse()
	}

	// TODO: try to init beta with different methods and Values
	// initialize beta slice
	beta := lialg.NewMatrix(1, points.N)
	for i := 0; i < points.N; i++ {
		beta.Matrix[0][i] = 0.1
	}

	// TODO: try to init alphaMatrix with different methods and Values
	// initialize alpha matrix
	alphaMatrix := lialg.NewMatrix(len(eigen.Values), points.N)
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
		cache = lialg.ComponentWiseMultiplication(*beta, cache)
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
	ans := lialg.NewMatrix(numOfSelectedPoints, points.M)
	for i := 0; i < numOfSelectedPoints; i++ {
		ans.Matrix[i] = tuples[i].coordinate
	}
	return *ans
}

func findAlpha(alphaMatrix, betaDiagonal, kMatrix lialg.Matrix, kkMatrices, eigenVectors []lialg.Matrix) lialg.Matrix {
	newAlphaMatrix := alphaMatrix
	for i := 0; i < alphaMatrix.N; i++ {
		cache := betaDiagonal.InverseDiagonal()

		cache = lialg.MatrixMultiplication(cache, kkMatrices[i])

		cache = lialg.MatrixMultiplication(cache, kMatrix)

		cache = lialg.MatrixMultiplication(cache, eigenVectors[i])

		newAlphaMatrix.Matrix[i] = cache.Matrix[0]
	}

	return newAlphaMatrix
}

func findBeta(beta, alphaMatrix, kk, k lialg.Matrix, eigen lialg.Eigen, lambda, sigma float64) lialg.Matrix {
	// prepare H matrix and f vector for qpsolver
	h := lialg.NewMatrix(beta.M, beta.M)
	f := lialg.NewMatrix(beta.M, 1)
	identity := lialg.CreateIdentity(beta.M)

	for i := 0; i < len(eigen.Values); i++ {
		alphaDiagonal := eigen.Vectors[i].VectorToDiagonalMatrix()

		cacheH := identity.MatrixScalarMultiplication(lambda * eigen.Values[i])
		cacheH = lialg.MatrixAddition(cacheH, kk)
		cacheH = lialg.MatrixMultiplication(alphaDiagonal, cacheH)
		cacheH = lialg.MatrixMultiplication(cacheH, alphaDiagonal)

		cacheH = lialg.MatrixAddition(cacheH, *h)
		h = &cacheH

		cacheF := eigen.Vectors[i].TransposeMatrix()
		cacheF = cacheF.MatrixScalarMultiplication(math.Sqrt(eigen.Values[i]))
		cacheF = lialg.MatrixMultiplication(cacheF, k)
		cacheF = lialg.MatrixMultiplication(cacheF, alphaDiagonal)
		cacheF = lialg.MatrixAddition(cacheF, *f)
		f = &cacheF
	}
	cacheF := f.MatrixScalarMultiplication(2)
	cache := lialg.CreateAllOnesVector(beta.M)
	cache = cache.TransposeMatrix()
	cache = cache.MatrixScalarMultiplication(sigma)
	cache = lialg.MatrixSubtraction(cache, cacheF)
	cache = cache.TransposeMatrix()
	f = &cache

	return lialg.QPSolve(*h, *f)
}
