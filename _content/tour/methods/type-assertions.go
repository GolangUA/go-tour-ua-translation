// +build no-run OMIT

package main

import "fmt"

func main() {
	var i interface{} = "вітаю"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // паніка!
	fmt.Println(f)
}
