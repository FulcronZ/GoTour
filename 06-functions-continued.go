package main

import "fmt"

// If a data type of an adjacent paramter is same, it can be omitted.
// func add(x int, y int) int {
func add(x, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(42, 13))
}
