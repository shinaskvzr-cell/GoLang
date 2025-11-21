package main
import "fmt"
func operations (a,b int) {
	fmt.Println("Sum = ",a+b)
	fmt.Println("Sub = ",a-b)
	fmt.Println("Mul = ",a*b)
	fmt.Println("Div = ",a/b)
}

func main(){
	operations(10,20)
}