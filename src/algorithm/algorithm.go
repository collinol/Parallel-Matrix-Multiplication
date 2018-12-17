package algorithm

import (
	"strconv"
	"strings"
	"structs"
)

// Reference for MatrixChain function
//https://labs.xjtudlc.com/labs/wldmt/reading%20list/books/Algorithms%20and%20optimization/Introduction%20to%20Algorithms.pdf
// ^(pages 375-377)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func MatrixChain(listOfMatrices []matrix_structs.Matrix) []string {

	n := len(listOfMatrices) + 1
	/* ***************************** /
	convert matrices into list of dimensions, listDim
	where ith matrix Ai is of dimension listDim[i-1] x listDim[i] 	*/
	listDim := []int{}
	for m := 0; m < n-1; m++ {
		if m == 0 {
			listDim = append(listDim, listOfMatrices[m].Rows)
			// append # of rows of left most matrix
		}
		listDim = append(listDim, listOfMatrices[m].Cols)
		// append # of cols of every matrix to the end of the list

	}
	/********************************/
	costTable := make([][]int, n)
	brackets := make([][]int, n)
	// initialize cost table and brackets table (for keeping track of order)
	for x := 0; x < n; x++ {
		costTable[x] = make([]int, n)
		costTable[x][x] = 0
		brackets[x] = make([]int, n)
	}
	counter := 1
	for l := 2; l < n; l++ {
		for i := 1; i < (n - l + 1); i++ {
			j := i + l - 1
			costTable[i][j] = MaxInt
			for k := i; k < j; k++ {

				cost := costTable[i][k] + costTable[k+1][j] +
					listDim[i-1]*listDim[k]*listDim[j]

				if cost < costTable[i][j] {
					costTable[i][j] = cost
					brackets[i][j] = k
					counter++

				}

			}
		}
	}
	var result string
	// Initially I thought I was going to be using chains of length >26
	// I intended to name them mod26 style, AA, AB, AC... BA, BB etc
	// however, this proved too difficult when parsing the resulting string
	// also, having a chain of that length GREATLY increased the run time
	matrixLetters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	pos := 0
	getMultiplicationOrder(1, n-1, n, brackets, matrixLetters, &pos, &result)
	/* ^ returns a string parenthesizing order of operations, but we need to
	be able to return a list containing the index of the matrices in the matrix list,
	in the correct order to feed them into the multiplication function*/
	pairings := []string{}
	order := result
	stack := []string{}
	index := -1
	// Using a stack technique with an index to annotate
	// the result string with proper indices so that I can slice
	// my array of matrices
	for i := 0; i < len(order); i++ {
		char := string(order[i])
		if char != "(" && char != ")" && char != "*" {
			index++
		}
		if char == "(" {
			stack = append(stack, strconv.Itoa(i))
		} else {
			if char == ")" && len(stack) > 0 {
				start, _ := strconv.Atoi(stack[len(stack)-1])
				stack = stack[0 : len(stack)-1]
				strippedString := strings.Replace(string(order[start+1:i]), "*", "", -1)
				strippedString = strings.Replace(strippedString, "(", "", -1)
				strippedString = strings.Replace(strippedString, ")", "", -1)
				pairings = append(pairings, strippedString)
			}
		}
	}
	strippedOrder := strings.Replace(order, "*", "", -1)
	strippedOrder = strings.Replace(strippedOrder, "(", "", -1)
	strippedOrder = strings.Replace(strippedOrder, ")", "", -1)
	orderIndices := []string{}
	startPosition := -1
	for i := 0; i < len(pairings); i++ {
		nextMult := pairings[i]

		startPosition = strings.Index(strippedOrder, nextMult)
		endPosition := startPosition + len(nextMult) - 1
		indexRange := ""
		for i := startPosition; i <= endPosition; i++ {
			if i != endPosition {
				indexRange += strconv.Itoa(i) + "*"
			} else {
				indexRange += strconv.Itoa(i)
			}
		}
		orderIndices = append(orderIndices, indexRange)

	}
	return orderIndices
}

/*https://www.geeksforgeeks.org/printing-brackets-matrix-chain-multiplication-problem/
	 A C++ reference I used for recursively printing parenthesis ^^
I had some issue trying to use this function to return a string, representing the
multiplication order, because of the recursion. So used a string ptr in the algorithm
function to save the results.  */
func getMultiplicationOrder(i int, j int, n int,
	bracket [][]int, mat []string, pos *int, result *string) {
	// If only one matrix left in current segment
	if i == j {
		if *pos < 26 {
			(*result) += mat[*pos]
		} else {
			repeatedChar := ""
			for l := 0; l < (*pos / 26); l++ {
				repeatedChar = mat[l] + mat[(*pos)%26]
			}
			(*result) += repeatedChar
		}
		*pos++
		(*result) += "*"
		return
	}

	(*result) += "("

	// Recursively put brackets around subexpression
	// from i to bracket[i][j].

	getMultiplicationOrder(i, bracket[i][j], n, bracket, mat, pos, result)

	// Recursively put brackets around subexpression
	// from bracket[i][j] + 1 to j.
	getMultiplicationOrder(bracket[i][j]+1, j, n, bracket, mat, pos, result)
	(*result) += ")"
}
