// +build no-build OMIT

package main

import "fmt"

// fibonacci це функція, що повертає
// функцію, що повертає int.
func fibonacci() func() int {
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
