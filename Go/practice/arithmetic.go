package main
import "fmt"
func main() {
	add:=func (a,b int)int{
		return a+b
	}(10,8)
	fmt.Println("Sum=",add)
}
