package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type AccountType string

const (
	Checking AccountType = "Checking"
	Savings  AccountType = "Savings"
)

type Currency string

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
)

type Account struct {
	AccountID         string
	AccountHolderName string
	Balance           float64
	Currency          Currency
	Type              AccountType
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
			Currency:          Currency(record[3]),
			Type:              AccountType(record[4]),
			CreatedAt:         createdAt,
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func sumBalancesByYearAndCurrency(accounts []Account) map[string]map[Currency]float64 {
	sums := make(map[string]map[Currency]float64)

	for _, account := range accounts {
		year := account.CreatedAt.Year()
		if _, ok := sums[strconv.Itoa(year)]; !ok {
			sums[strconv.Itoa(year)] = make(map[Currency]float64)
		}
		sums[strconv.Itoa(year)][account.Currency] += account.Balance
	}

	return sums
}

func writeCSV(filePath string, sums map[string]map[Currency]float64) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"Year", "Currency", "Sum of Balances"})

	// Write data
	for year, currencySums := range sums {
		for currency, sum := range currencySums {
			record := []string{year, string(currency), fmt.Sprintf("%.2f", sum)}
			writer.Write(record)
		}
	}

	return nil
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

	sums := sumBalancesByYearAndCurrency(accounts)
	fmt.Println("Sum of balances by year and currency:")
	for year, currencies := range sums {
		for currency, sum := range currencies {
			fmt.Printf("Year: %s, Currency: %s, Sum: %.2f\n", year, currency, sum)
		}
	}

	err = writeCSV("../assets/sum_balances.csv", sums)
	if err != nil {
		fmt.Println("Error writing CSV:", err)
	}
}
