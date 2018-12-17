package main

import (
	"algorithm"
	"bufio"
	"csv_generator"
	"evaluation"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"structs"
	"time"
)

func convertCsvToMatrices(linesFromCsv [][]string) matrix_structs.Matrix {
	n := len(linesFromCsv)
	m := len(linesFromCsv[0])
	newMatrix := matrix_structs.Matrix{Rows: n, Cols: m, Data: make([][]float64, n)}
	for i := 0; i < n; i++ {
		newMatrix.Data[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			data, _ := strconv.ParseFloat(linesFromCsv[i][j], 64)
			newMatrix.Data[i][j] = data
		}
	}
	return newMatrix
}

func getOrderOfOperations(algorithmFlag bool, list []matrix_structs.Matrix) []string {
	if algorithmFlag {
		return algorithm.MatrixChain(list)
	}
	// manually create the order of operations that's just
	//0*1, then *2, then *3, then *4... then *n
	leftToRight := []string{"0*1"}
	previous := 0
	for i := 2; i < len(list); i++ {
		leftToRight = append(leftToRight, leftToRight[previous]+"*"+strconv.Itoa(i))
		previous++
	}
	return leftToRight
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--help" {
		fmt.Println("Please specify flags")
		fmt.Println(" '-p = ' Upper bound on number of threads allowed to run")
		fmt.Println(" '-f(n) = ' file size you want to evaluate")
		fmt.Println(" 		available options: s = small matrices (dimensions between 2 and 10")
		fmt.Println(" 		                   m = medium matrices (dimensions between 10 and 500")
		fmt.Println(" 		                   l = large matrices (dimensions between 500 and 1000")
		fmt.Println("  *** if files of that size have already been created, include the n like -fn")
		fmt.Println("  *** this will tell the program to bypass creating new files for that directory. -f will populate the directory")
		fmt.Println(" '-a' Don't need to set equal to anything. If flag is present, the program will" +
			"run the matix chain multiplication algorithm to determine the most effecient order of operations" +
			"if not, the program will just execute each multiplication in order")
		fmt.Println("Note: if you don't already have this files in the necessary directory, this program will" +
			"create them for you first, which will take a little longer")
		fmt.Println("if you are getting a Failed to Open / No such directory or file error, " +
			"please check that there exists the directory structure /src/csv_generator/Small, and likewise for" +
			" Medium and Large.  ")
		os.Exit(3)
	}
	optimize := false
	version := "sequential"
	threadCap := -1
	var directory string
	size := "s"
	needsToBeCreated := true
	for args := 0; args < len(os.Args); args++ {
		if strings.Contains(os.Args[args], "-a") {
			optimize = true
		}
		if strings.Contains(os.Args[args], "-p") {
			version = "parallel"
			threadCap, _ = strconv.Atoi(strings.Split(os.Args[args], "=")[1])
			if threadCap <= 2 {
				version = "sequential"
			}
		}
		if strings.Contains(os.Args[args], "-f") {
			size = strings.Split(os.Args[args], "=")[1]
			if strings.Contains(os.Args[1], "-fn") {
				needsToBeCreated = false
			}

		}

	}
	switch size {
	case "s":
		directory = "Small"
	case "m":
		directory = "Medium"
	case "l":
		directory = "Large"

	}
	if needsToBeCreated {
		fmt.Println("Populating Matrix Chain Csvs")
		csvGenerator.CsvGenerator(directory)
	}
	timing(version, directory, threadCap, optimize)
}

func timing(version string, directory string, threadcap int, optimize bool) {
	totalDirectoryTime := []float64{}
	for csvFileNumber := 0; csvFileNumber < 10; csvFileNumber++ {

		infile, err := os.Open("csv_generator/" + directory + "/csv" + strconv.Itoa(csvFileNumber) + ".csv") //flag specific

		if err != nil {
			log.Println("failed opening ", infile, err)
			panic(err.Error())
		}
		defer infile.Close()
		scanner := bufio.NewScanner(infile)
		entries := [][]string{}
		list := []matrix_structs.Matrix{}
		//for each file in the directory, convert each block to matrix and append to list
		for scanner.Scan() {
			inputLine := strings.Split(scanner.Text(), ",")
			if inputLine[0] != "---" {
				entries = append(entries, inputLine)
			} else {
				list = append(list, convertCsvToMatrices(entries))
				entries = [][]string{}
			}

		}

		outputFile := "csv_generator/" + directory + "/csv" + strconv.Itoa(csvFileNumber) + ".csv"

		times := []float64{}
		// get average run time for 5 runs
		for iteration := 0; iteration < 5; iteration++ {
			start := time.Now()
			badorder := getOrderOfOperations(true, list)
			evaluation.Evaluate(version, list, badorder, outputFile, threadcap)
			times = append(times, time.Now().Sub(start).Seconds())
		}
		sum := 0.
		for avg := 0; avg < 5; avg++ {
			sum += times[avg]
		}
		totalDirectoryTime = append(totalDirectoryTime, sum/5.)

	}
	t := 0.
	for avg := 0; avg < 5; avg++ {
		t += totalDirectoryTime[avg]
	}
	fmt.Println("Directory:", directory, version, "time with optimize =", optimize, ":",
		t/10., "Upper Bound on threads =", threadcap)
}
