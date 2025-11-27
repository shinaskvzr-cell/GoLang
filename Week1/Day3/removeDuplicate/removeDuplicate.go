package removeDuplicate
import "fmt"

// func RemoveDuplicate(){
// 	arr := []int{1,2,3,4,4,3,4,5,2,4,5}
// 	seen := make(map[int]bool)
// 	unique := []int{}
// 	for _,v := range arr {
// 		if !seen[v]{
// 			unique = append(unique,v)
// 			seen[v]=true
// 		}
// 	}
// 	fmt.Println(unique)
// }

func RemoveDuplicate(){
	arr := []int{1,2,3,4,4,3,6,5,1,7,6,6}
	seen := make(map[int]bool)
	unique := []int{}

	for _,v := range arr {
		if !seen[v]{
			unique = append(unique,v)
			seen[v]=true
		}
	}
	fmt.Println("Unique Numbers are:",unique)
}