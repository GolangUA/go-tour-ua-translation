// +build no-build OMIT

package main

import "golang.org/x/tour/tree"

// Walk проходить по дереву `t` та надсилає всі значення
// з дерева в канал `ch`.
func Walk(t *tree.Tree, ch chan int)

// Same визначає чи два дерева
// `t1` та `t2` містять однакові значення.
func Same(t1, t2 *tree.Tree) bool

func main() {
}
