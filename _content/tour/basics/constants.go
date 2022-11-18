//go:build OMIT
// +build OMIT

package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "світ"
	fmt.Println("Привіт", World)
	fmt.Println("Добрий", Pi, "день")

	const Truth = true
	fmt.Println("Go рулить?", Truth)
}
