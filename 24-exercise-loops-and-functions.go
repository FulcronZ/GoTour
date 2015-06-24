package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    // Math library sqrt
    // return math.Sqrt(x)
    // Manual Newton's method
    z := 1.0
    for iter:=1; iter<=10; iter++ {
        z = z - ((math.Pow(z, 2) - x ) / (2 * z))
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
}
