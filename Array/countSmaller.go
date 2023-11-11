// Input:
// N = 5
// arr[] = {4 5 3 1 2}
// X = 3
// Output: 2
// Explanation: Here X = 3, and elements that
// are smaller than X are 1 and 2.

package main

import (
	"fmt"
)

func countSmallerNum(arr[] int) {
	var num,count int
	fmt.Printf("Enter the number : ")
	fmt.Scanln(&num)
	for i:=0;i<len(arr);i++ {
		if arr[i] < num {
			count++
		}
	}
	fmt.Println(count)
}

func main() {
	var size int
	fmt.Printf("Enter the size of array : ")
	fmt.Scanf("%d\n", &size)
	var array = make([]int, size)
	for i:=0;i<size;i++ {
		fmt.Printf("Enter the %d element : ", i)
		fmt.Scanf("%d\n", &array[i])
	}
	countSmallerNum(array)
}