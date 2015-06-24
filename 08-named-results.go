package main

import "fmt"

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return              // this is called "naked return"
}

func main() {
    fmt.Println(split(17))
}
