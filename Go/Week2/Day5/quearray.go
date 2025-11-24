package main
import "fmt"

type Queue struct{
	arr [5]int
	front int
	rear int
	itemCount int
}

func (q *Queue) Enqueue (val int){
	if q.itemCount == len(q.arr){
		fmt.Println("Queue is full")
		return
	}

	q.arr[q.rear]= val
	q.rear = (q.rear)+1%len(q.arr)
	q.itemCount++
}

func (q *Queue) Dequeue()int{
	if q.itemCount == 0{
		fmt.Println("Queue is empty")
		return -1
	}
	val := q.arr[q.front]
	q.front = (q.front +1)%len(q.arr)
	return val
}

func (q *Queue) Peek()int{
	if q.itemCount == 0{
		fmt.Println("Queue is empty")
		return -1 
	}
	return q.arr[q.front]
}

func (q *Queue) Display(){
	if q.itemCount == 0{
		fmt.Println("Queue is empty")
		return
	}
	i := q.front
	count := q.itemCount

	for count > 0{
		fmt.Println(q.arr[i]," ")
		i=(i+1)%len(q.arr)
		count --
	}
	fmt.Println()
}


func main(){

	q:= Queue{}

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.Enqueue(0)
	q.Enqueue(30)
	q.Display()
	fmt.Println(q.Peek())
	q.Dequeue()
	q.Display()
}