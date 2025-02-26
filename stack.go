package godatastructures

type Stack[T any] struct {
	data []T
}

type Err string

func (err Err) Error() string {
	return string(err)
}

func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.data) <= 0 {
		var zero T
		return zero, Err("Empty stack")
	}
	result := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	s.downsize()
	return result, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if len(s.data) <= 0 {
		var zero T
		return zero, Err("Empty stack")
	}
	return s.data[len(s.data)-1], nil
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) <= 0
}

func (s *Stack[T]) downsize() {
	length, capacity := len(s.data), cap(s.data)

	if length <= capacity/4 {
		newSlice := make([]T, length, capacity/2)
		copy(newSlice, s.data)
		s.data = newSlice
	}
}
