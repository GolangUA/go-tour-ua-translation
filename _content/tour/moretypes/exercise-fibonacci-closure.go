// +build no-build OMIT

package main

import "fmt"

// fibonacci це функція, яка повертає
// функцію, яка повертає int.
func fibonacci() func() int {
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
