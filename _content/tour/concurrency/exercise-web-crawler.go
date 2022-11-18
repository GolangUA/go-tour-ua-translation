// +build OMIT

package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch повертає тіло сторінки `body` отриманої за адресою `url`
	// та зріз `urls` з усіма адресами які були знайдені на сторінці.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl використовує переданий fetcher щоб рекурсивно пройти
// по всім сторінкам починаючи з адреси `url`, до максимальної глибини `depth`.
// вкладеності заданої depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Отримати сторінки для вкладених адрес паралельно.
	// TODO: Не отримувати сторінку з тією ж адресою двічі.
	// Наступне рішення не робить ні те ні інше:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("знайдено: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher це Fetcher який повертає заздалегідь створені данні.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("не знайдено: %s", url)
}

// fetcher це заповнена структура fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
