package godatastructures

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Queue[T any] struct {
	head *Node[T]
	size int
	tail *Node[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(item T) {
	newNode := Node[T]{value: item}
	if q.size == 0 {
		q.head = &newNode
		q.tail = &newNode
		q.size += 1
		return
	}
	q.tail.next = &newNode
	q.tail = &newNode
	q.size += 1
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.size <= 0 {
		var zero T
		return zero, fmt.Errorf("empty queue, can't dequeue")
	}
	var result T
	result, q.head = q.head.value, q.head.next
	q.size--

	if q.head == nil {
		q.tail = nil
	}
	return result, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if q.size <= 0 {
		var zero T
		return zero, fmt.Errorf("empty queue, can't peek")
	}
	return q.head.value, nil
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size <= 0
}

func (q *Queue[T]) Clear() {
	q.head = nil
	q.tail = nil
	q.size = 0
}

func (q *Queue[T]) Rear() (T, error) {
	if q.size <= 0 {
		var zero T
		return zero, fmt.Errorf("empty queue, can't rear")
	}
	return q.tail.value, nil
}
