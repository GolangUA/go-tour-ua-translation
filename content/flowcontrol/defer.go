//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	defer fmt.Println("світ")

	fmt.Println("Привіт")
}
