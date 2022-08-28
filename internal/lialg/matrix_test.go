package lialg

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEuclideanDistanceNormal(t *testing.T) {
	x := Matrix{N: 1, M: 3, Matrix: [][]float64{{1, 2, 3}}}
	y := Matrix{N: 1, M: 3, Matrix: [][]float64{{2, 3, 4}}}

	value := EuclideanDistance(x.Matrix[0], y.Matrix[0])
	expected := math.Sqrt(3)

	if value != expected {
		t.Errorf("Expected: %f ; Got: %f", expected, value)
	}
}

func TestEuclideanNormNormal(t *testing.T) {
	x := Matrix{N: 1, M: 3, Matrix: [][]float64{{1, 2, 3}}}

	value := EuclideanNorm(x.Matrix[0])
	expected := math.Sqrt(14)

	if value != expected {
		t.Errorf("Expected: %f ; Got: %f", expected, value)
	}
}

func TestSumNormNormal(t *testing.T) {
	x := Matrix{N: 1, M: 3, Matrix: [][]float64{{1, 2, -3}}}

	value := SumNorm(x.Matrix[0])
	expected := 6.0

	if value != expected {
		t.Errorf("Expected: %f ; Got: %f", expected, value)
	}
}

func TestSliceToDiagonalMatrixNormal(t *testing.T) {
	x := []float64{1, 2, 3}

	value := SliceToDiagonalMatrix(x)
	expected := Matrix{N: 3, M: 3, Matrix: [][]float64{{1, 0, 0}, {0, 2, 0}, {0, 0, 3}}}

	if !cmp.Equal(value, expected) {
		t.Errorf("Expected: %v ; Got: %v", expected, value)
	}
}

func TestDiagonalMatrixToSliceNormal(t *testing.T) {
	x := Matrix{N: 3, M: 3, Matrix: [][]float64{{1, 0, 0}, {0, 2, 0}, {0, 0, 3}}}

	value := x.DiagonalMatrixToSlice()
	expected := []float64{1, 2, 3}

	if !cmp.Equal(value, expected) {
		t.Errorf("Expected: %v ; Got: %v", expected, value)
	}
}

func TestVectorToDiagonalMatrixNormal(t *testing.T) {
	x := Matrix{N: 1, M: 3, Matrix: [][]float64{{1, 2, 3}}}

	value := x.VectorToDiagonalMatrix()
	expected := Matrix{N: 3, M: 3, Matrix: [][]float64{{1, 0, 0}, {0, 2, 0}, {0, 0, 3}}}

	if !cmp.Equal(value, expected) {
		t.Errorf("Expected: %v ; Got: %v", expected, value)
	}
}

func TestDiagonalMatrixToVectorNormal(t *testing.T) {
	x := Matrix{N: 3, M: 3, Matrix: [][]float64{{1, 0, 0}, {0, 2, 0}, {0, 0, 3}}}

	value := x.DiagonalMatrixToVector()
	expected := Matrix{N: 1, M: 3, Matrix: [][]float64{{1, 2, 3}}}

	if !cmp.Equal(value, expected) {
		t.Errorf("Expected: %v ; Got: %v", expected, value)
	}
}

func TestCreateIdentityNormal(t *testing.T) {
	value := CreateIdentity(3)
	expected := Matrix{N: 3, M: 3, Matrix: [][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}}

	if !cmp.Equal(value, expected) {
		t.Errorf("Expected: %v ; Got: %v", expected, value)
	}
}

func TestCreateAllOnesVectorNormal(t *testing.T) {
	value := CreateAllOnesVector(3)
	expected := Matrix{N: 1, M: 3, Matrix: [][]float64{{1, 1, 1}}}

	if !cmp.Equal(value, expected) {
		t.Errorf("Expected: %v ; Got: %v", expected, value)
	}
}

