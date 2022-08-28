package plt

import (
	"log"

	"github.com/h-waldschmidt/transductive/internal/lialg"
)

// only used for visualizing 2D Vectors
type Coordinate struct{ X1, X2 float64 }

func ConvertSliceToCoordinate(point []float64) Coordinate {
	// point has to have dimension 2
	if len(point) != 2 {
		log.Fatal("points do not have 2 dimensions")
	}
	return Coordinate{point[0], point[1]}
}

func ConvertMatrixToCoordinateSlice(matrix *lialg.Matrix) []Coordinate {
	// point has to have dimension 2
	if matrix.M != 2 {
		log.Fatal("points do not have 2 dimensions")
	}

	slice := make([]Coordinate, matrix.N)
	for i := 0; i < matrix.N; i++ {
		slice[i] = Coordinate{matrix.Matrix[i][0], matrix.Matrix[i][1]}
	}

	return slice
}

func ConvertCoordinatesToMatrix(points []Coordinate) lialg.Matrix {
	//initialize matrix
	matrix := lialg.NewMatrix(len(points), 2)

	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i][0] = points[i].X1
		matrix.Matrix[i][1] = points[i].X2
	}

	return *matrix
}
