package transductive

import (
	"math"
	"transductive-experimental-design/cmd/datamanager"
)

func SequentialOptimization(points []datamanager.Coordinate, numOfSelectedPoints int, lambda float64, sigma float64) []datamanager.Coordinate {
	var selectedPoints []datamanager.Coordinate

	//initialize the kVVMatrix
	kVVMatrix := datamanager.CalculateKernelMatrix(points, points, sigma)

	for len(selectedPoints) < numOfSelectedPoints {
		//select x to maximize the criteria
		var bestX datamanager.Coordinate
		bestValue := math.Inf(-1)

		for i := 0; i < len(points); i++ {
			currentX := points[i]
			currentValue := calculateCriteria(points, currentX, sigma, lambda)
			if currentValue > bestValue {
				bestValue = currentValue
				bestX = currentX
			}
		}

		// add it the newly found x to the set
		selectedPoints = append(selectedPoints, bestX)
		// normalize the Kvv function by removing the influence of x
		kVVMatrix = normalizeKvvMatrix(kVVMatrix, points, bestX, lambda, sigma)
	}
	return selectedPoints
}

// basically calculates the distance from all points to the given point
// and takes the datamanager.euclideanNorm of the resulting vector
func calculateCriteria(points []datamanager.Coordinate, currentX datamanager.Coordinate, sigma float64, lambda float64) float64 {
	kVxVector := datamanager.CalculateKernelVector(points, currentX, sigma)
	value, _ := datamanager.EuclideanNorm(kVxVector)
	value = math.Pow(value, 2) / (datamanager.RbfKernel(currentX, currentX, sigma) + lambda)
	return value
}

func normalizeKvvMatrix(kVVMatrix datamanager.Matrix, points []datamanager.Coordinate, point datamanager.Coordinate, lambda float64, sigma float64) datamanager.Matrix {
	VxMatrix := datamanager.CalculateKernelVector(points, point, sigma)
	xVMatrix := datamanager.TransposeMatrix(VxMatrix)
	VxxVMatrix, _ := datamanager.MatrixMultiplication(VxMatrix, xVMatrix)

	VxxVMatrix = datamanager.MatrixScalarMultiplication(VxxVMatrix, 1/(1+lambda))
	kVVMatrix, _ = datamanager.MatrixSubtraction(kVVMatrix, VxxVMatrix)

	return kVVMatrix
}
