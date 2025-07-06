package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	signal int
	ssid   string
)

func init() {
	const (
		defaultSignal = -55
		signalDesc    = "Signal in dBm to set as minimum viable signal"
		defaultSSID   = ""
		ssidDesc      = "The SSID substring to match when filtering for connections"
	)

	flag.IntVar(&signal, "signal", defaultSignal, signalDesc)
	flag.IntVar(&signal, "s", defaultSignal, signalDesc)
	flag.StringVar(&ssid, "ssid", defaultSSID, ssidDesc)
	flag.StringVar(&ssid, "id", defaultSSID, ssidDesc)
}

func main() {
	var (
		inputFilePath  = flag.String("in", "", "input file path")
		outputFilePath = flag.String("out", "output.csv", "output file path")
	)

	flag.Parse()

	// Flag validation
	// The inputFilePath variable is required,
	// and no default is set, as the name can change depending on what the name was set on the router itself.
	if inputFilePath == nil || *inputFilePath == "" {
		flag.Usage()
		log.Fatalln("Missing input file path, aborting")
	}

	file, err := os.Open(*inputFilePath)
	if err != nil {
		log.Fatalf("Error opening input file: %v", err)
	}
	defer func(fd *os.File) {
		err = fd.Close()
		if err != nil {
			log.Fatalf("Error closing input file: %v", err)
		}
	}(file)

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading csv entries: %v", err)
	}

	var validConnections [][]string
	for _, conn := range records {
		id := conn[1]
		isOwned := strings.Contains(id, ssid)
		sig, recordError := strconv.Atoi(conn[3])

		if recordError != nil || sig <= signal || !isOwned {
			continue
		}

		validConnections = append(validConnections, conn)
	}

	/* New file creation */
	wfd, err := os.Create(*outputFilePath)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer func(wfd *os.File) {
		err = wfd.Close()
		if err != nil {
			log.Fatalf("Error closing output file: %v", err)
		}
	}(wfd)

	fileWriter := csv.NewWriter(wfd)

	if err = fileWriter.WriteAll(validConnections); err != nil {
		log.Fatalf("Error writing csv entries: %v", err)
	}

	fmt.Printf("%d sectors identified\n", len(validConnections))
	return
}
