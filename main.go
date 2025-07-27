package main

import (
	"fmt"
	"math"
)

func calculate(num1,num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2,nil
	case "-":
		return  num1 - num2,nil
	case "*":
		return  num1 * num2,nil
	case "/":
		if num2 != 0 {
			return  num1 / num2,nil
		} else {
			return 0,fmt.Errorf("cannot divide by zero")
			
		}
	

	
case "%":
	if num2 != 0 {
		return (float64(int(num1) % int(num2))),nil
		} else {
			return 0,fmt.Errorf(" Cannot divide by zero for modulo")
			
		}
case "^":
	return math.Pow(num1, num2),nil

case "sqrt":
	if num1 >=0 {
		return math.Sqrt(num1),nil
	} else {
		return 0,fmt.Errorf("square root of negative number is not allowed")
		
	}
	
    
    default:
		return 0, fmt.Errorf("invalid operator")
		
	}
}

func main() {
	var history []string 
	for{
	var num1, num2 float64
	var operator string

	fmt.Println("Enter first number:")
	fmt.Scanln(&num1)

	if operator != "sqrt" {
		fmt.Println("Enter second number:")
		fmt.Scanln(&num2)
	} 


	fmt.Println("Enter operator (+, -, *, /,%,^, sqrt) or write 'exit' to quit,or write 'history',or write 'clear' to clear history: ")
	fmt.Scanln(&operator)

	if operator == "exit" {
		fmt.Println("Exiting , Goodbye!")
		break
	}

	if operator == "history" {
		if len(history) == 0{
			fmt.Printf("No Calculations Yet\n")
		} else {
			fmt.Println("History of previous calculations")
			for _,entry := range history {
				fmt.Println(entry)
			}
		}
		continue
	}

	if operator == "clear" {
		history = []string {}
			fmt.Println("History is cleared")
		continue
	}

	result,err :=calculate(num1,num2 ,operator)
	if err != nil {
		fmt.Println("Error:",err)
		continue
		} 

	 entry := fmt.Sprintf("%.2f, %s, %.2f = %.2f", num1,operator, num2, result )
	 if operator == "sqrt" {
		entry = fmt.Sprintf("sqrt(%.2f) = %.2f ", num1, result)
	 }
		
	 history = append(history, entry)

	fmt.Printf("Result: %.2f\n", result) 
}
}
