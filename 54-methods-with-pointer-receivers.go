package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Vertex and *Vertex can call these functions
// But a caller itself is passed by referenced
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())
	
	v2 := Vertex{5, 13}
	v2.Scale(4)
	fmt.Println(v2, v2.Abs())
}