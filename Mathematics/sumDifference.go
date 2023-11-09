package main

import (
	"fmt"
)

func diffOfSums(arr[] int, m int) {
	var result int
	for i := 0; i < len(arr); i++ {
		if i%m == 0 {
			result -= arr[i]
		} else {
			result += arr[i]
		}
	}
	fmt.Printf("Result: %d", result)
}

func main() {
	var divisor int
	fmt.Printf("Enter divisor value : ")
	fmt.Scanln(&divisor)
	var size int
	fmt.Printf("Enter size of your array: ")
	fmt.Scanln(&size)
	var arr = make([]int, size)
	fmt.Printf("Enter the elements : ")
	for i:=0; i<size; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	diffOfSums(arr, divisor)
}

