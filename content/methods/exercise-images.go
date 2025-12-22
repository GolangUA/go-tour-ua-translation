//go:build nobuild || OMIT

package main

import "golang.org/x/tour/pic"

type Image struct{}

// TODO: Реалізуйте методи ColorModel(), Bounds() і At(x, y int) color.Color.
// TODO: Поверніть прямокутник image.Rect(0, 0, w, h) у Bounds і color.RGBAModel у ColorModel.
// TODO: Використовуйте значення пікселя для формування color.RGBA{v, v, 255, 255} у At.
func main() {
	m := Image{}
	pic.ShowImage(m)
}
