package main
import "fmt"

type Node struct{
	data int
	next *Node
}

type Queue struct{
	front *Node
	rear *Node
}

func (q *Queue) Enqueue(val int){
	newNode := &Node{data:val}

	if q.front == nil{
		q.front = newNode
		q.rear = newNode
		return
	}

	q.rear.next = newNode
	q.rear = newNode
}

func (q *Queue) Dequeue()int{
	if q.front == nil{
		fmt.Println("Queue is empty")
		return -1
	}

	val := q.front.data
	q.front= q.front.next
	
	if q.front == nil{
		q.rear = nil
	}
	return val
}

func (q *Queue) Display(){
	if q.front == nil{
		fmt.Println("Queue is empty")
		return
	}
	curr := q.front
	for curr != nil{
		fmt.Println(curr.data," ")
		curr = curr.next
	}
	fmt.Println()
}


func main(){
	q := Queue{}
	q.Enqueue(10)
	q.Enqueue(990)
	q.Enqueue(120)
	q.Display()
}