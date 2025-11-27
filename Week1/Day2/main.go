package main
import (
	"fmt"
	"day2/calculator"
	"day2/errorHandling"
)

func main(){
	fmt.Println(calculator.Add(10,20))
	fmt.Println(calculator.Sub(10,20))
	fmt.Println(calculator.Mul(10,20))
	fmt.Println(calculator.Div(10,20))
	errorHandling.ErrorHandling()
	fmt.Println(arithmetic.Add(100,150))
}