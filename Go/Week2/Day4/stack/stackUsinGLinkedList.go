package main
import "fmt"

type Node struct{
	data int
	next *Node
}
 type Stack struct {
	top *Node
 }
 
func (s *Stack) Push(x int){
	newNode := &Node{data:x}
	newNode.next=s.top
	s.top=newNode
}

func (s *Stack) Pop()int{
	if s.top == nil{
		fmt.Println("Stack is empty")
		return -1
	}
	val := s.top.data
	s.top=s.top.next
	return val
}

func (s *Stack)Traverse(){
	if s.top == nil{
		fmt.Println("Stack is empty")
	}
	current := s.top
	for current != nil{
		fmt.Print(current.data,"->")
		current=current.next
	}
	fmt.Print("End")
}

 func main(){
	list:=Stack{}
	list.Push(10)
	list.Push(20)
	list.Push(30)

	fmt.Println(list.top.data)
	fmt.Println(list.top.next.data)
	fmt.Println(list.top.next.data)
	list.Traverse()


 }