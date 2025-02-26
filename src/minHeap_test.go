package godatastructures

import (
	"testing"
)

func TestMinHeap_NewMinHeap(t *testing.T) {
	heap := NewMinHeap[int]()
	
	if heap == nil {
		t.Error("NewMinHeap should return a non-nil heap")
	}
	
	if heap.data == nil {
		t.Error("NewMinHeap should initialize the data field")
	}
	
	if !heap.IsEmpty() {
		t.Error("New heap should be empty")
	}
	
	if heap.Size() != 0 {
		t.Errorf("Expected size 0, got %d", heap.Size())
	}
}

func TestMinHeap_BasicOperations(t *testing.T) {
	heap := NewMinHeap[int]()
	
	// Test IsEmpty on new heap
	if !heap.IsEmpty() {
		t.Error("New heap should be empty")
	}
	
	// Test Size on new heap
	if heap.Size() != 0 {
		t.Errorf("Expected size 0, got %d", heap.Size())
	}
	
	// Test Peek on empty heap
	_, err := heap.Peek()
	if err == nil {
		t.Error("Expected error on Peek from empty heap")
	}
	
	// Test Insert and Size
	heap.Insert(5)
	if heap.Size() != 1 {
		t.Errorf("Expected size 1, got %d", heap.Size())
	}
	
	// Test Peek after insert
	val, err := heap.Peek()
	if err != nil || val != 5 {
		t.Errorf("Expected 5, got %d with error: %v", val, err)
	}
	
	// Test Pop
	val, err = heap.Pop()
	if err != nil || val != 5 {
		t.Errorf("Expected 5, got %d with error: %v", val, err)
	}
	
	// Test IsEmpty after Pop
	if !heap.IsEmpty() {
		t.Error("Heap should be empty after Pop")
	}
}

func TestMinHeap_MultipleInserts(t *testing.T) {
	heap := NewMinHeap[int]()
	
	// Insert elements in non-sorted order
	elements := []int{5, 3, 8, 1, 2, 7, 6, 4}
	for _, e := range elements {
		heap.Insert(e)
	}
	
	// Verify size
	if heap.Size() != len(elements) {
		t.Errorf("Expected size %d, got %d", len(elements), heap.Size())
	}
	
	// Verify min element
	min, err := heap.Peek()
	if err != nil || min != 1 {
		t.Errorf("Expected min element 1, got %d with error: %v", min, err)
	}
	
	// Pop all elements and verify they come out in ascending order
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, e := range expected {
		val, err := heap.Pop()
		if err != nil || val != e {
			t.Errorf("Expected %d, got %d with error: %v", e, val, err)
		}
	}
	
	// Verify heap is empty
	if !heap.IsEmpty() {
		t.Error("Heap should be empty after all Pops")
	}
}

func TestMinHeap_ErrorCases(t *testing.T) {
	heap := NewMinHeap[int]()
	
	// Test Peek on empty heap
	_, err := heap.Peek()
	if err == nil {
		t.Error("Expected error on Peek from empty heap")
	}
	
	// Test Pop on empty heap
	_, err = heap.Pop()
	if err == nil {
		t.Error("Expected error on Pop from empty heap")
	}
}

func TestMinHeap_HeapifyUpDown(t *testing.T) {
	heap := NewMinHeap[int]()
	
	// Insert in reverse order to test heapifyUp
	for i := 10; i > 0; i-- {
		heap.Insert(i)
		// After each insert, the minimum should be the smallest value inserted so far
		min, err := heap.Peek()
		if err != nil || min != i {
			t.Errorf("After inserting %d, expected min to be %d, got %d with error: %v", i, i, min, err)
		}
	}
	
	// Pop half the elements to test heapifyDown
	for i := 1; i <= 5; i++ {
		val, err := heap.Pop()
		if err != nil || val != i {
			t.Errorf("Expected %d, got %d with error: %v", i, val, err)
		}
	}
	
	// Verify the remaining elements
	for i := 6; i <= 10; i++ {
		val, err := heap.Pop()
		if err != nil || val != i {
			t.Errorf("Expected %d, got %d with error: %v", i, val, err)
		}
	}
}

func TestMinHeap_DuplicateValues(t *testing.T) {
	heap := NewMinHeap[int]()
	
	// Insert duplicate values
	values := []int{3, 1, 2, 1, 3, 2, 1}
	for _, v := range values {
		heap.Insert(v)
	}
	
	// Expected output after sorting
	expected := []int{1, 1, 1, 2, 2, 3, 3}
	for _, e := range expected {
		val, err := heap.Pop()
		if err != nil || val != e {
			t.Errorf("Expected %d, got %d with error: %v", e, val, err)
		}
	}
}

