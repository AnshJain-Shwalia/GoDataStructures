package godatastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Run("NewStack", func(t *testing.T) {
		stack := Stack[int]{}
		assert.NotNil(t, stack)
		assert.Equal(t, 0, stack.Size())
	})

	t.Run("Push and Peek", func(t *testing.T) {
		stack := Stack[int]{}
		stack.Push(10)
		assert.Equal(t, 1, stack.Size())
		value, err := stack.Peek()
		assert.Nil(t, err)
		assert.Equal(t, 10, value)
	})

	t.Run("Push and Pop", func(t *testing.T) {
		stack := Stack[int]{}
		stack.Push(20)
		stack.Push(30)
		value, err := stack.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 30, value)
		assert.Equal(t, 1, stack.Size())
	})

	t.Run("Pop empty stack", func(t *testing.T) {
		stack := Stack[string]{}
		_, err := stack.Pop()
		assert.NotNil(t, err)
		assert.Equal(t, 0, stack.Size())
	})

	t.Run("Peek empty stack", func(t *testing.T) {
		stack := Stack[float64]{}
		_, err := stack.Peek()
		assert.NotNil(t, err)
		assert.Equal(t, 0, stack.Size())
	})

	t.Run("Multiple operations", func(t *testing.T) {
		stack := Stack[int]{}
		stack.Push(100)
		stack.Push(200)
		stack.Push(300)
		assert.Equal(t, 3, stack.Size())

		value, err := stack.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 300, value)

		value, err = stack.Peek()
		assert.Nil(t, err)
		assert.Equal(t, 200, value)

		stack.Pop()
		stack.Pop()
		assert.Equal(t, 0, stack.Size())
	})
}

func TestStackDownsizing(t *testing.T) {
	stack := Stack[int]{}

	// Push many elements
	numElements := 1000
	for i := 0; i < numElements; i++ {
		stack.Push(i)
	}

	initialCap := cap(stack.data)

	// Pop most elements to trigger downsizing
	for i := 0; i < numElements-10; i++ {
		stack.Pop()
	}

	finalCap := cap(stack.data)
	assert.Less(t, finalCap, initialCap, "Stack should have downsized")
	assert.Equal(t, 10, stack.Size())
}

func TestComplexTypes(t *testing.T) {
	t.Run("Struct type", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		stack := Stack[Person]{}
		p1 := Person{"Alice", 25}
		p2 := Person{"Bob", 30}

		stack.Push(p1)
		stack.Push(p2)

		popped, err := stack.Pop()
		assert.Nil(t, err)
		assert.Equal(t, p2, popped)
	})

	t.Run("Map type", func(t *testing.T) {
		stack := Stack[map[string]int]{}
		m1 := map[string]int{"a": 1}
		m2 := map[string]int{"b": 2}

		stack.Push(m1)
		stack.Push(m2)

		popped, err := stack.Pop()
		assert.Nil(t, err)
		assert.Equal(t, m2, popped)
	})

	t.Run("Interface type", func(t *testing.T) {
		stack := Stack[any]{}
		stack.Push(42)
		stack.Push("hello")
		stack.Push(3.14)

		val, err := stack.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 3.14, val)
	})
}

func TestEdgeCases(t *testing.T) {
	t.Run("Push after empty", func(t *testing.T) {
		stack := Stack[int]{}
		stack.Push(1)
		stack.Pop()
		stack.Push(2)
		val, err := stack.Peek()
		assert.Nil(t, err)
		assert.Equal(t, 2, val)
	})

	t.Run("Multiple empty checks", func(t *testing.T) {
		stack := Stack[int]{}
		assert.True(t, stack.IsEmpty())
		stack.Push(1)
		assert.False(t, stack.IsEmpty())
		stack.Pop()
		assert.True(t, stack.IsEmpty())
	})

	t.Run("Zero values", func(t *testing.T) {
		stack := Stack[int]{}
		stack.Push(0)
		val, err := stack.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 0, val)
	})
}

func BenchmarkStack(b *testing.B) {
	b.Run("Push", func(b *testing.B) {
		stack := Stack[int]{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			stack.Push(i)
		}
	})

	b.Run("Pop", func(b *testing.B) {
		stack := Stack[int]{}
		for i := 0; i < b.N; i++ {
			stack.Push(i)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			stack.Pop()
		}
	})

	b.Run("Mixed Operations", func(b *testing.B) {
		stack := Stack[int]{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if i%2 == 0 {
				stack.Push(i)
			} else {
				stack.Pop()
			}
		}
	})
}
