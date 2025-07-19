package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Enter first number:")
	fmt.Scanln(&num1)

	fmt.Println("Enter second number:")
	fmt.Scanln(&num2)

	fmt.Println("Enter operator (+, -, *, /,%,^):")
	fmt.Scanln(&operator)

	var result float64 

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Error: Cannot divide by zero")
			return
		}
	

	
case "%":
	if num2 != 0 {
		result = (float64(int(num1) % int(num2)))
		} else {
			fmt.Println("Error: Cannot divide by zero for modulo.")
			return
		}
case "^":
	result = math.Pow(num1, num2)
    
    default:
		fmt.Println("Invalid Operator.")
		return
	}


	fmt.Printf("Result: %.2f\n", result) 
}
