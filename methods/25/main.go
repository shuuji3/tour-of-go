package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	x0, y0, x1, y1 int
}

// ColorModel returns the Image's color model.
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (i Image) Bounds() image.Rectangle {
	return image.Rect(i.x0, i.y0, i.x1, i.y1)
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), uint8(x + y), 255}
}

func main() {
	m := Image{0, 0, 200, 200}
	pic.ShowImage(m)
}
