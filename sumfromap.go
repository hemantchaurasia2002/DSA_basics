package main

import (
	"fmt"
	"time"
)

func sumOfNaturalNumbers(numberOfElement int) int{
	sum := 0
	sum = (numberOfElement * (numberOfElement + 1) / 2)
	return sum

}

func main(){

	//Starting time
	start := time.Now()

	time.Sleep(time.Second * 2)
	//code for measuring time
	numberOfElement := 100
	totalSum := 0
	totalSum = sumOfNaturalNumbers(numberOfElement)
	fmt.Println("Total sum - ", totalSum)

	// calculating time and printing
	elapsed := time.Since(start)
	fmt.Printf("Program took %s", elapsed)


}