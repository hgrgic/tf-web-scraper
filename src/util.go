package src

import (
	"github.com/anaskhan96/soup"
	"os"
)

func ReadUrl(url string) soup.Root{
	resp, err := soup.Get(url)
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	return doc
}
