package errorHandling
import (
	"fmt"
	"errors"
)

func Division(a,b float64)(float64,error){
	if b==0{
		return 0, errors.New("Cannot divide by zero")
	}else{
		return a/b,nil
	}

}


func ErrorHandling(){
	var a float64
	fmt.Println("Enter a number")
	fmt.Scan(&a)
	
	var b float64
	fmt.Println("Enter another number")
	fmt.Scan(&b)

	result , err := Division(a,b)

	if err != nil{
		fmt.Println("Error",err)
		return
	}

	fmt.Println("Result:",result)
}