package main

import "fmt"

func main() {
	running := true

	for running {
		intro()

		var selection int
		fmt.Scan(&selection)

		x, y := getNumbers()

		if selection == 1 {
			fmt.Printf("%d + %d = %d\n", x, y, add(x, y))
		} else if selection == 2 {
			fmt.Printf("%d - %d = %d\n", x, y, subtract(x, y))
		} else if selection == 3 {
			fmt.Printf("%d * %d = %d\n", x, y, multiply(x, y))
		} else if selection == 4 {
			fmt.Printf("%d / %d = %d\n", x, y, divide(x, y))
		} else if selection == 5 {
			fmt.Printf("%d mod %d = %d\n", x, y, modulus(x, y))
		} else {
			fmt.Printf("Not a possible selection.")
		}

		running = keepGoing()
	}
}

func intro() {
	fmt.Println("Welcome to the amazing calculator!")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Modulus")
	fmt.Print("Select a function:")
}

func getNumbers() (int, int) {
	fmt.Print("Enter the first number: ")
	var num1 int
	fmt.Scan(&num1)

	fmt.Print("Enter the second number: ")
	var num2 int
	fmt.Scan(&num2)

	return num1, num2
}

func keepGoing() bool {
	going := true

	fmt.Print("Do you want to go again? (y/n): ")
	var response string
	fmt.Scan(&response)

	if response == "n" {
		going = false
	}

	return going
}

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

func divide(x int, y int) int {
	return x / y
}

func modulus(x int, y int) int {
	return x % y
}
