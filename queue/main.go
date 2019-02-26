package main

import (
	"log"
)

// Queue is a data structure to hold queues
type Queue struct {
	start int
	end   int
	data  []int
}

// NewQueue will create a new Queue data structure
func NewQueue(size int) *Queue {
	q := new(Queue)
	q.data = make([]int, size)
	q.start = 0
	q.end = 0
	return q
}

// Enqueue will add a new item to the queue and return if it was succesful.
func (q *Queue) Enqueue(x int) bool {
	if q.start >= len(q.data) {
		return false
	}

	q.data[q.end] = x
	q.end++
	log.Println("adding", x, "to the queue: ", q)

	return true
}

// Dequeue extracts an element from the queue
func (q *Queue) Dequeue() int {
	if q.start == q.end {
		return -1
	}

	x := q.data[q.start]
	q.data[q.start] = 0
	q.start++
	log.Println("removing", x, "from the queue: ", q)
	return x
}

func main() {

	a := make([]int, 10)
	log.Printf("a=%v p=%p", a, &a)
	a = append(a, 10, 5, 6, 2, 1, 2)
	log.Printf("a=%v p=%p", a, &a)

	q := NewQueue(10)
	q.Enqueue(5)
	q.Enqueue(6)
	q.Enqueue(7)

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()

	q.Enqueue(9)
	q.Enqueue(6)
	q.Enqueue(4)

}
