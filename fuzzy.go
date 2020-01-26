package main

import (
	"encoding/csv"
	"os"
	"strconv"
    "github.com/paul-mannino/go-fuzzywuzzy"
)

func readCsv(file string) [][]string {
	csvFile, _ := os.Open(file)
	defer csvFile.Close()
	r := csv.NewReader(csvFile)
	records, _ := r.ReadAll()
	return records
}

func main() {
    
	original:= readCsv("constituents.csv")
	target:= readCsv("sp500.csv")
	csvOut, _ := os.Create("results_go.csv")
	csvwriter := csv.NewWriter(csvOut)

	var targetData []string

	for _, line := range target[1:] {	
		targetData = append(targetData, line[1])
	}
	csvwriter.Write([]string{"name", "match_name", "match_score"})
	for _, line := range original[1:] {
		match, _ := fuzzy.ExtractOne(line[1], targetData)
		csvwriter.Write([]string{line[1], match.Match, strconv.Itoa(match.Score)})
    }
}
		