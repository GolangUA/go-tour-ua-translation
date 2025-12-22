//go:build OMIT

package main

import "fmt"

// Index повертає індекс x у s або -1, якщо не знайдено.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v і x мають тип T з обмеженням comparable,
		// тож можна використовувати оператор ==.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index працює зі зрізом int
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index також працює зі зрізом string
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "привіт"))
}
