//go:build OMIT

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Обріжемо слайс до нульової довжини.
	s = s[:0]
	printSlice(s)

	// Розширимо його довжину.
	s = s[:4]
	printSlice(s)

	// Відкинемо перші два значення.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("довжина=%d місткість=%d %v\n", len(s), cap(s), s)
}
