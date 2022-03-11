package transductive

import "math"

func sequentialOptimization(points []Coordinate, numOfPoints int, lambda float64, sigma float64) []Coordinate {
	var selectedPoints []Coordinate

	//initialize the kVVMatrix
	kVVMatrix := calculateKernelMatrix(points, points, sigma)

	for len(selectedPoints) < numOfPoints {
		//select x to maximize the criteria
		var bestX Coordinate
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
// and takes the euclideanNorm of the resulting vector
func calculateCriteria(points []Coordinate, currentX Coordinate, sigma float64, lambda float64) float64 {
	kVxVector := calculateKernelVector(points, currentX, sigma)
	value, _ := euclideanNorm(kVxVector)
	value = math.Pow(value, 2) / (rbfKernel(currentX, currentX, sigma) + lambda)
	return value
}

func normalizeKvvMatrix(kVVMatrix Matrix, points []Coordinate, point Coordinate, lambda float64, sigma float64) Matrix {
	VxMatrix := calculateKernelVector(points, point, sigma)
	xVMatrix := transposeMatrix(VxMatrix)
	VxxVMatrix, _ := matrixMultiplication(VxMatrix, xVMatrix)

	VxxVMatrix = matrixScalarMultiplication(VxxVMatrix, 1/(1+lambda))
	kVVMatrix, _ = matrixSubtraction(kVVMatrix, VxxVMatrix)

	return kVVMatrix
}
