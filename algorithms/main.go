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
