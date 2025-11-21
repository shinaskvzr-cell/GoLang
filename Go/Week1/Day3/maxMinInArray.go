package main
import "fmt"
func main(){
	arr:=[5]int{150,200,130,240,50}
	max := arr[0]
	for i:=0;i<len(arr);i++{
		if arr[i]>max{
			max=arr[i]
		}
	}
	fmt.Println(max)
}