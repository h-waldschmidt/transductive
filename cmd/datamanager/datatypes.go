package datamanager

// TODO: Change everything to support Vector instead of Coordinate
type Coordinate struct{ X1, X2 float64 }

/**
 * N represents the number of columns (so that a vector can be in one array)
 * M represents the number rows (e.g. dimension of a vector)
 */
type Matrix struct {
	N, M   int
	Matrix [][]float64
}
