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
