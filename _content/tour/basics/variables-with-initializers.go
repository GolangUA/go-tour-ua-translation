//go:build OMIT
// +build OMIT

package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "ні!"
	fmt.Println(i, j, c, python, java)
}
