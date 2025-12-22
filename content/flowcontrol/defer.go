//go:build OMIT

package main

import "fmt"

func main() {
	// Інструкція `defer` відкладає виконання виклику до моменту, коли поточна функція завершиться.
	defer fmt.Println("світе")

	fmt.Println("привіт")
}
