package main

import "fmt"

// function to print the array with array and
// size of the array as argument
func printArray(array []int, size int) {
    for i := 0; i < size; i++ {
        fmt.Print(array[i], " ")
    }
    fmt.Println()
}

// binary search function to find an element in the array
func binarySearch(array []int, size int, toFind int) int {
    var left, right, m int
    left = 0
    right = size

    for left < right {
        m = (right + left) / 2
        if array[m] == toFind {
            // if the element at index m is equal to the element we
            // are searching then returning m
            return m
        } else if array[m] > toFind {
            //If the element at index m is greater than the element
            //We are searching then we are skipping the right side
            // of the array
            right = m - 1
        } else {
            //If the element at index m is less than the element
            //We are searching then we are skipping the left side
            // of the array
            left = m
        }
    }

    return -1
}

func main() {
    // decalring and initializing the
    // array using the shorthand method
    array := []int{20, 44, 45, 54, 67, 88, 91}

    // decalring and initializing the
    // variable using the var keyword
    var toSearch int
    toSearch = 3

    fmt.Println("Golang program to find an element in an array using binary search using for loop.")
    fmt.Print("array: ")
    printArray(array, 6)
    // linear search function passing array and
    // variable as a parameter
    index := binarySearch(array, 6, toSearch)

    if index == -1 {
        fmt.Println(toSearch, "is not present in the array.")
    } else {
        fmt.Println(toSearch, "is present at index", index, "in the array.")
    }

}