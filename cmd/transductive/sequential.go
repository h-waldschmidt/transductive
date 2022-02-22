package transductive

import "math"

func sequentialOptimization(points []Coordinate, numOfPoints int, mean float64) {
	var selectedPoints []Coordinate

	// find out how to calculate the variance
	kVVMatrix := calculateKernelMatrix(points, points, 0)

	for len(selectedPoints) < numOfPoints {
		//select x to maximize function
		var currentX Coordinate
		bestValue := math.Inf(1)

		for i := 0; i < len(points); i++ {
			currentX = points[i]

		}
		// add it to set

		// normalize the Kvv function
	}

}

func calculateCriteria(points []Coordinate, currentX Coordinate, variance float64, mean float64) float64 {
	kVxVector := calculateKernelVector(points, currentX, variance)
	value := math.Pow(euclideanNorm(kVxVector), 2) / (rbfKernel(currentX, currentX, variance) + mean)
	return value
}
