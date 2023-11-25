package main

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
