//go:build OMIT

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	// Створюємо map і додаємо в неї елемент за ключем-рядком.
	m = make(map[string]Vertex)
	m["Лабораторії Белла"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Лабораторії Белла"])
}
