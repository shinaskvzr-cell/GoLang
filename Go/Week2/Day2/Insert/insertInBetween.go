package main 
import "fmt"
// Insertion at in between elements


func main(){
	arr := []int{3,6,9,12,15}
	//need to specify in which index should we insert
	index := 2
	//value to be inserted
	value := 99

	arr = append(arr,0)
	for i:=len(arr)-1;i>index;i--{
		arr[i]=arr[i-1]
	}

	arr[index]=value
	fmt.Println(arr)

}