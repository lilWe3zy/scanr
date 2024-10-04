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
		validConnections [][]string
		err              error
	)

	var maxSig = flag.Int("s", -55, "Provide a maximum allowed signal value")
	var file = flag.String("f", "t.csv", "Input csv file to process")
	var out = flag.String("o", "scanlist.csv", "Output csv file")
	var sub = flag.String("n", "RN", "Substring to match")

	flag.Parse()

	fd, err := os.Open(*file)
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
		isOwned := strings.Contains(towerName, *sub)
		signal, recordError := strconv.Atoi(connection[3])

		if recordError != nil || signal <= *maxSig || !isOwned {
			continue
		}

		validConnections = append(validConnections, connection)
	}

	/* New file creation */
	wfd, err := os.Create(*out)
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
