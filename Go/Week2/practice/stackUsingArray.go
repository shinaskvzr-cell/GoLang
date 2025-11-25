package main
import "fmt"

type Stack struct{
	arr [5]int
	top int
}


func (s *Stack) Push (val int){
	if s.top == len(s.arr)-1{
		fmt.Println("Stack is full")
		return
	}

	s.top++
	s.arr[s.top]=val

}


func (s *Stack) Pop ()int{
	if s.top == -1{
		fmt.Println("Stack is empty")
		return -1
	}

	val := s.arr[s.top]
	s.top--
	return val
}

func (s *Stack) Display(){
	if s.top ==-1{
		fmt.Println("Stack is empty")
		return
	}

	for i:=s.top;i>=0;i--{
		fmt.Println(s.arr[i])
	}
}

func main(){
	s := Stack{top :-1}
	s.Push(10)
	s.Push(9980)
	s.Push(20)
	s.Display()
}