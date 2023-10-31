//Program to find sum of N natural numbers
package main

import (
	"fmt"
)

func sumOfN(n int) int {
	var sum int
	sum = n*(n+1)/2
	return sum
}

func main() {
	var num int
	fmt.Printf("Enter the value of N : ")
	fmt.Scanf("%d", &num)
	fmt.Printf("Sum of first %d Natural Numbers are : %v",num,sumOfN(num))
}