func TestMinHeap_GenericTypes(t *testing.T) {
	// Test with string type
	heapStr := NewMinHeap[string]()
	
	strings := []string{"banana", "apple", "cherry", "date"}
	for _, s := range strings {
		heapStr.Insert(s)
	}
	
	// Strings should come out in lexicographical order
	expected := []string{"apple", "banana", "cherry", "date"}
	for _, e := range expected {
		val, err := heapStr.Pop()
		if err != nil || val != e {
			t.Errorf("Expected %s, got %s with error: %v", e, val, err)
		}
	}
	
	// Test with float type
	heapFloat := NewMinHeap[float64]()
	
	floats := []float64{3.14, 1.41, 2.71, 0.99}
	for _, f := range floats {
		heapFloat.Insert(f)
	}
	
	// Floats should come out in ascending order
	expectedFloats := []float64{0.99, 1.41, 2.71, 3.14}
	for _, e := range expectedFloats {
		val, err := heapFloat.Pop()
		if err != nil || val != e {
			t.Errorf("Expected %f, got %f with error: %v", e, val, err)
		}
	}
}

func TestMinHeap_LargeDataset(t *testing.T) {
	heap := NewMinHeap[int]()
	
	// Insert a large number of elements
	n := 1000
	for i := n; i > 0; i-- {
		heap.Insert(i)
	}
	
	// Verify size
	if heap.Size() != n {
		t.Errorf("Expected size %d, got %d", n, heap.Size())
	}
	
	// Verify elements come out in ascending order
	for i := 1; i <= n; i++ {
		val, err := heap.Pop()
		if err != nil || val != i {
			t.Errorf("Expected %d, got %d with error: %v", i, val, err)
			break // Stop after first failure to avoid too many errors
		}
	}
}

func TestMinHeap_InsertAfterPop(t *testing.T) {
	heap := NewMinHeap[int]()
	
	// Insert some elements
	for i := 5; i > 0; i-- {
		heap.Insert(i)
	}
	
	// Pop some elements
	for i := 1; i <= 3; i++ {
		val, err := heap.Pop()
		if err != nil || val != i {
			t.Errorf("Expected %d, got %d with error: %v", i, val, err)
		}
	}
	
	// Insert more elements
	heap.Insert(1)
	heap.Insert(2)
	
	// Verify the order
	expected := []int{1, 2, 4, 5}
	for _, e := range expected {
		val, err := heap.Pop()
		if err != nil || val != e {
			t.Errorf("Expected %d, got %d with error: %v", e, val, err)
		}
	}
}

func TestMinHeap_EdgeCase_SingleElement(t *testing.T) {
	heap := NewMinHeap[int]()
	
	// Insert and immediately pop a single element
	heap.Insert(42)
	val, err := heap.Pop()
	
	if err != nil || val != 42 {
		t.Errorf("Expected 42, got %d with error: %v", val, err)
	}
	
	// Verify heap is empty
	if !heap.IsEmpty() {
		t.Error("Heap should be empty after popping the only element")
	}
}

func TestMinHeap_EdgeCase_NegativeNumbers(t *testing.T) {
	heap := NewMinHeap[int]()
	
	// Insert negative numbers
	values := []int{-5, -10, -3, -7, -1}
	for _, v := range values {
		heap.Insert(v)
	}
	
	// Expected output after sorting
	expected := []int{-10, -7, -5, -3, -1}
	for _, e := range expected {
		val, err := heap.Pop()
		if err != nil || val != e {
			t.Errorf("Expected %d, got %d with error: %v", e, val, err)
		}
	}
}

func TestMinHeap_MixedOperations(t *testing.T) {
	heap := NewMinHeap[int]()
	
	// Insert some elements
	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(7)
	
	// Check min
	min, err := heap.Peek()
	if err != nil || min != 3 {
		t.Errorf("Expected min 3, got %d with error: %v", min, err)
	}
	
	// Pop min
	val, err := heap.Pop()
	if err != nil || val != 3 {
		t.Errorf("Expected 3, got %d with error: %v", val, err)
	}
	
	// Insert more elements
	heap.Insert(2)
	heap.Insert(4)
	
	// Check new min
	min, err = heap.Peek()
	if err != nil || min != 2 {
		t.Errorf("Expected min 2, got %d with error: %v", min, err)
	}
	
	// Pop all and verify order
	expected := []int{2, 4, 5, 7}
	for _, e := range expected {
		val, err := heap.Pop()
		if err != nil || val != e {
			t.Errorf("Expected %d, got %d with error: %v", e, val, err)
		}
	}
}
