package transductive

type Coordinate struct{ X1, X2 float64 }

type Matrix struct {
	N, M   int
	Matrix [][]float64
}
