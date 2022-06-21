//go:build unit

package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	sum := Sum(7, 7)
	if sum != 14 {
		t.Error("expected 14 but obtained", sum)
	}

	fmt.Println("Successful sum test")
}

func TestSubtract(t *testing.T) {
	subtract := Subtract(7, 7)
	if subtract != 0 {
		t.Error("expected 0 but obtained", subtract)
	}

	fmt.Println("Successful subtraction test")
}

func TestMultiplication(t *testing.T) {
	multiplication := Multiplication(7, 7)
	if multiplication != 49 {
		t.Error("expected 49 but obtained", multiplication)
	}

	fmt.Println("Successful multiplication test")
}

func TestDivision(t *testing.T) {
	division := Division(7, 7)
	if division != 1 {
		t.Error("expected 1 but obtained", division)
	}

	fmt.Println("Successful division test")
}

func TestAverage(t *testing.T) {
	average := Average(7, 7, 7, 7, 7)
	if average != 7 {
		t.Error("expected 7 but obtained", average)
	}

	fmt.Println("Successful average test")
}
