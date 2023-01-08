package main

import (
	"fmt"
)

func main() {
	numbers := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[i] < numbers[j] {
				var aux = numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = aux
			}
		}

	}

	fmt.Print(numbers)
}
