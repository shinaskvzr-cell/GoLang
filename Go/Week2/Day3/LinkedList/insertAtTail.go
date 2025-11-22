package main
import "fmt"

type Node struct{
	data int
	next *Node
}

type LinkedList struct{
	head *Node
}

func (l *LinkedList)InsertAtTail(value int){
	newNode := &Node{data:value}
	if l.head ==nil{
		l.head =  newNode
		return
	}
	current := l.head
	for current.next != nil{
		current = current.next
	}
	current.next=newNode
}

func (l *LinkedList)Display(){
	current := l.head
	for current != nil{
		fmt.Print(current.data,"->")
		current=current.next
	}
	fmt.Println("Nil")
}

func main(){
	list := LinkedList{}
	list.InsertAtTail(10)
	list.InsertAtTail(20)
	list.InsertAtTail(1)
	list.InsertAtTail(9)
	list.InsertAtTail(10)
	list.InsertAtTail(8)
	fmt.Println("Linked List Insertion At The End")

	list.Display()
}