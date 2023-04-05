//go:build OMIT
// +build OMIT

package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Привіт", "світ")
	fmt.Println(a, b)
}
