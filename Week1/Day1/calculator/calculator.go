package calculator
import "fmt"

func Calculator(){
	var a , b float64
	var operator string

	fmt.Println("Enter a number")
	fmt.Scan(&a)

	fmt.Println("Enter another number")
	fmt.Scan(&b)

	fmt.Println("Enter operator")
	fmt.Scan(&operator)

	switch(operator){
	case "+":
		fmt.Println("Result",a+b)
	case "-":
		fmt.Println("Result",a-b)
	case "*":
		fmt.Println("Result",a*b)
	case "/":
		if b == 0{
			fmt.Println("Cannot divide by zero")
		}else{
			fmt.Println("Result",a/b)
		}
	default :
		fmt.Println("Please enter a valid operator")
	}
}