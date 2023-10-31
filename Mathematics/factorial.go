//Program to find factorial of N
package main

import (
	"fmt"
)

func factOfN(n int) int{
	var result int = 1
	for i:=2;i<n+1;i++{
		result=result*i
	}
	return result
}

func main() {
	var n int
	fmt.Printf("Enter the value of n : ")
	fmt.Scanf("%d", &n)
	fmt.Printf("Factorial of %d is %v", n,factOfN(n))
}