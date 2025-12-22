//go:build OMIT

package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// Створюємо слайс (slice), взявши піддіапазон базового масиву.
	var s []int = primes[1:4]
	fmt.Println(s)
}
