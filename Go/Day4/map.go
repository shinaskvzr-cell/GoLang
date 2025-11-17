package main
import "fmt"
func main(){
	fmt.Println("There are three ways of declaring maps")
	fmt.Println(" 1) Using a map literal")
	person:=map[string]string{
		"name":"shinas",
		"city":"Kochi",
	}
	fmt.Println(person)

	fmt.Println("2) Using make()")
	scores:=make(map[string]int)
	scores["math"]=95
	scores["science"]came
}