package main

import "fmt"

func main() {
    a := make([]int, 5)
    printSlice("a", a)          // len=5 cap=5
    b := make([]int, 0, 5)
    printSlice("b", b)          // len=0 cap=5
    c := b[:2]
    printSlice("c", c)          // len=2 cap=5
    d := c[2:5]
    printSlice("d", d)          // len=3 cap=3
}

func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n",
        s, len(x), cap(x), x)
}
