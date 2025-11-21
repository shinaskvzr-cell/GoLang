package main
import "fmt"
func main(){
	arr:=[5]int{10,20,30,40,50}
	sum:=0
	for i:=0;i<len(arr);i++{
		sum+=arr[i]
	}
	fmt.Println("Average = ",sum/(len(arr)))
}