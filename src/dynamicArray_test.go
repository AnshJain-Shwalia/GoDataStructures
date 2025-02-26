package godatastructures

import (
	"testing"
)

func TestDynamicArray_BasicOperations(t *testing.T) {
	arr := &DynamicArray[int]{}

	// Test IsEmpty on new array
	if !arr.IsEmpty() {
		t.Error("New array should be empty")
	}

	// Test Size on new array
	if arr.Size() != 0 {
		t.Errorf("Expected size 0, got %d", arr.Size())
	}

	// Test Append and Size
	arr.Append(1)
	if arr.Size() != 1 {
		t.Errorf("Expected size 1, got %d", arr.Size())
	}

	// Test Get
	val, err := arr.Get(0)
	if err != nil || val != 1 {
		t.Errorf("Expected 1, got %d with error: %v", val, err)
	}

	// Test Set
	err = arr.Set(0, 2)
	if err != nil {
		t.Errorf("Unexpected error on Set: %v", err)
	}
	val, _ = arr.Get(0)
	if val != 2 {
		t.Errorf("Expected 2 after Set, got %d", val)
	}
}

func TestDynamicArray_ErrorCases(t *testing.T) {
	arr := &DynamicArray[int]{}

	// Test Get on empty array
	_, err := arr.Get(0)
	if err == nil {
		t.Error("Expected error on Get from empty array")
	}

	// Test Set on empty array
	err = arr.Set(0, 1)
	if err == nil {
		t.Error("Expected error on Set to empty array")
	}

	// Test Get out of bounds
	arr.Append(1)
	_, err = arr.Get(1)
	if err == nil {
		t.Error("Expected error on Get out of bounds")
	}

	// Test Set out of bounds
	err = arr.Set(1, 2)
	if err == nil {
		t.Error("Expected error on Set out of bounds")
	}

	// Test Pop on empty array
	arr.Clear()
	_, err = arr.Pop()
	if err == nil {
		t.Error("Expected error on Pop from empty array")
	}
}

func TestDynamicArray_MultipleOperations(t *testing.T) {
	arr := &DynamicArray[int]{}

	// Test multiple appends
	for i := 0; i < 5; i++ {
		arr.Append(i)
	}
	if arr.Size() != 5 {
		t.Errorf("Expected size 5, got %d", arr.Size())
	}

	// Test multiple pops
	for i := 4; i >= 0; i-- {
		val, err := arr.Pop()
		if err != nil || val != i {
			t.Errorf("Expected %d, got %d with error: %v", i, val, err)
		}
	}

	if !arr.IsEmpty() {
		t.Error("Array should be empty after all pops")
	}
}

func TestDynamicArray_Clear(t *testing.T) {
	arr := &DynamicArray[int]{}

	// Add some elements
	for i := 0; i < 5; i++ {
		arr.Append(i)
	}

	// Test Clear
	arr.Clear()
	if !arr.IsEmpty() {
		t.Error("Array should be empty after Clear")
	}
	if arr.Size() != 0 {
		t.Errorf("Expected size 0 after Clear, got %d", arr.Size())
	}
}

func TestDynamicArray_String(t *testing.T) {
	arr := &DynamicArray[int]{}

	// Test empty array string
	if arr.String() != "[]" {
		t.Errorf("Expected '[]', got '%s'", arr.String())
	}

	// Test non-empty array string
	arr.Append(1)
	arr.Append(2)
	expected := "[1 2]"
	if arr.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, arr.String())
	}
}

func TestDynamicArray_Downsize(t *testing.T) {
	arr := &DynamicArray[int]{}

	// Add many elements to force capacity increase
	for i := 0; i < 100; i++ {
		arr.Append(i)
	}

	initialCap := cap(arr.data)

	// Remove most elements to trigger downsize
	for i := 0; i < 92; i++ {
		arr.Pop()
	}

	// Capacity should have decreased
	if cap(arr.data) >= initialCap {
		t.Error("Array capacity should have decreased after removing many elements")
	}
}

func TestDynamicArray_Generic(t *testing.T) {
	// Test with string type
	strArr := &DynamicArray[string]{}
	strArr.Append("hello")
	strArr.Append("world")

	val, err := strArr.Get(0)
	if err != nil || val != "hello" {
		t.Errorf("Expected 'hello', got '%s' with error: %v", val, err)
	}

	// Test with float type
	floatArr := &DynamicArray[float64]{}
	floatArr.Append(1.5)
	floatArr.Append(2.5)

	fval, err := floatArr.Get(1)
	if err != nil || fval != 2.5 {
		t.Errorf("Expected 2.5, got %f with error: %v", fval, err)
	}
}

func TestNewDynamicArray(t *testing.T) {
	// Test creating a new dynamic array with initial capacity
	arr := NewDynamicArray[int](5)
	
	// Test that the array is initially empty
	if !arr.IsEmpty() {
		t.Error("New array should be empty")
	}

	// Test that we can append up to and beyond the initial capacity
	for i := 0; i < 7; i++ {
		arr.Append(i)
	}

	// Verify the contents
	for i := 0; i < 7; i++ {
		val, err := arr.Get(i)
		if err != nil || val != i {
			t.Errorf("Expected %d at index %d, got %d with error: %v", i, i, val, err)
		}
	}

	// Test with string type to verify generic behavior
	strArr := NewDynamicArray[string](3)
	strArr.Append("hello")
	strArr.Append("world")
	
	val, err := strArr.Get(0)
	if err != nil || val != "hello" {
		t.Errorf("Expected 'hello', got '%s' with error: %v", val, err)
	}
}
