//Program to count the digit of a number
package main

import (
	"fmt"
)

func digitCounter(x int) int{
	var result int = 0
	for x>0 {
		x=x/10
		result++
	}
	return result
}

func main() {
	var num int
	fmt.Printf("Enter the value of N : ")
	fmt.Scanf("%d", &num)
	fmt.Printf("The number of digits in %d are : %v",num,digitCounter(num))
}