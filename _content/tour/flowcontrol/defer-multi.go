//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	fmt.Println("рахуємо")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("завершено")
}
