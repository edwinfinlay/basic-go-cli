package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
	url string
)

func main() {
	flag.StringVar(&url, "url", "", "URL to check")

	flag.Parse()

	if url != "" {
		checkURL(url)
	}
}

func checkURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting url %s: %v", url, err)
		return
	}
	fmt.Printf("Checked url: %s, Status: %d", url, resp.StatusCode)
}
