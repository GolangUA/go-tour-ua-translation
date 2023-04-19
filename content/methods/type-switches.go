// +build OMIT

package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Подвоюємо %v отримуємо %v\n", v, v*2)
	case string:
		fmt.Printf("%q має довжину %v байт\n", v, len(v))
	default:
		fmt.Printf("Нічого не знаю про тип %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
