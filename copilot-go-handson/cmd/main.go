package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Account struct {
	AccountID         string
	AccountHolderName string
	Balance           float64
	Currency          string
	Type              string
	CreatedAt         time.Time
}

func readCSV(filePath string) ([]Account, error) {
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

	var accounts []Account
	for _, record := range records[1:] { // Skip header row
		balance, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, err
		}
		createdAt, err := time.Parse("2006-01-02", record[5])
		if err != nil {
			return nil, err
		}
		account := Account{
			AccountID:         record[0],
			AccountHolderName: record[1],
			Balance:           balance,
			Currency:          record[3],
			Type:              record[4],
			CreatedAt:         createdAt,
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func main() {
	accounts, err := readCSV("../assets/accounts.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, account := range accounts {
		fmt.Printf("%+v\n", account)
	}
}
