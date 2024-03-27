package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Define and parse input flags
	var signalMax int
	var fileName string
	var outDir string

	flag.IntVar(&signalMax, "signal", -55, "Provide a maximum allowed signal value")
	flag.StringVar(&fileName, "file", "t.csv", "Input csv file to process")
	flag.StringVar(&outDir, "outdir", "scanlist.csv", "Output csv file")

	flag.Parse()
	// Create final record structure
	var validConnections [][]string

	fd, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func(fd *os.File) {
		err := fd.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(fd)

	// read CSV file
	fileReader := csv.NewReader(fd)

	records, err := fileReader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	// Loop through csv slices
	for _, v := range records {
		towerName := v[1]
		isOwned := strings.Contains(towerName, "RN")
		signal, err := strconv.Atoi(v[3])

		if err != nil || signal <= signalMax || !isOwned {
			continue
		}
		// Append valid slices to new slice, for file writing
		validConnections = append(validConnections, v)
	}

	// Create new scan list csv file to write to
	wfd, err := os.Create(outDir)
	if err != nil {
		panic(err)
	}
	defer func(wfd *os.File) {
		err = wfd.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(wfd)

	fileWriter := csv.NewWriter(wfd)
	err = fileWriter.WriteAll(validConnections)
	if err != nil {
		fmt.Println("Could not write to file")
	}

	fmt.Printf("%d sectors identified\n", len(validConnections))
}
