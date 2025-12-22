//go:build OMIT

package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Двічі %v дорівнює %v\n", v, v*2)
	case string:
		fmt.Printf("%q має довжину %v байтів\n", v, len(v))
	default:
		fmt.Printf("Не знаю, що це за тип %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
