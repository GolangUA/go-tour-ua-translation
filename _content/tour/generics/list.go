package main

// List являє собою одиночний зв'язаний список, який містить
// у собі значення будь якого (any) типу.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
}
