package main

import (
	"io/ioutil"
	"net/http"
)

var urls = []string{
	"https://golang.org",
	"https://godoc.org",
	"https://play.golang.org",
	"https://gopl.io",
	"https://golang.org",
	"https://godoc.org",
	"https://play.golang.org",
	"https://gopl.io",
}

func httpGetCode(url string) func() ([]byte, error) {
	return func() ([]byte, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		return ioutil.ReadAll(resp.Body)
	}
}
