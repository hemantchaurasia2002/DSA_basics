//Recursion.........................
package main

import(
	"fmt"
)

// func legion(){
// 	fmt.Println("LegionX")
// 	legion()
// }

func Legion(n int) {
	if n==0 {
		return
	} else {
		fmt.Printf("%d. LegionX\n",n)
	}
	Legion(n-1)
}

func main(){
	//legion() 
	Legion(10)
}