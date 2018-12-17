package sequential_multiplication

import (
	"structs"
)

func SequentialMultiply(m1 matrix_structs.Matrix, m2 matrix_structs.Matrix) matrix_structs.Matrix {
	n := m1.Rows
	m := m1.Cols // or m2.Rows (doesn't matter)
	p := m2.Cols
	// a (nxm)*(mxp) matrix multiplication operation = (nxp) matrix
	m3 := matrix_structs.Matrix{Rows: n, Cols: p, Data: make([][]float64, n)}
	for x := 0; x < m3.Rows; x++ {
		m3.Data[x] = make([]float64, p) // initialize
	}
	for i := 0; i < n; i++ {
		for j := 0; j < p; j++ {
			sum := 0.
			for k := 0; k < m; k++ {
				sum = sum + m1.Data[i][k]*m2.Data[k][j]

			}
			m3.Data[i][j] = sum

		}
	}

	return m3
}
