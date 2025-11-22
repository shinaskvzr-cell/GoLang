package main
import "fmt"

type Node struct{
	data int
	next *Node
	prev *Node
}

type LinkedList struct{
	head *Node
	tail *Node
}

func (l *LinkedList) InsertAtEnd (value int){
	newNode := &Node{data:value}

	if l.head == nil {
		l.head = newNode
		l.tail = l.head
		return
	}
	curr := l.head
	for curr.next != nil{
		curr=curr.next
	}
	curr.next=newNode
	newNode.prev=curr
	l.tail=newNode
}

func (l *LinkedList)DisplayForward(){
	curr := l.head
	for curr != nil{
		fmt.Print(curr.data,"->")
		curr = curr.next
	}
	fmt.Println("End")
}
func (l *LinkedList)DisplayBackward(){
	curr := l.tail
	for curr != nil{
		fmt.Print(curr.data,"->")
		curr = curr.prev
	}
	fmt.Println("End")
}


func main(){
	list := LinkedList{}
	list.InsertAtEnd(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)
	list.DisplayForward()
	list.DisplayBackward()

}