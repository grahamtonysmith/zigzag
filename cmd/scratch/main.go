package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

const (
	FILEPATH = "/home/grahamtonysmith/code/go-alg-zigzag/data/gc.csv"
)

func main() {
	f, err := os.Open(FILEPATH)
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
