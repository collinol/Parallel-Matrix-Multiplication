package csvGenerator

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/**********************************************
Keeping the chain length the same and randomizing the data entry

eg)
varying:
	fileType = "Large" will have matrices of dimensions between 500-1000 x 500-1000
	fileType = "Small" will have matrices of dimensions between 2-10 x 2-10
	fileType = "Medium" 10-500

reference:
http://golangcookbook.blogspot.com/2012/11/generate-random-number-in-given-range.html
https://golangcode.com/write-data-to-a-csv-file/
	***********************************************/
func CsvGenerator(fileType string) {

	numOfChains := 10

	for i := 0; i < numOfChains; i++ {
		n := -1
		m := -1
		file, err := os.Create("")
		switch fileType {
		case "Small":
			file, err = os.Create("csv_generator/Small/" + "csv" + strconv.Itoa(i) + ".csv")
		case "Large":
			file, err = os.Create("csv_generator/Large/" + "csv" + strconv.Itoa(i) + ".csv")
		case "Medium":
			file, err = os.Create("csv_generator/Medium/" + "csv" + strconv.Itoa(i) + ".csv")
		}

		if err != nil {
			log.Fatal("Cannot Create File", err)
		}

		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()
		rand.Seed(time.Now().UTC().UnixNano()) //https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly
		// ^ otherwise rand.Intn() produces the same random output every run
		lengthOfChain := 10
		for j := 0; j < lengthOfChain; j++ {
			switch fileType {
			/*initial dimensions can be anything, but after the first matrix
			the rows of the next matrix need to be the same as the columns of the previous one*/
			// pass pointers to dimensions to modify them depending on filetype
			case "Large":
				getDimensions(&n, &m, 500, 1000)
			case "Small":
				getDimensions(&n, &m, 2, 10)
			case "Medium":
				getDimensions(&n, &m, 10, 500)
			}
			newMatrix := make([][]string, n)
			for k := 0; k < n; k++ {
				newMatrix[k] = make([]string, m)
				for l := 0; l < m; l++ {
					newMatrix[k][l] = strconv.Itoa(int(rand.Float64() * 10))

				}
				writer.Write(newMatrix[k])
			}
			writer.Write([]string{"---"})

		}
	}

}

func getDimensions(n *int, m *int, min int, max int) {
	if *n == -1 {
		*n = rand.Intn(max-min) + min
	} else {
		*n = *m
	}
	*m = rand.Intn(max-min) + min

}
