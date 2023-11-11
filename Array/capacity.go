package main

import (
	"fmt"
)

func main() {
	var numbers[] int
	for i := 0; i < 4; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println(numbers)
}