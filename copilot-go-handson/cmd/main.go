package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func readCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func main() {
	records, err := readCSV("../assets/accounts.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, record := range records {
		fmt.Println(record)
	}
}
