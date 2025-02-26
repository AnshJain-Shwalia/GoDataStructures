package godatastructures

import "fmt"

type DynamicArray[T any] struct {
	data []T
}

func NewDynamicArray[T any](size int) *DynamicArray[T] {
	return &DynamicArray[T]{data: make([]T, 0, size)}
}

func (da *DynamicArray[T]) Append(item T) {
	da.data = append(da.data, item)
}

func (da *DynamicArray[T]) Get(index int) (T, error) {
	if index >= len(da.data) || index < 0 {
		var zero T
		return zero, fmt.Errorf("index out of bounds")
	}
	return da.data[index], nil
}

func (da *DynamicArray[T]) Set(index int, item T) error {
	if index >= len(da.data) || index < 0 {
		return fmt.Errorf("index out of bounds")
	}
	da.data[index] = item
	return nil
}

func (da *DynamicArray[T]) Swap(index1, index2 int) error {
	if index1 >= len(da.data) || index1 < 0 || index2 >= len(da.data) || index2 < 0 {
		return fmt.Errorf("index out of bounds")
	}
	da.data[index1], da.data[index2] = da.data[index2], da.data[index1]
	return nil
}

func (da *DynamicArray[T]) Size() int {
	return len(da.data)
}

func (da *DynamicArray[T]) IsEmpty() bool {
	return len(da.data) == 0
}

func (da *DynamicArray[T]) Clear() {
	da.data = make([]T, 0)
}

func (da *DynamicArray[T]) String() string {
	return fmt.Sprint(da.data)
}

func (da *DynamicArray[T]) Pop() (T, error) {
	if da.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("empty array")
	}
	item := da.data[len(da.data)-1]
	da.data = da.data[:len(da.data)-1]
	da.downsize()
	return item, nil
}

func (da *DynamicArray[T]) downsize() {
	length, capacity := len(da.data), cap(da.data)
	if length <= capacity/4 {
		newSlice := make([]T, length, capacity/2)
		copy(newSlice, da.data)
		da.data = newSlice
	}
}
