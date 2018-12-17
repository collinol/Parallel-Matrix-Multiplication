package evaluation

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"parallel_multiplication"
	"sequential_multiplication"
	"strconv"
	"strings"
	"structs"
)

func specificEvaluation(version string,
	m1 matrix_structs.Matrix, m2 matrix_structs.Matrix, threadcap int) matrix_structs.Matrix {
	if version == "parallel" {
		return parallel_multiplication.ParallelExecution(m1, m2, threadcap)
	}
	return sequential_multiplication.SequentialMultiply(m1, m2)

}

/************************
Main Evaluator Function that takes the matrix list,
the order in which the multiplications should be done,
and the version (parllel or sequential) that should be run.

Calculates the cumulative total product and writes to file.

*************************/

func Evaluate(version string,
	matrixList []matrix_structs.Matrix,
	order []string, outputFile string, threadCap int) {

	// keep a map of which multiplications we've completed based on substrings
	multiplicationsCompleted := make(map[string]matrix_structs.Matrix)
	outfile, err := os.Create(strings.Replace(outputFile, ".csv", "_Result.csv", -1)) //flag specific

	if err != nil {
		log.Println("failed opening ", outfile, err)
		panic(err.Error())
	}
	defer outfile.Close()
	writer := csv.NewWriter(outfile)
	defer writer.Flush()
	for m := 0; m < len(order); m++ {
		if strings.Count(order[m], "*") == 1 { //two independent matrices need to be multiplied together
			// before being joined
			m1Pos, _ := strconv.Atoi(strings.Split(order[m], "*")[0])
			m2Pos, _ := strconv.Atoi(strings.Split(order[m], "*")[1])
			m1 := matrixList[m1Pos]
			m2 := matrixList[m2Pos]
			multiplicationsCompleted[order[m]] = specificEvaluation(version, m1, m2, threadCap)

		} else {

			//https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
			//for entries across the string
			for split := 0; split < len(order[m]); split++ {
				//tracking the string like [0 <-n]
				// an "in" keyword would be cool, but this works...
				if matrix, exists := multiplicationsCompleted[order[m][0:len(order[m])-split-1]]; exists {
					var m1 matrix_structs.Matrix
					if strings.Count(order[m][len(order[m])-split-1:], "*") == 1 {
						getLastMatrix := strings.Split(order[m], "*")
						rightHandSide, _ := strconv.Atoi(string(getLastMatrix[len(getLastMatrix)-1]))
						m1 = matrixList[rightHandSide]
					} else {
						m1 = multiplicationsCompleted[order[m][len(order[m])-split:]]

					}
					multiplicationsCompleted[order[m]] = specificEvaluation(version, matrix, m1, threadCap)
					break //avoid catching smaller subsets
				} else {
					//tracking the string like [0-> n]
					if matrix, exists := multiplicationsCompleted[order[m][split:]]; exists {
						var m2 matrix_structs.Matrix
						if split == 2 {
							leftHandSide, _ := strconv.Atoi(string(order[m][0]))
							m2 = matrixList[leftHandSide]
						} else {

							m2 = multiplicationsCompleted[order[m][0:split-1]]
						}
						multiplicationsCompleted[order[m]] = specificEvaluation(version, m2, matrix, threadCap)

						break
						//^avoid smaller subsets
					}
				}
			}
		}
	}
	// This is where I'm writinng out the results to the files
	// If I had more time, I'd like to get run times without this
	// and see how much of a negative effect the I/O operations are having.
	result := multiplicationsCompleted[order[len(order)-1]]
	for row := 0; row < len(result.Data); row++ {
		data := result.Data[row]
		writeableData := []string{}
		for j := 0; j < len(data); j++ {

			writeableData = append(writeableData, fmt.Sprintf("%.6f", result.Data[row][j]))
		}
		writer.Write(writeableData)
	}
}