func TestMatrixMultiplicationNormal(t *testing.T) {
	t.Run("ZeroMult", func(t *testing.T) {
		a := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}
		b := Matrix{2, 3, [][]float64{{0, 0, 0}, {0, 0, 0}}}

		value := MatrixMultiplication(a, b)
		expected := Matrix{2, 2, [][]float64{{0, 0}, {0, 0}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})

	t.Run("MatrixWithMatrix", func(t *testing.T) {
		a := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}
		b := Matrix{2, 3, [][]float64{{1, 0, 4}, {2, 1, 0}}}

		value := MatrixMultiplication(a, b)
		expected := Matrix{2, 2, [][]float64{{7, 9}, {8, 2}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})

	t.Run("MatrixWithVector", func(t *testing.T) {
		a := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}
		b := Matrix{1, 3, [][]float64{{1, 0, 4}}}

		value := MatrixMultiplication(a, b)
		expected := Matrix{1, 2, [][]float64{{7, 9}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})

	t.Run("VectorWithVector", func(t *testing.T) {
		a := Matrix{3, 1, [][]float64{{3}, {2}, {1}}}
		b := Matrix{1, 3, [][]float64{{1, 0, 4}}}

		value := MatrixMultiplication(b, a)
		expected := Matrix{3, 3, [][]float64{{3, 0, 12}, {2, 0, 8}, {1, 0, 4}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})
}

func TestTransposeMatrixNormal(t *testing.T) {
	t.Run("ZeroMatrixTranspose", func(t *testing.T) {
		a := Matrix{2, 3, [][]float64{{0, 0, 0}, {0, 0, 0}}}

		value := a.TransposeMatrix()
		expected := Matrix{3, 2, [][]float64{{0, 0}, {0, 0}, {0, 0}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})

	t.Run("VectorTranspose", func(t *testing.T) {
		a := Matrix{1, 3, [][]float64{{1, 0, 4}}}

		value := a.TransposeMatrix()
		expected := Matrix{3, 1, [][]float64{{1}, {0}, {4}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})

	t.Run("MatrixTranspose", func(t *testing.T) {
		a := Matrix{3, 2, [][]float64{{3, 2}, {2, 0}, {12, 8}}}

		value := a.TransposeMatrix()
		expected := Matrix{2, 3, [][]float64{{3, 2, 12}, {2, 0, 8}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})
}

func TestMatrixAdditionNormal(t *testing.T) {
	t.Run("ZeroAddition", func(t *testing.T) {
		a := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}
		b := Matrix{3, 2, [][]float64{{0, 0}, {0, 0}, {0, 0}}}

		value := MatrixAddition(a, b)
		expected := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})

	t.Run("VectorAddition", func(t *testing.T) {
		a := Matrix{1, 3, [][]float64{{3, 1, 2}}}
		b := Matrix{1, 3, [][]float64{{3, 1, 2}}}

		value := MatrixAddition(a, b)
		expected := Matrix{1, 3, [][]float64{{6, 2, 4}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})

	t.Run("MatrixAddition", func(t *testing.T) {
		a := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}
		b := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}

		value := MatrixAddition(a, b)
		expected := Matrix{3, 2, [][]float64{{6, 2}, {4, 0}, {2, 4}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})
}

func TestMatrixSubtractionNormal(t *testing.T) {
	t.Run("ZeroSubtraction", func(t *testing.T) {
		a := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}
		b := Matrix{3, 2, [][]float64{{0, 0}, {0, 0}, {0, 0}}}

		value := MatrixSubtraction(a, b)
		expected := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})

	t.Run("VectorSubtraction", func(t *testing.T) {
		a := Matrix{1, 3, [][]float64{{3, 1, 2}}}
		b := Matrix{1, 3, [][]float64{{2, 1, 1}}}

		value := MatrixSubtraction(a, b)
		expected := Matrix{1, 3, [][]float64{{1, 0, 1}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})

	t.Run("MatrixSubtraction", func(t *testing.T) {
		a := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}
		b := Matrix{3, 2, [][]float64{{2, 1}, {1, -3}, {0, -1}}}

		value := MatrixSubtraction(a, b)
		expected := Matrix{3, 2, [][]float64{{1, 0}, {1, 3}, {1, 3}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})
}

func TestMatrixScalarMultiplicationNormal(t *testing.T) {
	t.Run("ZeroMultiplication", func(t *testing.T) {
		a := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}

		value := a.MatrixScalarMultiplication(0)
		expected := Matrix{3, 2, [][]float64{{0, 0}, {0, 0}, {0, 0}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})

	t.Run("PositiveMultiplication", func(t *testing.T) {
		a := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}

		value := a.MatrixScalarMultiplication(5.5)
		expected := Matrix{3, 2, [][]float64{{16.5, 5.5}, {11, 0}, {5.5, 11}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})

	t.Run("NegativeMultiplication", func(t *testing.T) {
		a := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}

		value := a.MatrixScalarMultiplication(-5.5)
		expected := Matrix{3, 2, [][]float64{{-16.5, -5.5}, {-11, 0}, {-5.5, -11}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})

	t.Run("SameMatrix", func(t *testing.T) {
		a := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}

		value := a.MatrixScalarMultiplication(1)
		expected := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})
}

func TestComponentWiseMultiplicationNormal(t *testing.T) {
	t.Run("ZeroMult", func(t *testing.T) {
		a := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}
		b := Matrix{3, 2, [][]float64{{0, 0}, {0, 0}, {0, 0}}}

		value := ComponentWiseMultiplication(a, b)
		expected := Matrix{3, 2, [][]float64{{0, 0}, {0, 0}, {0, 0}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})

	t.Run("MatrixWithMatrix", func(t *testing.T) {
		a := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}
		b := Matrix{3, 2, [][]float64{{1, 0}, {2, 1}, {3, 4}}}

		value := ComponentWiseMultiplication(a, b)
		expected := Matrix{3, 2, [][]float64{{3, 0}, {4, 0}, {3, 8}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})

	t.Run("VectorWithVector", func(t *testing.T) {
		a := Matrix{1, 3, [][]float64{{1, 0, 4}}}
		b := Matrix{1, 3, [][]float64{{1, 0, 4}}}

		value := ComponentWiseMultiplication(b, a)
		expected := Matrix{1, 3, [][]float64{{1, 0, 16}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})
}

func TestCompAllCloseNormal(t *testing.T) {
	t.Run("SameMatrix", func(t *testing.T) {
		a := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}
		b := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}

		value := CompAllClose(a, b, 1e-08)
		expected := true

		if value != expected {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})

	t.Run("AlmostSameMatrix", func(t *testing.T) {
		a := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}
		b := Matrix{3, 2, [][]float64{{2, 1}, {1, -1}, {1, 2}}}

		// matrices are different, but the tolerance is big
		value := CompAllClose(a, b, 1)
		expected := true

		if value != expected {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})

	t.Run("DifferentMatrices", func(t *testing.T) {
		a := Matrix{3, 2, [][]float64{{3, 1}, {2, 0}, {1, 2}}}
		b := Matrix{3, 2, [][]float64{{2, 1}, {1, -1}, {1, 2}}}

		value := CompAllClose(a, b, 1e-08)
		expected := false

		if value != expected {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})
}

func TestInverseNormal(t *testing.T) {
	t.Run("IdentityInverse", func(t *testing.T) {
		a := Matrix{4, 4, [][]float64{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}}

		value := a.Inverse()
		expected := Matrix{4, 4, [][]float64{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})

	t.Run("MatrixInverse", func(t *testing.T) {
		a := Matrix{4, 4, [][]float64{{1, 1, 1, 0}, {0, 3, 1, 2}, {2, 3, 1, 0}, {1, 0, 2, 1}}}

		value := a.Inverse()
		expected := Matrix{4, 4, [][]float64{{-3, -0.5, 1.5, 1}, {1, 0.25, -0.25, -0.5}, {3, 0.25, -1.25, -0.5}, {-3, 0, 1, 1}}}

		if value.N != expected.N || value.M != expected.M || !CompAllClose(value, expected, 1e-08) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})

	t.Run("DiagonalInverse", func(t *testing.T) {
		a := Matrix{4, 4, [][]float64{{65, 0, 0, 0}, {0, -72, 0, 0}, {0, 0, 19, 0}, {0, 0, 0, 11}}}

		value := a.Inverse()
		expected := Matrix{4, 4, [][]float64{{1.0 / 65.0, 0, 0, 0}, {0, -1.0 / 72.0, 0, 0}, {0, 0, 1.0 / 19.0, 0}, {0, 0, 0, 1.0 / 11.0}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})
}

func TestInverseDiagonalNormal(t *testing.T) {
	t.Run("IdentityInverse", func(t *testing.T) {
		a := Matrix{4, 4, [][]float64{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}}

		value := a.InverseDiagonal()
		expected := Matrix{4, 4, [][]float64{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})

	t.Run("DiagonalInverse", func(t *testing.T) {
		a := Matrix{4, 4, [][]float64{{65, 0, 0, 0}, {0, -72, 0, 0}, {0, 0, 19, 0}, {0, 0, 0, 11}}}

		value := a.InverseDiagonal()
		expected := Matrix{4, 4, [][]float64{{1.0 / 65.0, 0, 0, 0}, {0, -1.0 / 72.0, 0, 0}, {0, 0, 1.0 / 19.0, 0}, {0, 0, 0, 1.0 / 11.0}}}

		if !cmp.Equal(value, expected) {
			t.Errorf("Expected: %v ; Got: %v", expected, value)
		}
	})
}
