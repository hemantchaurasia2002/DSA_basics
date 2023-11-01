//Program to find greatest common divisor of two numbers by Euclidean Algorithm
package main

import(
	"fmt"
)

func gcd(a,b int) int{
	for a!=b {
		if a>b {
			a=a-b
		}else{
			b=b-a
		}
	}
	return a
}

func main() {
	var a,b int
	fmt.Printf("Enter the 2 digits : ")
	fmt.Scanf("%d %d", &a,&b)
	fmt.Printf("GCD : %v", gcd(a,b))
}