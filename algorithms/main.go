package main

import (
	"math"
)

func main() {

}

func LinearSearch(array []int, target int) bool {
	for _, v := range array {
		if v == target {
			return true
		}
	}
	return false
}

func BinarySearch(array []int, target int) bool {
	lo := 0
	hi := len(array)

	for lo < hi {
		m := lo + (hi-lo)/2
		v := array[m]

		if v == target {
			return true
		} else if v > target {
			hi = m
		} else {
			lo = m + 1
		}
	}

	return false
}

func TwoCrystalBalls(floors []bool) int {
	buildingHeight := len(floors)
	interval := int(math.Sqrt(float64(buildingHeight)))
	currentFloor := 0
	prevFloor := 0

	// Drop the first ball at intervals until it breaks
	for currentFloor < buildingHeight {
		if floors[currentFloor] {
			break
		}
		prevFloor = currentFloor
		currentFloor += interval
		if currentFloor > buildingHeight {
			currentFloor = buildingHeight
		}
	}

	// Use the second ball for linear search in the narrowed range.
	for floor := prevFloor + 1; floor < currentFloor && floor < buildingHeight; floor++ {
		if floors[floor] {
			return floor
		}
	}
	return -1
}

func BubbleSort(array []int) []int {
	n := len(array)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if array[j] > array[j+1] {
				array[j+1], array[j] = array[j], array[j+1]
			}
		}
	}
	return array
}

type Queue struct {
	items []int
}

func (q *Queue) New() *Queue {
	return &Queue{}
}

func (q *Queue) Dequeue() int {
	item, items := q.items[0], q.items[1:]
	q.items = items
	return item
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Peek() int {
	return q.items[0]
}

type Stack struct {
	items []int
}

func (s *Stack) New() *Stack {
	return &Stack{}
}

func (s *Stack) Pop() int {
	item, items := s.items[len(s.items)-1], s.items[:len(s.items)-1]
	s.items = items
	return item
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Peek() int {
	return s.items[len(s.items)-1]
}

// ArrayBuffer

type ArrayBuffer struct {
	data []byte
}

func NewArrayBuffer(size int) *ArrayBuffer {
	return &ArrayBuffer{
		data: make([]byte, size),
	}
}

func (ab *ArrayBuffer) Get(index int) byte {
	return ab.data[index]
}

func (ab *ArrayBuffer) Set(index int, value byte) {
	ab.data[index] = value
}

func (ab *ArrayBuffer) Length() int {
	return len(ab.data)
}

// ArrayList

type ArrayList struct {
	data []interface{}
}

func NewArrayList() *ArrayList {
	return &ArrayList{
		data: make([]interface{}, 0),
	}
}

func (al *ArrayList) Add(value interface{}) {
	al.data = append(al.data, value)
}

func (al *ArrayList) Get(index int) interface{} {
	return al.data[index]
}

func (al *ArrayList) Size() int {
	return len(al.data)
}

// Maze Solver
type Point struct {
	X, Y int
}

var directions = []Point{
	{0, 1},  // down
	{1, 0},  // right
	{0, -1}, // up
	{-1, 0}, // left
}

// func printMaze(maze [][]int) {
// 	for _, row := range maze {
// 		for _, col := range row {
// 			fmt.Printf("%d ", col)
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println()
// }

func MazeSolver(maze [][]int, current, end Point) bool {
	if current.X < 0 || current.Y < 0 || current.X >= len(maze) || current.Y >= len(maze[0]) {
		return false
	}

	if maze[current.X][current.Y] == 0 {
		return false
	}

	if current == end {
		return true
	}

	maze[current.X][current.Y] = 0
	// printMaze(maze)

	for _, d := range directions {
		next := Point{current.X + d.X, current.Y + d.Y}
		if MazeSolver(maze, next, end) {
			return true
		}
	}

	return false
}

// QuickSort
func QuickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	pivot := array[0]
	var less, greater []int

	for _, v := range array[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	less = QuickSort(less)
	greater = QuickSort(greater)

	return append(append(less, pivot), greater...)
}

// Doubly Linked List
type Node struct {
	Next     *Node
	Previous *Node
	Value    interface{}
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (ll *DoublyLinkedList) Append(value interface{}) {
	node := &Node{Value: value}

	if ll.Head == nil {
		ll.Head = node
	} else {
		ll.Tail.Next = node
		node.Previous = ll.Tail
	}

	ll.Tail = node
}

func (ll *DoublyLinkedList) IsEmpty() bool {
	return ll.Head == nil
}

func (ll *DoublyLinkedList) Prepend(value interface{}) {
	node := &Node{Value: value}

	if ll.Head == nil {
		ll.Head = node
	} else {
		ll.Head.Previous = node
		node.Next = ll.Head
	}

	ll.Head = node
}

func (ll *DoublyLinkedList) Remove(value interface{}) {
	current := ll.Head

	for current != nil {
		if current.Value == value {
			if current == ll.Head {
				ll.Head = current.Next
				ll.Head.Previous = nil
			} else if current == ll.Tail {
				ll.Tail = current.Previous
				ll.Tail.Next = nil
			} else {
				current.Previous.Next = current.Next
				current.Next.Previous = current.Previous
			}
		}
		current = current.Next
	}
}

func (ll *DoublyLinkedList) RemoveTail() {
	if ll.Tail != nil {
		ll.Tail = ll.Tail.Previous
		ll.Tail.Next = nil
	}
}

func (ll *DoublyLinkedList) RemoveHead() {
	if ll.Head != nil {
		ll.Head = ll.Head.Next
		ll.Head.Previous = nil
	}
}

func (ll *DoublyLinkedList) Search(value interface{}) *Node {
	current := ll.Head

	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}

	return nil
}

func (ll *DoublyLinkedList) Size() int {
	count := 0
	current := ll.Head

	for current != nil {
		count++
		current = current.Next
	}

	return count
}
