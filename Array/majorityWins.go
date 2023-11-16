package main

import (
	"fmt"
)

func frequency(arr[] int) {
	var x,y int
	var xCount,yCount int
	fmt.Printf("Enter the value of X :")
	fmt.Scanf("%d\n", &x)
	fmt.Printf("Enter the value of Y :")
	fmt.Scanf("%d\n", &y)
	for i:=0;i<len(arr);i++ {
		if x==arr[i] {
			xCount++
		} else if y==arr[i] {
			yCount++
		} else {
			fmt.Println("No majority found")
		}
	}
	fmt.Printf("Frequency of %d : %d\nFrequency of %d : %d", x,xCount,y,yCount)

}

func main() {
	var size int
	fmt.Printf("Enter the size : ")
	fmt.Scanf("%d\n", &size)
	var array = make([]int, size)
	for i:=0;i<size;i++ {
		fmt.Printf("Enter the %d element : ", i)
		fmt.Scanf("%d\n", &array[i])
	}
	fmt.Printf("Array is : %v\n", array)
	frequency(array)

}