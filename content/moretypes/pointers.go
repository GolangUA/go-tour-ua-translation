//go:build OMIT

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // вказуємо на i
	fmt.Println(*p) // читаємо i через вказівник
	*p = 21         // змінюємо i через вказівник
	fmt.Println(i)  // бачимо нове значення i

	p = &j         // вказуємо на j
	*p = *p / 37   // ділимо j через вказівник
	fmt.Println(j) // бачимо нове значення j
}
