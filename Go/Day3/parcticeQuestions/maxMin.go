package main
import "fmt"
func main(){
	arr:=[5]int{24,22,13,75,9}
	max:=arr[0]
	min:=arr[0]
	for i:=0;i<len(arr);i++{
		if arr[i]>max{
			max=arr[i]
		}
		if arr[i]<min{
			min=arr[i]
		}
	}
	fmt.Println(max)
	fmt.Println(min)
}