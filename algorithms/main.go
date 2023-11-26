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
