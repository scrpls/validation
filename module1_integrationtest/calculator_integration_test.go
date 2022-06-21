//go:build integration

package main

import "testing"
import "fmt"
import "github.com/joho/godotenv"
import "os"
import "strconv"

func TestIntegrationCalculator(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatalf("could not load .env file: %v", err)
	}

	value_use, err := strconv.Atoi(os.Getenv("value_use"))
	average_total, err := strconv.Atoi(os.Getenv("average_total"))

	sum := Sum(value_use, value_use)
	subtract := Subtract(value_use, value_use)
	multiplication := Multiplication(value_use, value_use)
	division := Division(value_use, value_use)
	average := Average(value_use, value_use, value_use, value_use, value_use)

	totalaverage := Average(sum, subtract, multiplication, division, average)

	if totalaverage != average_total {
		t.Error("Expected 17 but obtained", totalaverage)
	}

	fmt.Println("Successful total average integration test")
}
