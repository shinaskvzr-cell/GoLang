package main
import "fmt"

type Node struct{
	data int
	next *Node
}

type Stack struct{
	top *Node
}

func (s *Stack) Push (val int){
	newNode := &Node{data:val}
	newNode.next= s.top
	s.top=newNode
}


func (s *Stack) Pop ()int {
	if s.top == nil{
		fmt.Println("Stack is empty")
		return -1
	}
	val := s.top.data
	s.top = s.top.next
	return val
}

func (s *Stack) Display(){
	if s.top == nil{
		fmt.Println("Stack is empty")
		return
	}
	curr := s.top

	for curr != nil{
		fmt.Print(curr.data,"->")
		curr = curr.next
	}
	fmt.Println("nil")
	
}


func main(){
	

}