package main
import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct{
	head *Node
}

func (l *LinkedList) InsertAtBeginning(value int){
	newNode := &Node{data:value}
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList)Display(){
	current := l.head
	for current != nil {
		fmt.Print(current.data,"->")
		current = current.next
	}
	fmt.Println("nil")
}

func main(){

	list := LinkedList{}
	list.InsertAtBeginning(10)
	list.InsertAtBeginning(10)
	list.InsertAtBeginning(10)
	list.InsertAtBeginning(30)
	fmt.Println("Linked List")
	list.Display()

}