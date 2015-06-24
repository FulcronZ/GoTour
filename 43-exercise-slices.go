package main

// Get the package by
//     $ go get golang.org/x/tour/gotour
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    MyPic := make([][]uint8, dy)
    for i:=0; i<dy; i++ {
        MyPic[i] = make([]uint8, dx)
    }

    for i:=0; i<dy; i++ {
        for j:=0; j<dx; j++ {
            MyPic[i][j] = (uint8) (i+j)/2
        }
    }

    return MyPic
}

func main() {
    pic.Show(Pic)
}
