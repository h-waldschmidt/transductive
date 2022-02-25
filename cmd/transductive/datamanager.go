package transductive

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

type Vector struct {
	N      int
	Vector []float64
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
	// initializing the vector
	vector := Vector{len(pointsX), make([]float64, len(pointsX))}

	// calculating all the values
	for i := 0; i < len(pointsX); i++ {
		vector.Vector[i] = rbfKernel(pointsX[i], point, variance)

	}

	return vector
}

func euclideanDistance(x Vector, y Vector) (float64, error) {

	//the size of the vectors need to be the same
	if len(x.Vector) != len(y.Vector) {
		return 0, fmt.Errorf("could not calculate euclidean Distance")
	}

	var distance float64
	for i := 0; i < len(x.Vector); i++ {
		distance += math.Pow(x.Vector[i]-y.Vector[i], 2)
	}

	return math.Sqrt(distance), nil
}

func euclideanNorm(x Vector) float64 {

	var norm float64
	for i := 0; i < len(x.Vector); i++ {
		norm += math.Pow(x.Vector[i], 2)
	}

	return math.Sqrt(norm)
}

func matrixMultiplication(matrix1 Matrix, matrix2 Matrix) (Matrix, error) {

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
