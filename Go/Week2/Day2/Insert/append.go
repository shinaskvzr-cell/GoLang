//Insert an element at the end WITHOUT using append()

package main
import "fmt"

func main(){
	slice := []int{10,20,30}
	newSlice := make([]int,len(slice)+1)
	for i:=0;i<len(slice);i++{
		newSlice[i]=slice[i]
	}
	newSlice[len(newSlice)-1]=40
	fmt.Println(newSlice)
}