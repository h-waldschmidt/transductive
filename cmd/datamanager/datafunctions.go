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

func CalculateKernelMatrix(pointsX Matrix, pointsY Matrix, sigma float64) (Matrix, error) {

	//x and y need to have the same dimensions
	if pointsX.N != pointsY.N {
		return Matrix{0, 0, make([][]float64, 0)}, fmt.Errorf("could not use RBFKernel")
	}

	// initializing the matrix
	matrix := Matrix{pointsY.M, pointsX.M, make([][]float64, pointsY.M)}
	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i] = make([]float64, pointsX.M)
	}

	// calculating all the values
	for i := 0; i < matrix.N; i++ {
		for j := 0; j < matrix.M; j++ {
			matrix.Matrix[j][i], _ = RbfKernel(pointsX.Matrix[i], pointsY.Matrix[j], sigma)
		}
	}

	return matrix, nil
}

func CalculateKernelVector(pointsX Matrix, point []float64, sigma float64) (Matrix, error) {

	//x and y need to have the same dimensions
	if pointsX.N != len(point) {
		return Matrix{0, 0, make([][]float64, 0)}, fmt.Errorf("could not use RBFKernel")
	}

	// initializing the vector(Matrix with M=1)
	vector := Matrix{1, pointsX.N, make([][]float64, 1)}
	vector.Matrix[0] = make([]float64, vector.M)

	// calculating all the values
	for i := 0; i < pointsX.N; i++ {
		vector.Matrix[0][i], _ = RbfKernel(pointsX.Matrix[i], point, sigma)
	}

	return vector, nil
}

func EuclideanDistance(x []float64, y []float64) (float64, error) {

	//x and y need to be vectors and have the same dimensions
	if len(x) != len(y) {
		return 0, fmt.Errorf("could not calculate euclidean Distance")
	}

	var distance float64
	for i := 0; i < len(x); i++ {
		distance += math.Pow(x[i]-y[i], 2)
	}

	return math.Sqrt(distance), nil
}

func EuclideanNorm(x []float64) float64 {

	var norm float64
	for i := 0; i < len(x); i++ {
		norm += math.Pow(x[i], 2)
	}

	return math.Sqrt(norm)
}

func MatrixMultiplication(matrix1 Matrix, matrix2 Matrix) (Matrix, error) {

	// The inner dimensions need to be the same
	if matrix1.M != matrix2.N {
		return Matrix{0, 0, [][]float64{}}, fmt.Errorf("could not multiply the matrices")
	}

	// initialize the matrix
	matrix := Matrix{matrix2.M, matrix1.N, make([][]float64, matrix2.M)}
	for i := 0; i < matrix.M; i++ {
		matrix.Matrix[i] = make([]float64, matrix1.N)
	}

	// need to test if this is the cache efficient version of matrix multiplication
	for i := 0; i < matrix1.N; i++ {
		for j := 0; j < matrix2.M; j++ {
			var sum float64
			for k := 0; k < matrix1.M; k++ {
				sum += matrix1.Matrix[k][i] * matrix2.Matrix[j][k]
			}
			matrix.Matrix[j][i] = sum
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
		matrix.Matrix[i] = make([]float64, matrix.M)
	}

	for i := 0; i < matrix.N; i++ {
		for j := 0; j < matrix.M; j++ {
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
		matrix.Matrix[i] = make([]float64, matrix.M)
	}

	for i := 0; i < matrix.N; i++ {
		for j := 0; j < matrix.M; j++ {
			matrix.Matrix[i][j] = matrix1.Matrix[i][j] - matrix2.Matrix[i][j]
		}
	}

	return matrix, nil
}

func MatrixScalarMultiplication(matrix Matrix, scalar float64) Matrix {

	for i := 0; i < matrix.N; i++ {
		for j := 0; j < matrix.M; j++ {
			matrix.Matrix[i][j] *= scalar
		}
	}
	return matrix
}

func ConvertSliceToCoordinate(point []float64) (Coordinate, error) {
	// point has to have dimension 2
	if len(point) != 2 {
		return Coordinate{}, fmt.Errorf("could not add the matrices")
	}
	return Coordinate{point[0], point[1]}, nil
}
func ConvertMatrixToCoordinateSlice(matrix Matrix) ([]Coordinate, error) {
	// point has to have dimension 2
	if matrix.M != 2 {
		return nil, fmt.Errorf("could not add the matrices")
	}

	slice := make([]Coordinate, matrix.N)
	for i := 0; i < matrix.N; i++ {
		slice[i] = Coordinate{matrix.Matrix[i][0], matrix.Matrix[i][1]}
	}

	return slice, nil
}

func ConvertCoordinatesToMatrix(points []Coordinate) Matrix {
	//initialize matrix
	matrix := Matrix{len(points), 2, make([][]float64, len(points))}
	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i] = make([]float64, matrix.M)
	}

	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i][0] = points[i].X1
		matrix.Matrix[i][1] = points[i].X2
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

func PlotSelectedPoints(items []Coordinate, selectedPoints []Coordinate, path string) error {
	var itemsXYs plotter.XYs
	for _, xy := range items {
		itemsXYs = append(itemsXYs, struct{ X, Y float64 }{xy.X1, xy.X2})
	}

	var points plotter.XYs
	for _, xy := range selectedPoints {
		points = append(points, struct{ X, Y float64 }{xy.X1, xy.X2})
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

	// Add the items as an scatter plot
	po, err := plotter.NewScatter(points)
	if err != nil {
		return fmt.Errorf("could not create scatter: %v", err)
	}

	s.GlyphStyle.Shape = draw.CrossGlyph{}
	s.Color = color.RGBA{R: 255, A: 255}
	p.Add(s)

	po.GlyphStyle.Shape = draw.CrossGlyph{}
	po.Color = color.RGBA{R: 0, A: 255}
	p.Add(po)

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
