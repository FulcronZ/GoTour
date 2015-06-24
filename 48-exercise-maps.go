package main

import (
    "golang.org/x/tour/wc"
)

import "strings"

func WordCount(s string) map[string]int {
    words := strings.Fields(s)      // split by whitespace char
    MyMap := make(map[string]int)
    for _, v := range words {
        MyMap[v] += 1
    }
    return MyMap
}

func main() {
    wc.Test(WordCount)
}
