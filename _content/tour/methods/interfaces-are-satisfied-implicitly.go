//go:build OMIT
// +build OMIT

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// Наявність метода M() означає, що даний тип T реалізовує інтерфейс I;
// зауважте, що ніяким чином не потрібно ЯВНО вказувати, що тип T реалізує конкретний інтерфейс I.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"вітаю"}
	i.M()
}
