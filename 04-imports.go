package main

// It's a good style to factor import statement
import (
    "fmt"
    "math"
)

// One can import a package one by one like below
// import "fmt"
// import "math"

func main() {
    fmt.Printf("Now you have %g problems.", math.Nextafter(2,3))
}
