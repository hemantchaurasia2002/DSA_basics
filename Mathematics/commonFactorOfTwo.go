package main

import (
	"fmt"
)

func commonFactors(a int, b int) {
	var ans int = 0
	g := gcd(a, b)
	for x := 1; x <= g; x++ {
		if g%x == 0 {
			ans++
		}
	}
	fmt.Printf("The common factor of %v and %v is %v\n", a,b,ans)
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	var a,b int
	fmt.Printf("Enter the value of A : ")
	fmt.Scanf("%d", &a)
	fmt.Printf("Enter the value of B : ")
	fmt.Scanf("%d", &b)
	gcd(a,b)
	commonFactors(a,b)
}