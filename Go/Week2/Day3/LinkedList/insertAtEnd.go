package main
import "fmt"

type Node struct{
	data int
	next *Node
}

type LinkedList struct{
	head *Node
}

func (l *LinkedList) InsertAtEnd (value int){
	newNode := &Node{data:value}

	if l.head == nil {
		l.head = newNode
		return
	}
	curr := l.head
	for curr.next != nil{
		curr=curr.next
	}
	curr.next=newNode
}

func (l *LinkedList)Display(){
	curr := l.head
	for curr != nil{
		fmt.Print(curr.data,"->")
		curr = curr.next
	}
	fmt.Println("End")
}


func main(){
	list := LinkedList{}
	list.InsertAtEnd(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)
	list.Display()

}