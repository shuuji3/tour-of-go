package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var rows = make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		var row = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			row[x] = uint8((x + y) / 2)
		}
		rows[y] = row
	}
	return rows
}

func main() {
	pic.Show(Pic)
}
