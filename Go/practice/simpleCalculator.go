package main

import "fmt"

func main() {
	
	var num1,num2 int
	var operator string 

	fmt.Println("Enter a number")
	fmt.Scan(&num1)

	fmt.Println("Enter another number")
	fmt.Scan(&num2)

	fmt.Println("Enter operator (+,-,*,/)")
	fmt.Scan(&operator)

	switch operator{
	case "+":
		fmt.Println("Outpur",num1+num2)
	case "-":
		fmt.Println("Output",num1-num2)
	case "*":
		fmt.Println("Output",num1*num2)
	case "/":
		if num2!=0{
			fmt.Println("Output",num1/num2)
		}else{
			fmt.Println("Error: division by zero")
		}
	default:
		fmt.Println("Invalid operator")
	}

}