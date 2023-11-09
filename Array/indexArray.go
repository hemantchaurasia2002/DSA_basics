package main 

import (
	"fmt"
)

func valueByIndex(arr []int, n,index int) {
	for i:=0;i<n;i++ {
		if i == index {
			fmt.Printf("Number is %d", arr[i])
		}
	}
}

func main() {
	fmt.Printf("Enter size of your array: ")
	var size,index int
	fmt.Scanln(&size)
	fmt.Printf("Enter the index : ")
	fmt.Scanln(&index)
	var arr = make([]int, size)
	fmt.Printf("Enter the elements : ")
	for i:=0; i<size; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	valueByIndex(arr, size, index)
}