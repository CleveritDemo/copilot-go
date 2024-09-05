package main

import (
	"reflect"
	"testing"
	"time"
)

func TestSumBalancesByYearAndCurrency(t *testing.T) {
	// Create sample data
	accounts := []Account{
		{AccountID: "1", AccountHolderName: "John Doe", Balance: 1000.50, Currency: USD, Type: Checking, CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)},
		{AccountID: "2", AccountHolderName: "Jane Smith", Balance: 2000.75, Currency: EUR, Type: Savings, CreatedAt: time.Date(2021, 6, 15, 0, 0, 0, 0, time.UTC)},
		{AccountID: "3", AccountHolderName: "Alice Brown", Balance: 1500.00, Currency: USD, Type: Savings, CreatedAt: time.Date(2022, 3, 10, 0, 0, 0, 0, time.UTC)},
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
