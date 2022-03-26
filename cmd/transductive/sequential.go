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
			currentValue := calculateCriteria(kVVMatrix, currentX, i, sigma, lambda)
			if currentValue > bestValue {
				bestValue = currentValue
				bestX = currentX
			}
		}

		// add it the newly found x to the set
		selectedPoints = append(selectedPoints, bestX)
		// normalize the Kvv function by removing the influence of x
		kVVMatrix = normalizeKvvMatrix(kVVMatrix, points, bestX, sigma, lambda)
	}
	return selectedPoints
}

// basically calculates the distance from all points to the given point
// and takes the datamanager.euclideanNorm of the resulting vector
func calculateCriteria(kVVMatrix datamanager.Matrix, currentX datamanager.Coordinate, index int, sigma float64, lambda float64) float64 {
	kVxVector := datamanager.Matrix{kVVMatrix.N, 1, make([][]float64, kVVMatrix.N)}
	for i := 0; i < kVxVector.N; i++ {
		kVxVector.Matrix[i] = []float64{kVVMatrix.Matrix[index][i]}
	}

	kxVVector := datamanager.TransposeMatrix(kVxVector)
	//value, _ := datamanager.EuclideanNorm(kVxVector)
	value, _ := datamanager.MatrixMultiplication(kxVVector, kVxVector)
	result := value.Matrix[0][0] / (datamanager.RbfKernel(currentX, currentX, sigma) + lambda)
	return result
}

func normalizeKvvMatrix(kVVMatrix datamanager.Matrix, points []datamanager.Coordinate, point datamanager.Coordinate, lambda float64, sigma float64) datamanager.Matrix {
	VxMatrix := datamanager.CalculateKernelVector(points, point, sigma)
	xVMatrix := datamanager.TransposeMatrix(VxMatrix)
	VxxVMatrix, _ := datamanager.MatrixMultiplication(VxMatrix, xVMatrix)

	VxxVMatrix = datamanager.MatrixScalarMultiplication(VxxVMatrix, 1/(datamanager.RbfKernel(point, point, sigma)+lambda))
	kVVMatrix, _ = datamanager.MatrixSubtraction(kVVMatrix, VxxVMatrix)

	return kVVMatrix
}
