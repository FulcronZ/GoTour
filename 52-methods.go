package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Go does not have classes
// But, you can define methods on struct types
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
}
