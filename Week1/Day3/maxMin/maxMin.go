package maxMin
import "fmt"

func MaxAndMin(){
	arr := []int{100,23,45,61,23,85,234}
	largest := arr[0]
	smallest := arr[0]
	for i:= 0; i<len(arr);i++{
		if largest<arr[i]{
			largest = arr[i]
		}
		if arr[i]<smallest{
			smallest = arr[i]
		}
	}
	fmt.Println("Largest is :",largest)
	fmt.Println("Smallest is :",smallest)
}