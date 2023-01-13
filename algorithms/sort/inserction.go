package main

import (
	"fmt"
)

func main() {
	numbers := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	for i := 1; i < len(numbers); i++ {
		for j := i; j > 0 && numbers[j-1] > numbers[j]; j-- {
			numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
		}
	}

	fmt.Print(numbers)
}
