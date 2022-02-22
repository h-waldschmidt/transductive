package transductive

import "math"

func Transductive() {

}

func rbfKernel(x1 Coordinate, x2 Coordinate, variance float64) float64 {
	value := -math.Pow(x1.X1-x2.X1, 2) + math.Pow(x1.X2-x2.X2, 2)
	value = math.Sqrt(value)
	value /= 2 * variance
	value = math.Exp(value)

	return value
}

func calculateKernelMatrix(pointsX []Coordinate, pointsY []Coordinate, variance float64) Matrix {
	// initializing the matrix
	matrix := Matrix{len(pointsX), len(pointsY), make([][]float64, len(pointsX))}
	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i] = make([]float64, len(pointsY))
	}

	// calculating all the values
	for i := 0; i < len(pointsX); i++ {
		for j := 0; j < len(pointsY); j++ {
			matrix.Matrix[i][j] = rbfKernel(pointsX[i], pointsY[j], variance)
		}
	}

	return matrix
}

func calculateKernelVector(pointsX []Coordinate, point Coordinate, variance float64) Vector {
	// initializing the matrix
	vector := Vector{len(pointsX), make([]float64, len(pointsX))}

	// calculating all the values
	for i := 0; i < len(pointsX); i++ {
		vector.Vector[i] = rbfKernel(pointsX[i], point, variance)

	}

	return vector
}

func euclideanDistance(x Coordinate, y Coordinate) float64 {
	distance := math.Pow(x.X1-y.X1, 2) + math.Pow(x.X2-y.X2, 2)

	return math.Sqrt(distance)
}
