//go:build nobuild || OMIT

package main

import "fmt"

// fibonacci — функція, яка повертає
// функцію, що повертає int.
func fibonacci() func() int {
	// TODO: реалізуйте замикання, яке генерує послідовність Фібоначчі.
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
