package main

import (
    "fmt"
    "math/rand"
)

func main() {
    // This is deterministic unless rand.Seed() with a different number
    fmt.Println("My favorite number is", rand.Intn(10))
}
