package datamanager

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

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
