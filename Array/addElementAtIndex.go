//program to update the element of an array at a particular index value
//Input:
//n = 5
//array = [1 2 3 4 5]
//index = 4,element = 8
//Output = [1 2 3 4 8]
package main

import (
	"fmt"
)


func elementAtIndex(array[] int, index, value int) {
	for i:=0;i<len(array);i++ {
		if i==index {
			array[i]=value
		}
	}
	fmt.Printf("ARRAY is : %v", array)
}

func main() {
	var size,i,value int
	fmt.Printf("Enter the index : ")
	fmt.Scanln(&i)
	fmt.Printf("Enter size of your array: ")
	fmt.Scanln(&size)
	var arr = make([]int, size)
	fmt.Printf("Enter the replacement value : ")
	fmt.Scanln(&value)
	fmt.Println("Enter the element : ")
	for i:=0; i<size;i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("Old Array : ", arr)

	elementAtIndex(arr,i,value)
}