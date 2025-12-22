//go:build nobuild || OMIT

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// TODO: побудуйте та поверніть двовимірний зріз значень uint8.
}

func main() {
	pic.Show(Pic)
}
