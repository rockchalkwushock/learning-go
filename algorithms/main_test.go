package main

import "testing"

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
