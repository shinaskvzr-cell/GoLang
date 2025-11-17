package main
import "fmt"
func main(){
	var n int
	fmt.Println("Enter a number")
	fmt.Scan(&n)

	a,b:=0,1
	fmt.Println("Fibonacci series upto %d are",n)
	for i:=0; i<n; i++{
		fmt.Printf("%d ",a)
		a, b=b, a+b
	}
	fmt.Println()
}