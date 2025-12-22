//go:build OMIT

package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append працює з nil-слайсами.
	s = append(s, 0)
	printSlice(s)

	// Слайс збільшується за потреби.
	s = append(s, 1)
	printSlice(s)

	// Можна додати більше ніж один елемент за раз.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("довжина=%d місткість=%d %v\n", len(s), cap(s), s)
}
