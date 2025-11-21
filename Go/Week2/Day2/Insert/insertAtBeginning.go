package main
import "fmt"

func main(){
	sli := []int{5,10,15,20}
	value := 2
	sli = append([]int{value},sli...)
	fmt.Print(sli)
}