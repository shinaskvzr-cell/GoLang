package main
import "fmt"

func main(){
	arr := []int{10,20,30}
	value := 40
	arr = append(arr,value)
	fmt.Println(arr)
}