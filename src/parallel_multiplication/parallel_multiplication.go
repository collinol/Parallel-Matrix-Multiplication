package parallel_multiplication

import (
	"structs"
	"sync"
)

func ParallelExecution(m1 matrix_structs.Matrix, m2 matrix_structs.Matrix, threadcap int) matrix_structs.Matrix {
	ElementResults := make(chan matrix_structs.Element)
	n := m1.Rows
	p := m2.Cols
	// a (nxm)*(mxp) matrix multiplication operation = (nxp) matrix
	m3 := matrix_structs.Matrix{Rows: n, Cols: p, Data: make([][]float64, n)}
	wg := sync.WaitGroup{}
	wg.Add(1)
	// one routine to listen for new results and write to the product
	go feedResults(m3, ElementResults, &wg)

	for x := 0; x < m3.Rows; x++ {
		m3.Data[x] = make([]float64, p)
	}
	capChannel := make(chan bool, threadcap)
	for i := 0; i < n; i++ {
		capChannel <- true
		go parallelMultiply(m1, m2, i, p, ElementResults, capChannel)
		// thrad cap wait barrier
		for len(capChannel) > threadcap {
			continue
		}
	}

	wg.Wait()
	return m3
}

// the actual calculation
func parallelMultiply(
	m1 matrix_structs.Matrix, m2 matrix_structs.Matrix,
	i int, p int, Elements chan<- matrix_structs.Element, cap chan bool) {

	m := m1.Cols
	for j := 0; j < p; j++ {
		sum := 0.
		for k := 0; k < m; k++ {
			sum = sum + m1.Data[i][k]*m2.Data[k][j]
		}

		Elements <- matrix_structs.Element{Value: sum, I: i, J: j}
	}
	<-cap

}

func feedResults(m3 matrix_structs.Matrix, readFrom <-chan matrix_structs.Element, wg *sync.WaitGroup) {
	for index := 0; index < m3.Rows*m3.Cols; index++ {
		input := <-readFrom
		m3.Data[input.I][input.J] = input.Value
	}
	wg.Done()
}
