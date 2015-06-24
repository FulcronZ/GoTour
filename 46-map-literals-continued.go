package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

// if top-level type is just a type name, it can be omitted
var m = map[string]Vertex{
    "Bell Labs":    {40.68433, -74.39967},
    "Google":       {37.42202, -122.08408},
}

func main() {
    fmt.Println(m)
}
