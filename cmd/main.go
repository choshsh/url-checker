package main

import (
	"fmt"
	"url-checker/pkg/urlChecker"
)

func main() {
	urlChecker := urlChecker.NewUrlChekcer("1")
	fmt.Println(urlChecker)
}
