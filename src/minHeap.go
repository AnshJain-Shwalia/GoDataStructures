package godatastructures

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type MinHeap[T constraints.Ordered] struct {
	data *DynamicArray[T]
}

func NewMinHeap[T constraints.Ordered]() *MinHeap[T] {
	return &MinHeap[T]{data: NewDynamicArray[T](0)}
}

func (h *MinHeap[T]) Size() int {
	return h.data.Size()
}

func (h *MinHeap[T]) IsEmpty() bool {
	return h.data.IsEmpty()
}

func (h *MinHeap[T]) Peek() (T, error) {
	return h.data.Get(0)
}

func (h *MinHeap[T]) Insert(item T) {
	h.data.Append(item)
	h.heapifyUp()
}

func (h *MinHeap[T]) Pop() (T, error) {
	if h.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("empty heap.")
	}
	h.data.Swap(0, h.Size()-1)
	value, _ := h.data.Pop()
	h.heapifyDown()
	return value, nil
}

func (h *MinHeap[T]) heapifyUp() {
	currentIndex := h.data.Size() - 1
	currentValue, err1 := h.data.Get(currentIndex)
	parentIndex := parent(currentIndex)
	parentValue, err2 := h.data.Get(parentIndex)
	for err1 == nil && err2 == nil && currentValue < parentValue {
		err := h.data.Swap(currentIndex, parentIndex)
		if err != nil {
			return
		}
		currentIndex = parentIndex
		parentIndex = parent(currentIndex)
		currentValue, err1 = h.data.Get(currentIndex)
		parentValue, err2 = h.data.Get(parentIndex)
	}
}

func (h *MinHeap[T]) heapifyDown() {
	currentIdx := 0

	// Continue until we reach a leaf node or the heap property is satisfied
	for {
		// Get indices of left and right children
		leftIdx := leftIndex(currentIdx)
		rightIdx := rightIndex(currentIdx)

		// Assume current node is the smallest initially
		smallestIdx := currentIdx

		// Get current node value
		currentVal, err := h.data.Get(currentIdx)
		if err != nil {
			return // Invalid index, should not happen
		}

		// Check if left child exists and is smaller than current smallest
		leftVal, leftErr := h.data.Get(leftIdx)
		if leftErr == nil && leftVal < currentVal {
			smallestIdx = leftIdx
		}

		// If smallest is now left child, use its value for comparison with right child
		smallestVal, _ := h.data.Get(smallestIdx)

		// Check if right child exists and is smaller than current smallest
		rightVal, rightErr := h.data.Get(rightIdx)
		if rightErr == nil && rightVal < smallestVal {
			smallestIdx = rightIdx
		}

		// If smallest is still the current node, heap property is satisfied
		if smallestIdx == currentIdx {
			return
		}

		// Swap current node with the smallest child
		if err := h.data.Swap(currentIdx, smallestIdx); err != nil {
			return // Swap failed, should not happen
		}

		// Move down to the child we swapped with
		currentIdx = smallestIdx
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func leftIndex(index int) int {
	return 2*index + 1
}

func rightIndex(index int) int {
	return 2*index + 2
}
