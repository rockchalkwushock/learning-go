package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)

	answer := 1
	// char is represented as rune here (int32)
	for _, char := range input {
		if char >= 'A' && char <= 'Z' {
			answer++
		}
		// str := string(char)
		// if strings.ToUpper(str) == str {
		// 	answer++
		// }
	}

	fmt.Println(answer)
}
