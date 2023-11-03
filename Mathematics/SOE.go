// program of Sieve of Eratosthenes
package main

import(
	"fmt"
	"math"
)

func primeNum(num int) bool{
	if num < 2 {
	   fmt.Println("Number must be greater than 2.")
	   return false
	}
	sq_root := int(math.Sqrt(float64(num)))
	for i:=2; i<=sq_root; i++{
	   if num % i == 0 {
		  return false
	   }
	}
	return true
}

func SOE(n int) {
	for i:=2;i<=n;i++ {
		if primeNum(i)==true {
			fmt.Printf("%d ", i)
		}
	}
	return
}

func main(){
	var num int
	fmt.Printf("Enter the value of num : ")
	fmt.Scanf("%d", &num)
   	SOE(num)
}