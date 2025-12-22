//go:build nobuild || OMIT

package main

import "golang.org/x/tour/tree"

// Walk проходить дерево t, надсилаючи всі значення
// з дерева в канал ch.
// TODO: реалізуйте проходження та надсилання значень у ch.
func Walk(t *tree.Tree, ch chan int)

// Same визначає, чи дерева
// t1 і t2 містять однакові значення.
// TODO: реалізуйте порівняння, використовуючи Walk.
func Same(t1, t2 *tree.Tree) bool

func main() {
}
