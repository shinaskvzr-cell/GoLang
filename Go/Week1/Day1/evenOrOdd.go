package main
import "fmt"
func main(){
	var n int
	fmt.Println("Enter a number")
	fmt.Scan(&n)

	if n%2 == 0{
		fmt.Printf("%d is Even",n)
	}else{
		fmt.Printf("%d is Odd",n)
	}
}