package main
import "fmt"

type Node struct{
	data int
	next *Node
}

type LinkedList struct{
	head *Node
}

func (l *LinkedList)Delete(n int){

	if l.head ==nil || n<=0{
		return
	}
	if n==1{
		l.head=l.head.next
	}
	current:=l.head
	for i:=1;i<n-1&&current!=nil;i++{
		current=current.next
	}
	if current==nil || current.next==nil{
		return
	}
	current.next=current.next.next
}

func (l *LinkedList) Display(){
	current:=l.head
	for current !=nil{
		fmt.Print(current.data,"->")
		current=current.next
	}
	fmt.Print("End")
}

func main(){
	list := LinkedList{}

	 n1 := &Node{data: 1}
    n2 := &Node{data: 2}
    n3 := &Node{data: 3}

    n1.next = n2
    n2.next = n3

    list.head = n1

	list.Display()
	list.Delete(5)
	list.Display()

	

}