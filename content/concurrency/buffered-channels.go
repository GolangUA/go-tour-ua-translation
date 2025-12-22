//go:build OMIT

package main

import "fmt"

func main() {
	ch := make(chan int, 2) // буферизований канал місткістю 2
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
