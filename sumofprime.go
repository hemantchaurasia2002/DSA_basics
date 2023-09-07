package main

import (
	"fmt"
	"time"
)

func primeNumbers(N int) {
   for i := 2; i <= N; i++ {
			isPrime := true
      

   for j := 2; j <= i/2; j++ {
      if i%j == 0 {
         isPrime = false
         break
      }
   }
   if isPrime {
         fmt.Println(i)
      }
   }
}
func main() {

	//Starting time
	start := time.Now()

	time.Sleep(time.Second * 2)

  var n int = 101
  primeNumbers(n)

	// calculating time and printing
	elapsed := time.Since(start)
 	fmt.Printf("Program took %s", elapsed)
}