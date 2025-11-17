package main
import (
	"errors"
	"fmt"
)
func divide(a,b int)(int,error){
	if b==0{
		return 0, errors.New("Cannot divide by zero")
	}
	return a/b,nil
}
func main(){
	result,err:=divide(10,8)
	if err != nil{
		fmt.Println("Errror", err)
	}else{
		fmt.Println("Result",result)
	}
}