package main

/*func main() {
	sum := Sum(2, 2)
	fmt.Println("The result of the sum is: ", sum)

	subtract := Subtract(2, 2)
	fmt.Println("The result of the subtraction is: ", subtract)

	multiplication := Multiplication(2, 2)
	fmt.Println("The result of the multiplication is: ", multiplication)

	division := Division(2, 2)
	fmt.Println("The result of the division is: ", division)
}*/

func Sum(number1, number2 int) (result int) {
	result = number1 + number2
	return
}

func Subtract(number1, number2 int) (result int) {
	result = number1 - number2
	return
}

func Multiplication(number1, number2 int) (result int) {
	result = number1 * number2
	return
}

func Division(number1, number2 int) (result int) {
	result = number1 / number2
	return
}
