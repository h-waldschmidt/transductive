package transductive

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

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
