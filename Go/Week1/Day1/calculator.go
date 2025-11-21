package main
import "fmt"
func main(){
	var num1 float64
	fmt.Println("Enter first number")
	fmt.Scan(&num1)
	
	var num2 float64
	fmt.Println("Enter second number")
	fmt.Scan(&num2)
	
	var operator string
	fmt.Println("Enter the operaot (+,-,*,/)")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Printf("Result : %f",num1+num2)
	case "-":
		fmt.Printf("Result : %f",num1-num2)
	case "*":
		fmt.Printf("Result : %f",num1*num2)
	case "/":
		if num2 != 0{
			fmt.Printf("Result : %f",num1/num2)
		}else{
			fmt.Printf("Division is not possible")
		}
	default:
		fmt.Printf("Invalid operator")
	}
}