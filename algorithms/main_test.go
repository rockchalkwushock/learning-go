package main

import (
	"reflect"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	target := 3
	want := true
	got := LinearSearch(array, target)
	if got != want {
		t.Errorf("LinearSearch(%v, %v) = %v, want %v", array, target, got, want)
	}
}

func TestBinarySearch(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	target := 3
	want := true
	got := BinarySearch(array, target)
	if got != want {
		t.Errorf("BinarySearch(%v, %v) = %v, want %v", array, target, got, want)
	}
}

func TestTwoCrystalBalls(t *testing.T) {
	array := []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, true}
	want := 14
	got := TwoCrystalBalls(array)
	if got != want {
		t.Errorf("TwoCrystalBalls(%v) = %v, want %v", array, got, want)
	}
}

func TestBubbleSort(t *testing.T) {
	array := []int{4, 2, 1, 3, 5}
	want := []int{1, 2, 3, 4, 5}
	got := BubbleSort(array)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("BubbleSort(%v) = %v, want %v", array, got, want)
	}
}

func TestQueue(t *testing.T) {
	q := Queue{}

	if !q.IsEmpty() {
		t.Errorf("Queue.IsEmpty() = %v, want %v", q.IsEmpty(), true)
	}

	q.Enqueue(1)

	if q.IsEmpty() {
		t.Errorf("Queue.IsEmpty() = %v, want %v", q.IsEmpty(), false)
	}

	if q.Peek() != 1 {
		t.Errorf("Queue.Peek() = %v, want %v", q.Peek(), 1)
	}

	if q.Dequeue() != 1 {
		t.Errorf("Queue.Dequeue() = %v, want %v", q.Dequeue(), 1)
	}

	if !q.IsEmpty() {
		t.Errorf("Queue.IsEmpty() = %v, want %v", q.IsEmpty(), true)
	}
}

func TestStack(t *testing.T) {
	s := Stack{}

	if !s.IsEmpty() {
		t.Errorf("Stack.IsEmpty() = %v, want %v", s.IsEmpty(), true)
	}

	s.Push(1)

	if s.IsEmpty() {
		t.Errorf("Stack.IsEmpty() = %v, want %v", s.IsEmpty(), false)
	}

	if s.Pop() != 1 {
		t.Errorf("Stack.Pop() = %v, want %v", s.Pop(), 1)
	}

	if !s.IsEmpty() {
		t.Errorf("Stack.IsEmpty() = %v, want %v", s.IsEmpty(), true)
	}
}

func TestArrayBuffer(t *testing.T) {
	ab := NewArrayBuffer(10)

	if ab.Length() != 10 {
		t.Errorf("Expected length to be 10, got %d", ab.Length())
	}

	ab.Set(0, 'H')
	ab.Set(1, 'i')

	if ab.Get(0) != 'H' || ab.Get(1) != 'i' {
		t.Errorf("Unexpected values in ArrayBuffer")
	}
}

func TestArrayList(t *testing.T) {
	al := NewArrayList()

	if al.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", al.Size())
	}

	al.Add("Hello")
	al.Add("World")

	if al.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", al.Size())
	}

	if al.Get(0) != "Hello" || al.Get(1) != "World" {
		t.Errorf("Unexpected values in ArrayList")
	}
}

// MazeSolver
func TestMazeSolver(t *testing.T) {
	testCases := []struct {
		maze  [][]int
		start Point
		end   Point
		want  bool
	}{
		{
			maze: [][]int{
				{1, 1, 1, 1, 1},
				{1, 0, 0, 0, 1},
				{1, 1, 1, 0, 1},
			},
			start: Point{0, 0},
			end:   Point{2, 4},
			want:  true,
		},
		// Add more test cases as needed.
	}

	for _, tc := range testCases {
		got := MazeSolver(tc.maze, tc.start, tc.end)
		if got != tc.want {
			t.Errorf("MazeSolver(%v, %v, %v) = %v, want %v", tc.maze, tc.start, tc.end, got, tc.want)
		}
	}
}

