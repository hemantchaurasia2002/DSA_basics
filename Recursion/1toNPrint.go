package main

import (
	"fmt"
)

func PrintWithoutLoop(n int) {
	if n>=1 {
		PrintWithoutLoop(n-1)
		fmt.Print(" ", n)
	}
}

func main() {
	var num int
	fmt.Printf("Enter the value of N : ")
	fmt.Scanf("%d", &num)
	PrintWithoutLoop(num)
}