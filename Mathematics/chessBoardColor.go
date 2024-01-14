package main

import "fmt"

func squareIsWhite(coordinates string) bool {
	letter := coordinates[0]
	digit := coordinates[1]
	return letter%2 != digit%2
}

func main() {
	coordinates := "a1"
	fmt.Println(squareIsWhite(coordinates))
}

