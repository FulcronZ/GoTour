package main

import "fmt"

func main() {
    sum := 1
    // while in Go is a for wo/ var init
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum)
}
