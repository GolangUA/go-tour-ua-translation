// +build OMIT

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // вказівник на i
	fmt.Println(*p) // читання i через вказівник
	*p = 21         // задати i через вказівник
	fmt.Println(i)  // вивсети нове значення i

	p = &j         // вказівник на j
	*p = *p / 37   // ділення j через вказівник
	fmt.Println(j) // вивсети нове значення j
}
