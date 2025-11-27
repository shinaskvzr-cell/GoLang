package main
import "fmt"

type Node struct{
	data int
	next *Node
}

type Stack struct{
	top *Node
}

func (s *Stack) Push(val int){
	newNode := &Node{data:val}

	if s.top ==nil{
		s.top=newNode
		return
	}
	curr := s.top

	for curr.next != nil{
		curr.next=newNode
	}
	curr.next = newNode
}

func (s *Stack) Display(){
	if s.top == nil{
		fmt.Println("Stack is empty")
		return
	}
	curr := s.top
	for curr !=nil{
		fmt.Println(curr.data,"->")
		curr=curr.next
	}
	fmt.Println("nil")
}


func main(){
	s := Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Display()
}