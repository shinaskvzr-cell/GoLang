package main
import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack struct{
	top *Node
}

func (s *Stack) Push (x int){
	newNode := &Node {data:x}
	newNode.next=s.top
	s.top=newNode
}

func (s *Stack) Pop ()int{
	if s.top == nil{
		fmt.Println("Stack Underflow")
		return -1
	}
	val := s.top.data
	s.top=s.top.next
	return val
}

func (s *Stack) Peek()int{
	if s.top ==nil{
		fmt.Println("Stack is empty")
		return -1
	}
	return s.top.data
}

func main(){
	s:=&Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println(s.Pop())
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
}