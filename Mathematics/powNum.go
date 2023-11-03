package main

import(
	"fmt"
)

func powerNum(x,y int) {
	result:= 1
    for y != 0 {
        result *= x
        y -= 1
    }
    fmt.Printf("POWER : %d", result)
	return
}


func main() {
    var exponent, base int
    fmt.Print("Enter Base:")
    fmt.Scanln( & base)
    fmt.Print("Enter exponent:")
    fmt.Scanln( & exponent)
	powerNum(base,exponent)
}