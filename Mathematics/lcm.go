package main

import(
	"fmt"
)

func gcd(x,y int) int{
	if y==0 {
		return x
	}
	return gcd(y, x%y)
}

func lcm(a,b int) int{
	return a*b/gcd(a,b)

}

func main() {
	var a,b int
	fmt.Printf("Enter the 2 digits : ")
	fmt.Scanf("%d %d", &a,&b)
	fmt.Printf("LCM : %v\n", lcm(a,b))

}