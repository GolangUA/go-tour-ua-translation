// +build OMIT

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // має тип Vertex
	v2 = Vertex{X: 1}  // Y:0 є неявним
	v3 = Vertex{}      // X:0 та Y:0
	p  = &Vertex{1, 2} // має тип *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
