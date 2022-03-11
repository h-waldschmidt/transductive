package datamanager

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

// TODO: Change everything to support Vector instead of Coordinate
type Coordinate struct{ X1, X2 float64 }

type Matrix struct {
	N, M   int
	Matrix [][]float64
}

func CalculateKernelMatrix(pointsX []Coordinate, pointsY []Coordinate, variance float64) Matrix {
	// initializing the matrix
	matrix := Matrix{len(pointsX), len(pointsY), make([][]float64, len(pointsX))}
	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i] = make([]float64, len(pointsY))
	}

	// calculating all the values
	for i := 0; i < len(pointsX); i++ {
		for j := 0; j < len(pointsY); j++ {
			matrix.Matrix[i][j] = RbfKernel(pointsX[i], pointsY[j], variance)
		}
	}

	return matrix
}

func CalculateKernelVector(pointsX []Coordinate, point Coordinate, variance float64) Matrix {

	// initializing the vector(Matrix with M=1)
	vector := Matrix{len(pointsX), 1, make([][]float64, 1)}
	vector.Matrix[0] = make([]float64, len(pointsX))

	// calculating all the values
	for i := 0; i < len(pointsX); i++ {
		vector.Matrix[0][i] = RbfKernel(pointsX[i], point, variance)
	}

	return vector
}

func EuclideanDistance(x Matrix, y Matrix) (float64, error) {

	//x and y need to be vectors and have the same dimensions
	if x.N != y.N || x.M > 1 || y.M > 1 {
		return 0, fmt.Errorf("could not calculate euclidean Distance")
	}

	var distance float64
	for i := 0; i < x.N; i++ {
		distance += math.Pow(x.Matrix[0][i]-y.Matrix[0][i], 2)
	}

	return math.Sqrt(distance), nil
}

func EuclideanNorm(x Matrix) (float64, error) {
	//x need to be a vector
	if x.M > 1 {
		return 0, fmt.Errorf("could not calculate euclidean Norm")
	}

	var norm float64
	for i := 0; i < x.N; i++ {
		norm += math.Pow(x.Matrix[0][i], 2)
	}

	return math.Sqrt(norm), nil
}

func MatrixMultiplication(matrix1 Matrix, matrix2 Matrix) (Matrix, error) {

	// The inner dimensions need to be the same
	if matrix1.M != matrix2.N {
		return Matrix{0, 0, [][]float64{}}, fmt.Errorf("could not multiply the matrices")
	}

	// initialize the matrix
	matrix := Matrix{matrix1.N, matrix2.M, make([][]float64, matrix1.N)}
	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i] = make([]float64, matrix2.M)
	}

	// need to test if this is the cache efficient version of matrix multiplication
	for i := 0; i < matrix1.N; i++ {
		for j := 0; j < matrix2.M; j++ {
			var sum float64
			for k := 0; k < matrix1.M; k++ {
				sum += matrix1.Matrix[i][k] * matrix2.Matrix[k][j]
			}
			matrix.Matrix[i][j] = sum
		}
	}

	return matrix, nil
}

func TransposeMatrix(matrix Matrix) Matrix {
	//initialize the transpose matrix
	transpose := Matrix{matrix.M, matrix.N, make([][]float64, matrix.M)}
	for i := 0; i < transpose.N; i++ {
		transpose.Matrix[i] = make([]float64, transpose.M)
	}

	for i := 0; i < transpose.N; i++ {
		for j := 0; j < transpose.M; j++ {
			transpose.Matrix[i][j] = matrix.Matrix[j][i]
		}
	}

	return transpose
}

func MatrixAddition(matrix1 Matrix, matrix2 Matrix) (Matrix, error) {

	// the dimensions of the matrices have to match
	if matrix1.N != matrix2.N || matrix1.M != matrix2.M {
		return Matrix{0, 0, [][]float64{}}, fmt.Errorf("could not add the matrices")
	}

	//initialize the matrix
	matrix := Matrix{matrix1.N, matrix1.M, make([][]float64, matrix1.N)}
	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i] = make([]float64, matrix1.M)
	}

	for i := 0; i < matrix1.N; i++ {
		for j := 0; j < matrix1.M; j++ {
			matrix.Matrix[i][j] = matrix1.Matrix[i][j] + matrix2.Matrix[i][j]
		}
	}

	return matrix, nil
}

func MatrixSubtraction(matrix1 Matrix, matrix2 Matrix) (Matrix, error) {
	// the dimensions of the matrices have to match
	if matrix1.N != matrix2.N || matrix1.M != matrix2.M {
		return Matrix{0, 0, [][]float64{}}, fmt.Errorf("could not add the matrices")
	}

	//initialize the matrix
	matrix := Matrix{matrix1.N, matrix1.M, make([][]float64, matrix1.N)}
	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i] = make([]float64, matrix1.M)
	}

	for i := 0; i < matrix1.N; i++ {
		for j := 0; j < matrix1.M; j++ {
			matrix.Matrix[i][j] = matrix1.Matrix[i][j] + matrix2.Matrix[i][j]
		}
	}

	return matrix, nil
}

func MatrixScalarMultiplication(matrix Matrix, scalar float64) Matrix {

	for i := 0; i < matrix.N; i++ {
		for j := 0; j < matrix.M; j++ {
			matrix.Matrix[i][j] /= scalar
		}
	}
	return matrix
}

func CreateNormalDistribution(mean float64, standardDeviation float64, numberOfItems int) []Coordinate {
	var distribution []Coordinate

	for i := 0; i < numberOfItems; i++ {
		xy := Coordinate{rand.NormFloat64()*standardDeviation + mean, rand.NormFloat64()*standardDeviation + mean}
		distribution = append(distribution, xy)
	}
	return distribution
}

func PlotDistribution(items []Coordinate, path string) error {
	var itemsXYs plotter.XYs
	for _, xy := range items {
		itemsXYs = append(itemsXYs, struct{ X, Y float64 }{xy.X1, xy.X2})
	}
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create %s: %v", path, err)
	}

	p := plot.New()

	// Add the items as an scatter plot
	s, err := plotter.NewScatter(itemsXYs)
	if err != nil {
		return fmt.Errorf("could not create scatter: %v", err)
	}
	s.GlyphStyle.Shape = draw.CrossGlyph{}
	s.Color = color.RGBA{R: 255, A: 255}
	p.Add(s)

	wt, err := p.WriterTo(256, 256, "png")
	if err != nil {
		return fmt.Errorf("could not create writer: %v", err)
	}
	_, err = wt.WriteTo(f)
	if err != nil {
		return fmt.Errorf("could not write to %s: %v", path, err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("could not close %s: %v", path, err)
	}
	return nil
}
