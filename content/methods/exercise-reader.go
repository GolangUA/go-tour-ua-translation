//go:build nobuild || OMIT

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Додайте метод Read([]byte) (int, error) до MyReader.

func main() {
	reader.Validate(MyReader{})
}
