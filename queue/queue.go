package main

import "fmt"

// Queue FIFO
type Queue struct {
	slice []int
}

// Enqueue 入队列
func (q *Queue) Enqueue(value int) {
	q.slice = append(q.slice, value)
}

// Dequeue 出队列
func (q *Queue) Dequeue() int {
	out := q.slice[0]
	q.slice = q.slice[1:]
	return out
}

func main() {
	fmt.Println("Queue")

	queue := &Queue{}
	data := []int{100, 200, 300}
	for _, v := range data {
		queue.Enqueue(v)
	}
	fmt.Println(queue)
	queue.Dequeue()
	fmt.Println(queue)
}
