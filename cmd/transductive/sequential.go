package transductive

import "math"

func sequentialOptimization(points []Coordinate, numOfPoints int, mean float64) []Coordinate {
	var selectedPoints []Coordinate

	//TODO: find out how to calculate the variance
	var variance float64
	//initialize the kVVMatrix
	kVVMatrix := calculateKernelMatrix(points, points, variance)

	for len(selectedPoints) < numOfPoints {
		//select x to maximize the criteria
		var bestX Coordinate
		bestValue := math.Inf(1)

		for i := 0; i < len(points); i++ {
			currentX := points[i]
			currentValue := calculateCriteria(points, currentX, variance, mean)
			if currentValue < bestValue {
				bestValue = currentValue
				bestX = currentX
			}
		}

		// add it the newly found x to the set
		selectedPoints = append(selectedPoints, bestX)
		// normalize the Kvv function by removing the influence of x
		kVVMatrix = normalizeKvvMatrix(kVVMatrix, points, bestX, mean, variance)
	}
	return selectedPoints
}

// basically calculates the distance from all points to the given point
// and takes the euclideanNorm of the resulting vector
func calculateCriteria(points []Coordinate, currentX Coordinate, variance float64, mean float64) float64 {
	kVxVector := calculateKernelVector(points, currentX, variance)
	value, _ := euclideanNorm(kVxVector)
	value = math.Pow(value, 2) / (rbfKernel(currentX, currentX, variance) + mean)
	return value
}

func normalizeKvvMatrix(kVVMatrix Matrix, points []Coordinate, point Coordinate, mean float64, variance float64) Matrix {
	VxMatrix := calculateKernelVector(points, point, variance)
	xVMatrix := transposeMatrix(VxMatrix)
	VxxVMatrix, _ := matrixMultiplication(VxMatrix, xVMatrix)

	VxxVMatrix = matrixScalarMultiplication(VxxVMatrix, 1/(1+mean))
	kVVMatrix, _ = matrixSubtraction(kVVMatrix, VxxVMatrix)

	return kVVMatrix
}
