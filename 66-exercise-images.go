package main

import "golang.org/x/tour/pic"
import (
	"image"
	"image/color"	
)

type Image struct{}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

// set image size
func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

// fill color in each pixel
func (im Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x+x), uint8(x+y), uint8(y+y), 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}