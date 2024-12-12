package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func readCsvFile(filepath string) [][]string {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Unable to read input file "+filepath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filepath, err)
	}
	return records
}
func main() {
	records := readCsvFile("problems.csv")
	fmt.Println(records)
}
