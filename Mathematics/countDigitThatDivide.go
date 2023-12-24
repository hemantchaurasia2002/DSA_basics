package main 

import (
	"fmt"
)

func countDigits(num int) {
    temp := num
    var count int = 0
    for temp!=0 {
        rem := temp%10
        if (num%rem)==0 {
            count++
        }
        temp = temp/10
    }
    fmt.Printf("Output : %v\n", count)    
}

func main() {
	var n int
	fmt.Printf("Enter the value of digit : ")
	fmt.Scanf("%d", &n)
	countDigits(n)
}