package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int

func fibonacci() func() int {
    count := 0
    val := make([]int, 2)
    var sum int
    return func() int {
        count++
        switch count-1 {
        case 0:
            val[0] = 1
            return val[0]
        case 1:
            val[1] = 1
            return val[1]
        default:
            sum = val[0] + val[1]
            switch count % 2 {
            case 0:
                val[0] = sum
            default:
                val[1] = sum
            }
            return sum
        }
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(i, f())
    }
}
