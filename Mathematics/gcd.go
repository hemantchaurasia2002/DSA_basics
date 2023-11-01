//Program to find greatest common divisor of two numbers by Euclidean Algorithm
package main

import(
	"fmt"
)

func gcd(x,y int) int{
	for x!=y {
		if x>y {
			x=x-y
		}else{
			y=y-x
		}
	}
	return x
}

func gcdOptimized(x,y int) int{
	if y==0 {
		return x
	}
	return gcd(y, x%y)
}

func main() {
	var a,b int
	fmt.Printf("Enter the 2 digits : ")
	fmt.Scanf("%d %d", &a,&b)
	fmt.Printf("GCD : %v\n", gcd(a,b))
	fmt.Printf("Optimized GCD : %v", gcdOptimized(a,b))
}