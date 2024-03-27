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
	var (
		signalMax        int
		fileName         string
		outDir           string
		name             string
		validConnections [][]string
		err              error
	)

	flag.IntVar(&signalMax, "s", -55, "Provide a maximum allowed signal value")
	flag.StringVar(&fileName, "f", "t.csv", "Input csv file to process")
	flag.StringVar(&outDir, "o", "scanlist.csv", "Output csv file")
	flag.StringVar(&name, "n", "RN", "Substring to match")
	flag.Parse()

	fd, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	/* Explicit ignore with File closure */
	defer func(fd *os.File) {
		_ = fd.Close()
	}(fd)

	fileReader := csv.NewReader(fd)

	records, err := fileReader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, connection := range records {
		towerName := connection[1]
		isOwned := strings.Contains(towerName, name)
		signal, recordError := strconv.Atoi(connection[3])

		if recordError != nil || signal <= signalMax || !isOwned {
			continue
		}

		validConnections = append(validConnections, connection)
	}

	/* New file creation */
	wfd, err := os.Create(outDir)
	if err != nil {
		panic(err)
	}
	defer func(wfd *os.File) {
		_ = wfd.Close()
	}(wfd)

	fileWriter := csv.NewWriter(wfd)

	err = fileWriter.WriteAll(validConnections)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d sectors identified\n", len(validConnections))
	return
}
