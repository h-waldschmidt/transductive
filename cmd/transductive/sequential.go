package transductive

import (
	"math"
	"transductive-experimental-design/cmd/datamanager"
	"transductive-experimental-design/cmd/kernelregression"
)

func sequentialOptimization(points []datamanager.Coordinate, numOfPoints int, lambda float64, sigma float64) []datamanager.Coordinate {
	var selectedPoints []datamanager.Coordinate

	//initialize the kVVMatrix
	kVVMatrix := datamanager.calculateKernelMatrix(points, points, sigma)

	for len(selectedPoints) < numOfPoints {
		//select x to maximize the criteria
		var bestX datamanager.Coordinate
		bestValue := math.Inf(1)

		for i := 0; i < len(points); i++ {
			currentX := points[i]
			currentValue := calculateCriteria(points, currentX, sigma, lambda)
			if currentValue < bestValue {
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
	kVxVector := datamanager.calculateKernelVector(points, currentX, sigma)
	value, _ := datamanager.euclideanNorm(kVxVector)
	value = math.Pow(value, 2) / (kernelregression.RbfKernel(currentX, currentX, sigma) + lambda)
	return value
}

func normalizeKvvMatrix(kVVMatrix datamanager.Matrix, points []datamanager.Coordinate, point datamanager.Coordinate, lambda float64, sigma float64) datamanager.Matrix {
	VxMatrix := datamanager.calculateKernelVector(points, point, sigma)
	xVMatrix := datamanager.transposeMatrix(VxMatrix)
	VxxVMatrix, _ := datamanager.matrixMultiplication(VxMatrix, xVMatrix)

	VxxVMatrix = datamanager.matrixScalarMultiplication(VxxVMatrix, 1/(1+lambda))
	kVVMatrix, _ = datamanager.matrixSubtraction(kVVMatrix, VxxVMatrix)

	return kVVMatrix
}
