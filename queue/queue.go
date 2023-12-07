// Package queue implements queue with linkedlist
package queue

import "fmt"

type queueNode struct {
	val  int
	next *queueNode
}

type queue struct {
	head *queueNode
	len  int
}

func (q *queue) enqueue(val int) {
	newNode := &queueNode{
		val:  val,
		next: nil,
	}
	if q.head == nil {
		q.head = newNode
		q.len++
	} else {
		newNode.next = q.head
		q.head = newNode
		q.len++
	}
}

func (q *queue) dequeue() int {
	var prev *queueNode
	current := q.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	prev.next = nil
	q.len--
	return current.val
}

func main() {
	q := queue{
		head: nil,
		len:  0,
	}
	q.enqueue(4)
	q.enqueue(3)
	q.enqueue(2)
	q.enqueue(1)
	fmt.Printf("Queue lenght is %d, queue head is %v\n", q.len, q.head)
	fmt.Printf("Dequeue result is %d\n", q.dequeue())
	fmt.Printf("Dequeue result is %d\n", q.dequeue())
}