// QuickSort
func TestQuickSort(t *testing.T) {
	testCases := []struct {
		array []int
		want  []int
	}{
		{
			array: []int{4, 2, 1, 3, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		// Add more test cases as needed.
	}

	for _, tc := range testCases {
		got := QuickSort(tc.array)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("QuickSort(%v) = %v, want %v", tc.array, got, tc.want)
		}
	}
}

// DoublyLinkedList
func TestDoublyLinkedList(t *testing.T) {
	dll := DoublyLinkedList{}

	if !dll.IsEmpty() {
		t.Errorf("DoublyLinkedList.IsEmpty() = %v, want %v", dll.IsEmpty(), true)
	}

	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	if dll.IsEmpty() {
		t.Errorf("DoublyLinkedList.IsEmpty() = %v, want %v", dll.IsEmpty(), false)
	}

	if dll.Size() != 3 {
		t.Errorf("DoublyLinkedList.Size() = %v, want %v", dll.Size(), 3)
	}

	if dll.Head.Value != 1 {
		t.Errorf("DoublyLinkedList.Head.Value = %v, want %v", dll.Head.Value, 1)
	}

	if dll.Tail.Value != 3 {
		t.Errorf("DoublyLinkedList.Tail.Value = %v, want %v", dll.Tail.Value, 3)
	}

	dll.RemoveHead()

	if dll.Head.Value != 2 {
		t.Errorf("DoublyLinkedList.Head.Value = %v, want %v", dll.Head.Value, 2)
	}

	dll.RemoveTail()

	if dll.Tail.Value != 2 {
		t.Errorf("DoublyLinkedList.Tail.Value = %v, want %v", dll.Tail.Value, 2)
	}
}

// MergeSort
func TestMergeSort(t *testing.T) {
	testCases := []struct {
		array []int
		want  []int
	}{
		{
			array: []int{4, 2, 1, 3, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		// Add more test cases as needed.
	}

	for _, tc := range testCases {
		got := MergeSort(tc.array)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("MergeSort(%v) = %v, want %v", tc.array, got, tc.want)
		}
	}
}

// MinHeap
func TestMinHeap(t *testing.T) {
	h := NewMinHeap()

	if !h.IsEmpty() {
		t.Errorf("Expected heap to be empty")
	}

	h.Add(3)
	h.Add(2)
	h.Add(1)

	if h.Poll() != 1 {
		t.Errorf("Expected 1, got %d", h.Poll())
	}

	if h.Poll() != 2 {
		t.Errorf("Expected 2, got %d", h.Poll())
	}

	if h.Poll() != 3 {
		t.Errorf("Expected 3, got %d", h.Poll())
	}

	if !h.IsEmpty() {
		t.Errorf("Expected heap to be empty")
	}
}

// MaxHeap
func TestMaxHeap(t *testing.T) {
	h := NewMaxHeap()

	if !h.IsEmpty() {
		t.Errorf("Expected heap to be empty")
	}

	h.Add(1)
	h.Add(2)
	h.Add(3)

	if h.Poll() != 3 {
		t.Errorf("Expected 3, got %d", h.Poll())
	}

	if h.Poll() != 2 {
		t.Errorf("Expected 2, got %d", h.Poll())
	}

	if h.Poll() != 1 {
		t.Errorf("Expected 1, got %d", h.Poll())
	}

	if !h.IsEmpty() {
		t.Errorf("Expected heap to be empty")
	}
}

func TestDijkstra(t *testing.T) {
	g := NewGraph()

	node1 := g.AddNode("1")
	node2 := g.AddNode("2")
	node3 := g.AddNode("3")
	node4 := g.AddNode("4")

	g.AddEdge(node1, node2, 1)
	g.AddEdge(node2, node3, 2)
	g.AddEdge(node1, node3, 3)
	g.AddEdge(node3, node4, 4)

	path := g.Dijkstra("1", "4")

	expectedPath := []string{"1", "2", "3", "4"}

	if !reflect.DeepEqual(path, expectedPath) {
		t.Errorf("Expected path to be %v, got %v", expectedPath, path)
	}
}
