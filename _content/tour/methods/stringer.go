// +build OMIT

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v років)", p.Name, p.Age)
}

func main() {
	a := Person{"Нестор Махно", 45}
	z := Person{"Леся Українка", 42}
	fmt.Println(a, z)
}
