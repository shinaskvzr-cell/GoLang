package main
import "fmt"
func main() {
	arr:=[8]int{1,2,3,4,5,6,7,8}
	evenCount:=0
	oddCount:=0
	for i:=0;i<len(arr);i++{
		if arr[i] % 2 == 0 {
			evenCount++
		}else{
			oddCount++
		}
 	}
	fmt.Println("Even Count",evenCount)
	fmt.Println("Odd Count",oddCount)
}