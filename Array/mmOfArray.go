package main

import (
	"fmt"
)

func mean(arr[] int) {
	var sum int = 0
	for i:=0;i<len(arr);i++ {
		sum+=arr[i]
	}
	fmt.Printf("Mean of %v is : %d", arr,sum)
}

func median(arr[]int ,  size int) {
	for i:=0;i<len(arr)-1;i++ {
		for j:=0;j<len(arr)-1-i;j++ {
			if arr[j] > arr[j+1] {
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		
		}
	}
	fmt.Printf("Sorted array : %v\n", arr)
	var median int
	if size%2==0 {
		median=arr[size/2]
	} else{
		median=(arr[size/2]+(arr[size/2]+1))/2
	}
	fmt.Printf("Median : %d", median)

}


func main() {
	var size int
	fmt.Printf("Enter the size of the array : ")
	fmt.Scanf("%d\n", &size)
	var array = make([]int, size)
	for i:=0;i<size;i++ {
		fmt.Printf("Enter the %d element : ", i)
		fmt.Scanf("%d\n", &array[i])
	}
	fmt.Println(array)
	fmt.Printf("\n")
	mean(array)
	fmt.Printf("\n")
	median(array, size)

}
