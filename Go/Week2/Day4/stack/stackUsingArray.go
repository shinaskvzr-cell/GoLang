package main
import "fmt"

type Stack struct {
	data []int
	top int
}

func NewStack(size int)*Stack{
	return &Stack{
		data : make([]int , size),
		top : -1,
	}
}

func (s *Stack) Push (x int){
	if s.top ==len(s.data)-1{
		fmt.Println("Stack Overflow")
		return 
	}
	s.top++
	s.data[s.top]=x
	
}

func (s *Stack) Pop ()int{
	if s.top == -1{
		fmt.Println("Stack underflow")
		return -1
	}
	val := s.data[s.top]
	s.data[s.top] = 0
	s.top--
	return val
}

func (s *Stack) Peek ()int{
	if s.top == -1{
		fmt.Println("Stack is empty")
		return -1
	}
	return s.data[s.top]
}

func main(){
	s:=NewStack(5)
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println(s)
	fmt.Println(s.data)
	fmt.Println(s.Pop())
	fmt.Println(s.Peek())
	fmt.Println(s)
}