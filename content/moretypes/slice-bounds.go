//go:build OMIT

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	// Вибираємо елементи з індексами 1..3 (4 не включно).
	s = s[1:4]
	fmt.Println(s)

	// Обрізаємо слайс до довжини 2.
	s = s[:2]
	fmt.Println(s)

	// Відкидаємо перший елемент.
	s = s[1:]
	fmt.Println(s)
}
