package main

import "fmt"

func finalTime(arrivalTime, delayedTime int) int {
	return (arrivalTime + delayedTime) % 24
}

func main() {
	arrivalTime := 10
	delayedTime := 5
	time := finalTime(arrivalTime, delayedTime)
	fmt.Println(time)
}

