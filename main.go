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
	defer fd.Close()

	// read CSV file
	fileReader := csv.NewReader(fd)

	records, err := fileReader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range records {
		towerName := v[1]
		isOwned := strings.Contains(towerName, "RN")
		signal, err := strconv.Atoi(v[3])

		if err != nil {
			fmt.Println("Could not convert to integer")
			continue
		}

		if signal <= signalMax || !isOwned {
			continue
		}
		validConnections = append(validConnections, v)
	}

	// Create new scanlist csv file to write to
	w, err := os.Create(outDir)
	if err != nil {
		fmt.Println("Could not create file")
	}
	defer w.Close()

	fileWriter := csv.NewWriter(w)
	err = fileWriter.WriteAll(validConnections)
	if err != nil {
		fmt.Sprintln("Could not write to file")
	}

	fmt.Println("Sectors filtered")
}
