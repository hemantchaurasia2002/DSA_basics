package main

import (
	"fmt"
)


func reverseCountNum(n int) {
	if n == 0 {
		fmt.Printf("-----------------------\n")
		return
	}
	fmt.Printf("%d\n", n)
	reverseCountNum(n-1)
	fmt.Printf("%d\n", n)
	
}

// func countNum(n int) {
// 	if n <= 10 {
// 		fmt.Printf("%d\n", n)
// 		countNum(n+1)
// 	}
// }

func main(){
	
	// var num int = 1
	// countNum(num)
	var x int = 10
	
	reverseCountNum(x)
}