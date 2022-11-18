// +build no-build OMIT

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("Створюємо нову структуру Reader зі строки")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
