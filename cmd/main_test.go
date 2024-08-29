package main

import (
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"
)

func TestReadCSV(t *testing.T) {
	// Create a temporary CSV file
	file, err := os.CreateTemp("", "accounts*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	// Write sample data to the file
	data := `AccountID,AccountHolderName,Balance,Currency,Type,CreatedAt
1,John Doe,1000.50,USD,Checking,2022-01-01
2,Jane Smith,2000.75,EUR,Saving,2021-06-15`
	if _, err := file.Write([]byte(data)); err != nil {
		t.Fatal(err)
	}
	file.Close()

	// Call readCSV
	accounts, err := readCSV(file.Name())
	if err != nil {
		t.Fatal(err)
	}

	// Verify the returned data
	expected := []Account{
		{AccountID: "1", AccountHolderName: "John Doe", Balance: 1000.50, Currency: USD, Type: Checking, CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)},
		{AccountID: "2", AccountHolderName: "Jane Smith", Balance: 2000.75, Currency: EUR, Type: Saving, CreatedAt: time.Date(2021, 6, 15, 0, 0, 0, 0, time.UTC)},
	}
	if !reflect.DeepEqual(accounts, expected) {
		t.Errorf("Expected %v, got %v", expected, accounts)
	}
}

func TestSumBalancesByYearAndCurrency(t *testing.T) {
	// Create sample data
	accounts := []Account{
		{AccountID: "1", AccountHolderName: "John Doe", Balance: 1000.50, Currency: USD, Type: Checking, CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)},
		{AccountID: "2", AccountHolderName: "Jane Smith", Balance: 2000.75, Currency: EUR, Type: Saving, CreatedAt: time.Date(2021, 6, 15, 0, 0, 0, 0, time.UTC)},
		{AccountID: "3", AccountHolderName: "Alice Brown", Balance: 1500.00, Currency: USD, Type: Saving, CreatedAt: time.Date(2022, 3, 10, 0, 0, 0, 0, time.UTC)},
	}

	// Call sumBalancesByYearAndCurrency
	sums := sumBalancesByYearAndCurrency(accounts)

	// Verify the returned data
	expected := map[string]map[Currency]float64{
		"2021": {EUR: 2000.75},
		"2022": {USD: 2500.50},
	}
	if !reflect.DeepEqual(sums, expected) {
		t.Errorf("Expected %v, got %v", expected, sums)
	}
}

func TestWriteCSV(t *testing.T) {
	// Create a temporary file
	file, err := os.CreateTemp("", "sums*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	// Create sample data
	sums := map[string]map[Currency]float64{
		"2021": {EUR: 2000.75},
		"2022": {USD: 2500.50},
	}

	// Call writeCSV
	if err := writeCSV(file.Name(), sums); err != nil {
		t.Fatal(err)
	}

	// Read the file and verify its contents
	content, err := os.ReadFile(file.Name())
	if err != nil {
		t.Fatal(err)
	}

	expected := `Year,Currency,Sum of Balances
2021,EUR,2000.75
2022,USD,2500.50
`

	// Sort the lines before comparing
	expectedLines := strings.Split(expected, "\n")
	actualLines := strings.Split(string(content), "\n")
	sort.Strings(expectedLines)
	sort.Strings(actualLines)

	if !reflect.DeepEqual(expectedLines, actualLines) {
		t.Errorf("Expected %v, got %v", expectedLines, actualLines)
	}
}
