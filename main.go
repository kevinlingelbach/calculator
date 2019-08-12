package main

import "fmt"

func main() {
	fmt.Println("Welcome to the amazing calculator!")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide \n\n")
	fmt.Print("Select a function:")

	var selection int
	fmt.Scan(&selection)

	x, y := getNumbers()
	fmt.Println(x, " + ", y, " = ", add(x, y))
}

func getNumbers() (int, int) {
	fmt.Print("Enter number 1: ")
	var num1 int
	fmt.Scan(&num1)

	fmt.Print("Enter number 2: ")
	var num2 int
	fmt.Scan(&num2)

	return num1, num2
}

func add(x int, y int) int {
	return x + y
}
