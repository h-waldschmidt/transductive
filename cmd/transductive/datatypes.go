package transductive

type Coordinate struct{ X1, X2 float64 }

type Matrix struct {
	n, m   int
	matrix [][]float64
}
