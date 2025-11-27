package fibonacci
import "fmt"

func Fibonacci(){
	var limit int
	fmt.Println("Enter a limit")
	fmt.Scan(&limit)

	a := 0
	b := 1
	for i := 0; i<limit ;i++{
		fmt.Println(a)
		c := a+b
		a = b
		b = c
	}
}