// +build OMIT

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Здійснити зріз зрізу, щоб встановити йому нульову довжину.
	s = s[:0]
	printSlice(s)

	// Збільшити його довжину
	s = s[:4]
	printSlice(s)

	// Відкинути перші два значення.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
