package main

import "fmt"

func numberOfMatches(n int) int {
    if n > 1 {
        return n/2 + numberOfMatches((n+1)/2)
    }
    return 0
}

func main() {
    n := 10
    fmt.Println(numberOfMatches(n))
}

