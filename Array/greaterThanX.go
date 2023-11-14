// Input:
// N = 5
// arr[] = {4 5 3 1 2}
// X = 3
// Output: 2
// Explanation: Here X is 3. The elements
// greater than X are 4 and 5.

package main

import (
	"fmt"
)

func greaterThanX(arr[] int) {
	var count,check int
	fmt.Printf("Enter the value of X : ")
	fmt.Scanf("%d\n", &check)
	for i:=0;i<len(arr);i++ {
		if arr[i] > check {
			count++
		}
	}
	fmt.Printf("Element which are greater than X are : %d", count)

}

func main (){
	var size int
	fmt.Printf("Enter the size of array : ")
	fmt.Scanf("%d\n", &size)
	var array = make([]int, size)
	for i:=0;i<size;i++ {
		fmt.Printf("Enter the %d element : ", i)
		fmt.Scanf("%d\n", &array[i])
	}
	fmt.Printf("Array is : %v\n", array)
	greaterThanX(array)
}