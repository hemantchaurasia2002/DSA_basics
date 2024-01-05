package main

import "fmt"

func binarySearch(arr []int, low int, high int, number int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2

	if number == arr[mid] {
		return mid
	} else if number < arr[mid] {
		return binarySearch(arr, low, mid-1, number)
	} else {
		return binarySearch(arr, mid+1, high, number)
	}
}

func main() {
	var number int
	var size int
	fmt.Print("Enter the value of arr size: ")
	fmt.Scan(&size)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Print("Enter array value: ")
		fmt.Scan(&arr[i])
	}
	fmt.Print("Enter the target value: ")
	fmt.Scan(&number)
	low := 0
	high := len(arr) - 1
	index := binarySearch(arr, low, high, number)
	if index != -1 {
		fmt.Printf("Element found at index %d", index)
	} else {
		fmt.Println("Element not found in the array")
	}
}

