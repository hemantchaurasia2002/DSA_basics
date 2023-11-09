package main

import(
	"fmt"
)

func evenOddArray(array []int) {
	var even []int
	var odd []int
	for i:=0;i<len(array);i++ {
		if array[i]%2==0 {
			even = append(even, array[i])
		} else if array[i]%2!=0 {
				odd = append(odd, array[i])
			}
		}
		fmt.Println("Even Number : ", even)
		fmt.Println("Odd Number : ", odd)
}
	

func main() {
		fmt.Printf("Enter size of your array: ")
		var size int
		fmt.Scanln(&size)
		var arr = make([]int, size)
		fmt.Printf("Enter the elements : ")
		for i:=0; i<size; i++ {
		   fmt.Scanf("%d", &arr[i])
		}
		evenOddArray(arr)
}