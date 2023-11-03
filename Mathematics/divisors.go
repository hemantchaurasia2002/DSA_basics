package main

import(
	"fmt"
)

func divisors(n int) {
	var i int = 1
	for i*i<=n {
		if n%i==0 {
			fmt.Printf("%d", i)
			if i!=n/i {
				fmt.Printf("%d", n/i)
			}
		}
		i+=1
	}
}

ffunc main(){
	var num int
	fmt.Printf("Enter the value of num : ")
	fmt.Scanf("%d", &num)
	divisors(num)
}