//go:build OMIT
// +build OMIT

package main

import "fmt"

// Index повертає індекс змінної "x" у слайсі "s", або -1 якщо "х" не було знайдено.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// "v" та "x" являються типами T, які мають обмеження 'comparable',
		// то ж ми можемо використовувати порівняння через == у данному випадку.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// 'Index' працює зі слайсом типу int
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// 'Index' також працює зі слайсом типу string
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}
