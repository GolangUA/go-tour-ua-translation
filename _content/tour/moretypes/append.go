// +build OMIT

package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append працює на 'nil' зрізах.
	s = append(s, 0)
	printSlice(s)

	//  Зріз росте за потреби.
	s = append(s, 1)
	printSlice(s)

	// Одночасно можливо додати більше одного елементу.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
