package transductive

import (
	"math"
	"transductive-experimental-design/cmd/datamanager"
)

func SequentialOptimization(points datamanager.Matrix, numOfSelectedPoints int, lambda float64, sigma float64) datamanager.Matrix {
	selectedPoints := datamanager.Matrix{numOfSelectedPoints, points.M, make([][]float64, numOfSelectedPoints)}
	for i := 0; i < selectedPoints.N; i++ {
		selectedPoints.Matrix[i] = make([]float64, selectedPoints.M)
	}
	//initialize the kVVMatrix
	kVVMatrix, _ := datamanager.CalculateKernelMatrix(points, points, sigma)

	for j := 0; j < numOfSelectedPoints; j++ {
		//select x to maximize the criteria
		var bestX []float64
		bestValue := math.Inf(-1)

		for i := 0; i < points.N; i++ {
			currentX := points.Matrix[i]
			currentValue := calculateCriteria(kVVMatrix, currentX, i, sigma, lambda)
			if currentValue > bestValue {
				bestValue = currentValue
				bestX = currentX
			}
		}

		// add it the newly found x to the set
		selectedPoints.Matrix[j] = bestX
		// normalize the Kvv function by removing the influence of x
		kVVMatrix = normalizeKvvMatrix(kVVMatrix, points, bestX, sigma, lambda)
	}
	return selectedPoints
}

// basically calculates the distance from all points to the given point
// and takes the datamanager.euclideanNorm of the resulting vector
func calculateCriteria(kVVMatrix datamanager.Matrix, currentX []float64, index int, sigma float64, lambda float64) float64 {
	kVxVector := datamanager.Matrix{1, kVVMatrix.N, make([][]float64, 1)}
	kVxVector.Matrix[0] = kVVMatrix.Matrix[index]

	kxVVector := datamanager.TransposeMatrix(kVxVector)
	//value, _ := datamanager.EuclideanNorm(kVxVector)
	value, _ := datamanager.MatrixMultiplication(kxVVector, kVxVector)
	result := value.Matrix[0][0] / (1 + lambda)
	return result
}

func normalizeKvvMatrix(kVVMatrix datamanager.Matrix, points datamanager.Matrix, point []float64, lambda float64, sigma float64) datamanager.Matrix {
	VxMatrix, _ := datamanager.CalculateKernelVector(points, point, sigma)
	xVMatrix := datamanager.TransposeMatrix(VxMatrix)
	VxxVMatrix, _ := datamanager.MatrixMultiplication(VxMatrix, xVMatrix)

	VxxVMatrix = datamanager.MatrixScalarMultiplication(VxxVMatrix, 1/(1+lambda))
	kVVMatrix, _ = datamanager.MatrixSubtraction(kVVMatrix, VxxVMatrix)

	return kVVMatrix
}
