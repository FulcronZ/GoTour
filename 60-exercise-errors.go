package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return 0.0, ErrNegativeSqrt(x) 
	}
	
	// copied from 24-exercise-loops-and-functions
	z := 1.0
    for iter:=1; iter<=10; iter++ {
        z = z - ((math.Pow(z, 2) - x ) / (2 * z))
    }
	return z, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// if not convert e to float64, it causes infinite loop
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}