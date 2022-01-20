package main

import (
	"flag"
	"log"
	"url-checker/pkg/urlChecker"
)

func main() {
	var url string
	var count int

	// cmd 파라미터 매핑
	flag.StringVar(&url, "url", "", "체크하고 싶은 url 입력 (Ex. https://www.google.com")
	flag.IntVar(&count, "count", 3, "체크할 횟수")
	flag.Parse()
	if len(url) == 0 {
		log.Fatal("URL 값이 비어있음")
	}

	urlChecker := urlChecker.NewUrlChekcer(url, count)
	urlChecker.Print()

}
