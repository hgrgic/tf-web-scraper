package main

import (
	"fmt"
	"tf-scrapper/src"
)

func main() {
	l := src.ScrapeLeague("nebitno")
	fmt.Println(l)
}
