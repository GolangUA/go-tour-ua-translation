// +build OMIT

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Відповідь"] = 42
	fmt.Println("Значення:", m["Відповідь"])

	m["Відповідь"] = 48
	fmt.Println("Значення:", m["Відповідь"])

	delete(m, "Відповідь")
	fmt.Println("Значення:", m["м"])

	v, ok := m["Відповідь"]
	fmt.Println("Значення:", v, "Присутнє?", ok)
}
