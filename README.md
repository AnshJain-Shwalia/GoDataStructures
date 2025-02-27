# GoDataStructures

A Go package providing generic implementations of common data structures.

## Table of Contents
- [Installation](#installation)
- [Data Structures](#data-structures)
  - [Dynamic Array](#dynamic-array)
  - [Stack](#stack)
  - [Queue](#queue)
  - [Min Heap](#min-heap)
- [Usage Examples](#usage-examples)
- [Contributing](#contributing)
- [License](#license)

## Installation

To use this package in your Go project:

```bash
go get github.com/AnshJain-Shwalia/GoDataStructures
```

Make sure you have Go 1.23.0 or later installed, as this package uses Go generics.

## Data Structures

### Dynamic Array

A resizable array implementation that grows and shrinks dynamically.

```go
import "github.com/AnshJain-Shwalia/GoDataStructures/godatastructures"

// Create a new dynamic array with initial capacity
arr := godatastructures.NewDynamicArray[int](10)

// Add elements
arr.Append(5)
arr.Append(10)
arr.Append(15)

// Get element at index
val, err := arr.Get(1) // Returns 10

// Set element at index
err = arr.Set(1, 20) // Sets the element at index 1 to 20

// Get size
size := arr.Size() // Returns 3

// Check if empty
isEmpty := arr.IsEmpty() // Returns false

// Remove the last element
lastVal, err := arr.Pop() // Returns 15

// Clear the array
arr.Clear()
```

### Stack

A Last-In-First-Out (LIFO) data structure.

```go
import "github.com/AnshJain-Shwalia/GoDataStructures/godatastructures"

// Create a new stack
stack := godatastructures.NewStack[string]()

// Push elements
stack.Push("first")
stack.Push("second")
stack.Push("third")

// Peek at the top element without removing it
top, err := stack.Peek() // Returns "third"

// Pop the top element
val, err := stack.Pop() // Returns "third"

// Get size
size := stack.Size() // Returns 2

// Check if empty
isEmpty := stack.IsEmpty() // Returns false
```

### Queue

A First-In-First-Out (FIFO) data structure.

```go
import "github.com/AnshJain-Shwalia/GoDataStructures/godatastructures"

// Create a new queue
queue := godatastructures.NewQueue[float64]()

// Add elements to the queue
queue.Enqueue(1.1)
queue.Enqueue(2.2)
queue.Enqueue(3.3)

// Peek at the front element without removing it
front, err := queue.Peek() // Returns 1.1

// Get the rear element without removing it
rear, err := queue.Rear() // Returns 3.3

// Remove and return the front element
val, err := queue.Dequeue() // Returns 1.1

// Get size
size := queue.Size() // Returns 2

// Check if empty
isEmpty := queue.IsEmpty() // Returns false

// Clear the queue
queue.Clear()
```

### Min Heap

A binary heap data structure that maintains the min-heap property.

```go
import "github.com/AnshJain-Shwalia/GoDataStructures/godatastructures"

// Create a new min heap
heap := godatastructures.NewMinHeap[int]()

// Insert elements
heap.Insert(5)
heap.Insert(3)
heap.Insert(8)
heap.Insert(1)

// Peek at the minimum element without removing it
min, err := heap.Peek() // Returns 1

// Remove and return the minimum element
min, err = heap.Pop() // Returns 1

// Get size
size := heap.Size() // Returns 3

// Check if empty
isEmpty := heap.IsEmpty() // Returns false
```

## Usage Examples

Here's a complete example showing how to use multiple data structures together:

```go
package main

import (
	"fmt"

	"github.com/AnshJain-Shwalia/GoDataStructures/godatastructures"
)

func main() {
	// Create a stack and push some elements
	stack := godatastructures.NewStack[int]()
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	
	// Create a queue to store elements from the stack
	queue := godatastructures.NewQueue[int]()
	
	// Pop elements from stack and enqueue them
	for !stack.IsEmpty() {
		val, _ := stack.Pop()
		queue.Enqueue(val)
		fmt.Printf("Moved %d from stack to queue\n", val)
	}
	
	// Create a dynamic array to store elements from the queue
	arr := godatastructures.NewDynamicArray[int](5)
	
	// Dequeue elements and add them to the array
	for !queue.IsEmpty() {
		val, _ := queue.Dequeue()
		arr.Append(val)
		fmt.Printf("Moved %d from queue to array\n", val)
	}
	
	// Create a min heap and add elements from the array
	heap := godatastructures.NewMinHeap[int]()
	
	for i := 0; i < arr.Size(); i++ {
		val, _ := arr.Get(i)
		heap.Insert(val)
		fmt.Printf("Added %d to min heap\n", val)
	}
	
	// Extract all elements from the heap in sorted order
	fmt.Println("Elements in sorted order:")
	for !heap.IsEmpty() {
		val, _ := heap.Pop()
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the terms of the license included in the repository.
