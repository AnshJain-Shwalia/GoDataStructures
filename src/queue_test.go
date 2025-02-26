package godatastructures

import (
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("New Queue should be empty", func(t *testing.T) {
		q := Queue[int]{}
		if !q.IsEmpty() {
			t.Error("New queue should be empty")
		}
		if q.Size() != 0 {
			t.Errorf("Expected size 0, got %d", q.Size())
		}
	})

	t.Run("Enqueue operations", func(t *testing.T) {
		q := Queue[int]{}
		
		// Test single enqueue
		q.Enqueue(1)
		if q.Size() != 1 {
			t.Errorf("Expected size 1, got %d", q.Size())
		}
		if q.IsEmpty() {
			t.Error("Queue should not be empty after enqueue")
		}

		// Test multiple enqueues
		q.Enqueue(2)
		q.Enqueue(3)
		if q.Size() != 3 {
			t.Errorf("Expected size 3, got %d", q.Size())
		}

		// Test peek after enqueue
		val, err := q.Peek()
		if err != nil {
			t.Errorf("Unexpected error on peek: %v", err)
		}
		if val != 1 {
			t.Errorf("Expected peek value 1, got %v", val)
		}

		// Test rear after enqueue
		rear, err := q.Rear()
		if err != nil {
			t.Errorf("Unexpected error on rear: %v", err)
		}
		if rear != 3 {
			t.Errorf("Expected rear value 3, got %v", rear)
		}
	})

	t.Run("Dequeue operations", func(t *testing.T) {
		q := Queue[int]{}
		
		// Test dequeue on empty queue
		_, err := q.Dequeue()
		if err == nil {
			t.Error("Expected error when dequeuing from empty queue")
		}

		// Setup queue for dequeue tests
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		// Test dequeue
		val, err := q.Dequeue()
		if err != nil {
			t.Errorf("Unexpected error on dequeue: %v", err)
		}
		if val != 1 {
			t.Errorf("Expected dequeued value 1, got %v", val)
		}
		if q.Size() != 2 {
			t.Errorf("Expected size 2 after dequeue, got %d", q.Size())
		}

		// Test multiple dequeues
		val, _ = q.Dequeue()
		if val != 2 {
			t.Errorf("Expected dequeued value 2, got %v", val)
		}
		val, _ = q.Dequeue()
		if val != 3 {
			t.Errorf("Expected dequeued value 3, got %v", val)
		}
		if !q.IsEmpty() {
			t.Error("Queue should be empty after dequeuing all elements")
		}
	})

	t.Run("Peek operations", func(t *testing.T) {
		q := Queue[string]{}

		// Test peek on empty queue
		_, err := q.Peek()
		if err == nil {
			t.Error("Expected error when peeking empty queue")
		}

		// Test peek with elements
		q.Enqueue("first")
		q.Enqueue("second")

		val, err := q.Peek()
		if err != nil {
			t.Errorf("Unexpected error on peek: %v", err)
		}
		if val != "first" {
			t.Errorf("Expected peek value 'first', got %v", val)
		}

		// Verify peek doesn't remove elements
		if q.Size() != 2 {
			t.Errorf("Peek should not modify queue size, expected 2, got %d", q.Size())
		}
	})

	t.Run("Rear operations", func(t *testing.T) {
		q := Queue[string]{}

		// Test rear on empty queue
		_, err := q.Rear()
		if err == nil {
			t.Error("Expected error when getting rear of empty queue")
		}

		// Test rear with elements
		q.Enqueue("first")
		q.Enqueue("second")
		q.Enqueue("third")

		val, err := q.Rear()
		if err != nil {
			t.Errorf("Unexpected error on rear: %v", err)
		}
		if val != "third" {
			t.Errorf("Expected rear value 'third', got %v", val)
		}
	})

	t.Run("Clear operations", func(t *testing.T) {
		q := Queue[int]{}
		
		// Add some elements
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		// Clear the queue
		q.Clear()

		if !q.IsEmpty() {
			t.Error("Queue should be empty after clear")
		}
		if q.Size() != 0 {
			t.Errorf("Queue size should be 0 after clear, got %d", q.Size())
		}

		// Verify operations after clear
		_, err := q.Peek()
		if err == nil {
			t.Error("Expected error when peeking after clear")
		}
		_, err = q.Dequeue()
		if err == nil {
			t.Error("Expected error when dequeuing after clear")
		}
	})

	t.Run("Mixed operations", func(t *testing.T) {
		q := Queue[int]{}

		// Test sequence of operations
		q.Enqueue(1)
		q.Enqueue(2)
		val, _ := q.Dequeue()
		if val != 1 {
			t.Errorf("Expected dequeued value 1, got %v", val)
		}
		
		q.Enqueue(3)
		val, _ = q.Peek()
		if val != 2 {
			t.Errorf("Expected peek value 2, got %v", val)
		}

		rear, _ := q.Rear()
		if rear != 3 {
			t.Errorf("Expected rear value 3, got %v", rear)
		}

		if q.Size() != 2 {
			t.Errorf("Expected size 2, got %d", q.Size())
		}
	})

	t.Run("Generic type support", func(t *testing.T) {
		// Test with string type
		strQueue := Queue[string]{}
		strQueue.Enqueue("hello")
		strQueue.Enqueue("world")
		val, _ := strQueue.Dequeue()
		if val != "hello" {
			t.Errorf("Expected 'hello', got %v", val)
		}

		// Test with float type
		floatQueue := Queue[float64]{}
		floatQueue.Enqueue(1.5)
		floatQueue.Enqueue(2.5)
		val2, _ := floatQueue.Peek()
		if val2 != 1.5 {
			t.Errorf("Expected 1.5, got %v", val2)
		}
	})
}
