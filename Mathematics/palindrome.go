package main 

import (
	"fmt"
)

func isPal(n int) bool{
	var reverse int = 0
	var temp int = n
	var le int
	for temp!=0 {
		le = temp % 10
		reverse = reverse*10+le
		temp = temp/10
	}
	return reverse == n
}

func main() {
	var num int
	fmt.Printf("Enter the value of N : ")
	fmt.Scanf("%d", &num)
	fmt.Printf("%v", isPal(num))
}