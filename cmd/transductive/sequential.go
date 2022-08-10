package transductive

import (
	"math"
	"sync"
	"transductive-experimental-design/cmd/lialg"
)

// TODO: create global variables

// Sequential Algorithm for Transductive Experimental Design
// Searches in every iteration the best point useing the criterion
func SequentialOptimization(points lialg.Matrix, numOfSelectedPoints int, lambda, sigma float64) lialg.Matrix {
	selectedPoints := lialg.NewMatrix(numOfSelectedPoints, points.M)

	//initialize the kVVMatrix
	kVVMatrix := points.CalculateKernelMatrix(points, sigma)

	for j := 0; j < numOfSelectedPoints; j++ {
		//select x to maximize the criteria
		var bestX []float64
		bestValue := math.Inf(-1)
		var wg sync.WaitGroup
		var mutex sync.Mutex
		for i := 0; i < points.N; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				currentX := points.Matrix[i]
				currentValue := calculateCriteria(kVVMatrix, currentX, i, sigma, lambda)
				mutex.Lock()
				{
					if currentValue > bestValue {
						bestValue = currentValue
						bestX = currentX
					}
				}
				mutex.Unlock()
			}(i)
		}
		wg.Wait()

		// add it the newly found x to the set
		selectedPoints.Matrix[j] = bestX
		// normalize the Kvv function by removing the influence of x
		kVVMatrix = normalizeKvvMatrix(kVVMatrix, points, bestX, sigma, lambda)
	}
	return *selectedPoints
}

// Criteria used for finding the best points
func calculateCriteria(kVVMatrix lialg.Matrix, currentX []float64, index int, sigma, lambda float64) float64 {
	//initialize kVxMatrix
	kVxVector := lialg.Matrix{N: 1, M: kVVMatrix.N, Matrix: make([][]float64, 1)}
	kVxVector.Matrix[0] = kVVMatrix.Matrix[index]

	kxVVector := kVxVector.TransposeMatrix()

	value := lialg.MatrixMultiplication(kxVVector, kVxVector)

	result := value.Matrix[0][0] / (1 + lambda)
	return result
}

// After selecting a point the kVVMatrix has to be normalized,
// meaning the influence of the selected point has to be removed
func normalizeKvvMatrix(kVVMatrix, points lialg.Matrix, point []float64, lambda, sigma float64) lialg.Matrix {
	VxMatrix := points.CalculateKernelVector(point, sigma)

	xVMatrix := VxMatrix.TransposeMatrix()
	VxxVMatrix := lialg.MatrixMultiplication(VxMatrix, xVMatrix)

	VxxVMatrix.MatrixScalarMultiplication(1 / (1 + lambda))
	kVVMatrix = lialg.MatrixSubtraction(kVVMatrix, VxxVMatrix)

	return kVVMatrix
}
