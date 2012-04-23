package main

import (
	"image"
	"image/color"
	"tour/pic"
)

type Image struct{
	p [][]uint8
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(i.p[0]), len(i.p))
}

func (i Image) At(x int, y int) color.Color {
	return color.RGBA{i.p[y][x],i.p[y][x],255,255}
}

func Pic(dx, dy int) (p [][]uint8) {
	p = make([][]uint8,dy)
	for y := 0; y < dy; y++ {
		p[y] = make([]uint8,dx)
		for x := 0; x < dx; x++ {
			p[y][x] = uint8(x ^ y)
		}
	}
	return
}

func main() {
	m := Image{Pic(256,256)}
	pic.ShowImage(m)
}