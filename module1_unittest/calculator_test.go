package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	sum := Sum(7, 7)
	if sum != 14 {
		t.Error("30 was expected and we obtained", sum)
	}

	fmt.Println("Successful sum test")
}

func TestSubtract(t *testing.T) {
	subtract := Subtract(7, 7)
	if subtract != 0 {
		t.Error("30 was expected and we obtained", subtract)
	}

	fmt.Println("Successful subtraction test")
}

func TestMultiplication(t *testing.T) {
	multiplication := Multiplication(7, 7)
	if multiplication != 49 {
		t.Error("30 was expected and we obtained", multiplication)
	}

	fmt.Println("Successful multiplication test")
}

func TestDivision(t *testing.T) {
	division := Division(7, 7)
	if division != 1 {
		t.Error("30 was expected and we obtained", division)
	}

	fmt.Println("Successful division test")
}
