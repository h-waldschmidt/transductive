package datamanager

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
