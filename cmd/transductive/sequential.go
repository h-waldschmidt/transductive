package transductive

import (
	"math"
	"transductive-experimental-design/cmd/datamanager"
)

// TODO: create global variables

//Sequential Algorithm for Transductive Experimental Design
//Searches in every iteration the best point useing the criterion
func SequentialOptimization(points datamanager.Matrix, numOfSelectedPoints int, lambda float64, sigma float64) datamanager.Matrix {
	selectedPoints := datamanager.NewMatrix(numOfSelectedPoints, points.M)

	//initialize the kVVMatrix
	kVVMatrix := datamanager.CalculateKernelMatrix(points, points, sigma)

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

// Criteria used for finding the best points
func calculateCriteria(kVVMatrix datamanager.Matrix, currentX []float64, index int, sigma float64, lambda float64) float64 {
	//initialize kVxMatrix
	kVxVector := datamanager.Matrix{1, kVVMatrix.N, make([][]float64, 1)}
	kVxVector.Matrix[0] = kVVMatrix.Matrix[index]

	kxVVector := kVxVector.TransposeMatrix()

	value := datamanager.MatrixMultiplication(kxVVector, kVxVector)

	result := value.Matrix[0][0] / (1 + lambda)
	return result
}

// After selecting a point the kVVMatrix has to be normalized,
//meaning the influence of the selected point has to be removed
func normalizeKvvMatrix(kVVMatrix datamanager.Matrix, points datamanager.Matrix, point []float64, lambda float64, sigma float64) datamanager.Matrix {
	VxMatrix := datamanager.CalculateKernelVector(points, point, sigma)

	xVMatrix := VxMatrix.TransposeMatrix()
	VxxVMatrix := datamanager.MatrixMultiplication(VxMatrix, xVMatrix)

	VxxVMatrix = VxxVMatrix.MatrixScalarMultiplication(1 / (1 + lambda))
	kVVMatrix = datamanager.MatrixSubtraction(kVVMatrix, VxxVMatrix)

	return kVVMatrix
}
