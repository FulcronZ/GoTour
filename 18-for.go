package main

import "fmt"

func main() {
    sum := 0
    // Careful. Go's for statement doesn't have () as C
    for i:= 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)
}
