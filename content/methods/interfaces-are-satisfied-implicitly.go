//go:build OMIT

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// Цей метод означає, що тип T реалізує інтерфейс I,
// але нам не потрібно оголошувати це явно.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"привіт"}
	i.M()
}
