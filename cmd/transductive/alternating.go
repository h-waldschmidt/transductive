package transductive

import (
	"transductive-experimental-design/cmd/datamanager"
)

func AlternatingOptimization(points datamanager.Matrix, numOfSelectedPoints int, lambda float64, sigma float64) datamanager.Matrix {
	//repeat until no major improvement
	// find optimal alpha
	// find optimal beta
	// normalize Beta Matrix

	// extract selected Points from Beta Matrix,
	// by selecting the numOfSelectedPoints biggest points
}

func findAlpha() {}

func findBeta() {}

func normalizeBetaMatrix() {}
