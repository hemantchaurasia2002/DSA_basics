package main

import "fmt"

func numberOfSteps(num int) int {
    cnt := 0
    for num != 0 {
        if num&1 == 1 {
            num--
        } else {
            num /= 2
        }
        cnt++
    }
    return cnt
}

func main() {
    fmt.Println(numberOfSteps(14))
}

