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

type Coordinate struct{ X, Y float64 }

func CreateNormalDistribution(mean float64, standardDeviation float64, numberOfItems int) {
	var items plotter.XYs

	for i := 0; i < numberOfItems; i++ {
		xy := Coordinate{rand.NormFloat64() * 0.1, rand.NormFloat64() * 0.1}
		items = append(items, struct{ X, Y float64 }{xy.X, xy.Y})
	}
	plotDistribution(items, "../../plots/test.png")
}

func plotDistribution(items plotter.XYs, path string) error {

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create %s: %v", path, err)
	}

	p := plot.New()

	// create scatter with all data points
	s, err := plotter.NewScatter(items)
	if err != nil {
		return fmt.Errorf("could not create scatter: %v", err)
	}
	s.GlyphStyle.Shape = draw.CrossGlyph{}
	s.Color = color.RGBA{R: 255, A: 255}
	p.Add(s)

	var x, c float64
	x = 1.2
	c = -3

	// create fake linear regression result
	l, err := plotter.NewLine(plotter.XYs{
		{3, 3*x + c}, {20, 20*x + c},
	})
	if err != nil {
		return fmt.Errorf("could not create line: %v", err)
	}
	p.Add(l)

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
