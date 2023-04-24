//go:build ignore || OMIT
// +build ignore OMIT

package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // тип MyFloat реалізовує інтерфейс Abser
	a = &v // тип *Vertex реалізовує інтерфейс Abser

	// Наступний рядок коду є невірним, на що компілятор видасть помилку.
	// Змінна v має тип Vertex (не *Vertex) і таким чином не реалізовує
	// інтерфейс Abser, тому неможливо виконати присвоювання до змінної а,
	// яка має інтерфейсний тип Abser.
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
