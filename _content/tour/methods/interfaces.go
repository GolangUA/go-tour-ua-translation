// +build no-build OMIT

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

	a = f  // Тип MyFloat реалізовує інтерфейс Abser
	a = &v // Тип *Vertex реалізовує інтерфейс Abser

	// Наступна строка кода є невірною, на що компілятор видасть помилку,
	// бо змінна v має тип Vertex (не *Vertex) і таким чином не реалізовує
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
