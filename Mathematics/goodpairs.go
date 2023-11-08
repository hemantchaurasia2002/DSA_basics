//Number of Good Paris
// Given an array of integers nums, return the number of good pairs.
// A pair (i, j) is called good if nums[i] == nums[j] and i < j.// Example 1:
// Input: nums = [1,2,3,1,1,3]
// Output: 4
// Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.

package main

import "fmt"

func noGoodPairs(nums []int) int {
    count := 0
    for i := 0; i < len(nums)-1; i++ {
        for j := i + 1; j < len(nums); j++ {
            if i < j && nums[i] == nums[j] {
                count++
            }
        }
    }
    return count
}

func main() {
    nums := []int{1, 2, 3, 1, 1, 3}
    result := noGoodPairs(nums)
    fmt.Println(result)
}


func goodpair(nums []int) int {
    var count int 
    
}