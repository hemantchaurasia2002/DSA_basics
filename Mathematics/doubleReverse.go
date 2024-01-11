package main

import "fmt"

func doubleReverse(num int) bool {
    return num == 0 || num % 10 != 0
}

func main() {
    num := 12345
    fmt.Println(doubleReverse(num))
}