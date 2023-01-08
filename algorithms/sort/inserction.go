package main

import (
	"fmt"
)

func main() {
	numbers := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	for i, v := range numbers {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
