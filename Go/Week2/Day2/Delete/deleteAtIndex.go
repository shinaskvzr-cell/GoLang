package main
import "fmt"

func main(){

	arr := [5]int{10,20,30,40,50}
	index :=2
	for i:=index ;i<len(arr)-1;i++{
		arr[i]=arr[i+1]
	}
	arr[len(arr)-1]=0
	fmt.Println(arr)

}