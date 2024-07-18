package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	PATH string
)

func init() {
	flag.StringVar(&PATH, "path", "/path/to/csv", "filepath to csv")
	flag.Parse()
}

func main() {
	f, err := os.Open(PATH)
	if err != nil {
		log.Fatalln("Unable to read CSV", err)
	}
	defer f.Close()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse CSV", err)
	}

	for i, record := range records {
		log.Println(fmt.Sprintf("[RECORD %d]", i), record)
	}
}
