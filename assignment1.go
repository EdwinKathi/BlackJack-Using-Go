package main

import (
	"fmt"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	if b != 0 {
		return a / b
	}
	return 0
}

func main() {
	var num1, num2 float64

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	fmt.Println("Select an operation:")
	fmt.Println("1. Addition (+)")
	fmt.Println("2. Subtraction (-)")
	fmt.Println("3. Multiplication (*)")
	fmt.Println("4. Division (/)")

	var choice int
	fmt.Scanln(&choice)

	var result float64

	switch choice {
	case 1:
		result = add(num1, num2)
		fmt.Printf("The sum of %.2f and %.2f is %.2f\n", num1, num2, result)
	case 2:
		result = subtract(num1, num2)
		fmt.Printf("The difference between %.2f and %.2f is %.2f\n", num1, num2, result)
	case 3:
		result = multiply(num1, num2)
		fmt.Printf("The product of %.2f and %.2f is %.2f\n", num1, num2, result)
	case 4:
		result = divide(num1, num2)
		if num2 != 0 {
			fmt.Printf("The division of %.2f by %.2f is %.2f\n", num1, num2, result)
		} else {
			fmt.Println("Cannot divide by zero!")
		}
	default:
		fmt.Println("Invalid choice!")
	}
}
