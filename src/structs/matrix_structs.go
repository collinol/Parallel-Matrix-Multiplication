package matrix_structs

// The (i,j)th element of matrices
type Element struct {
	Value float64
	I     int
	J     int
}

type Matrix struct {
	Rows int
	Cols int
	Data [][]float64
}
