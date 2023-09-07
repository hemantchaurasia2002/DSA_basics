package main

import (
	"fmt"
	"time"
)

func main (){

	//Starting time
	start := time.Now()

	time.Sleep(time.Second * 2)
	//code for measuring time
	numberOfElement := 100
	sumOfNaturalNumbers := 0
	for i := 1; i <= numberOfElement; i++ {
		sumOfNaturalNumbers += i
	}
	fmt.Println("Total Sum - ", sumOfNaturalNumbers) 
	// calculating time and printing
	elapsed := time.Since(start)
 	fmt.Printf("Program took %s", elapsed)

}

