package main
import "fmt"
func main(){
	arr:=[10]int{1,2,3,4,5,6,7,8,9,10}
	var num int
	fmt.Println("Enter the number to search")
	fmt.Scan(&num)

	found:=false
	for i:=0;i<len(arr);i++{
		if arr[i]==num{
			found=true
			break
		}
	}
	if found==true{
		fmt.Println("Element found")
	}else{
		fmt.Println("Element not found")
	}
}