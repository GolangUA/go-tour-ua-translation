//go:build OMIT
// +build OMIT

package main

import "fmt"

func main() {
	fmt.Println("Рахуємо")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Завершено")
}
