package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

var csvPath string

func init() {

	flag.StringVar(
		&csvPath,
		"csv",
		"problems.csv",
		"a csv file in the format of 'question,answer' (default \"problems.csv\")",
	)
}

func main() {

	file, err := os.Open(csvPath)

	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Ошибка чтения: %v\n", err)
		return
	}

	for _, record := range records {
		fmt.Printf("Запись: %v\n", record)
	}
	// Parse flag with csv file name, if empty read default 'problems.csv'

	// Read csv file
	// Save question in map [question]answer

	// initialize map with stats (questions size, right answers, wrong answers)

	// Loop questions
	// Ask N question. If right or wrong save in stats

	// Print results like "You scored N out of len(questions)"
}
