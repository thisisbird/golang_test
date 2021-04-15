package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}
//使用Page可以對應URL與Size

func main() {
	pages := make(chan Page)
	urls := []string{"http://example.com", "http://golang.org/", "http://golang.org/doc"}

	for _, url := range urls {
		go responseSize2(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page:= <-pages
		fmt.Printf("%s: %d\n",page.URL,page.Size)
	}
}

func responseSize2(url string, channel chan Page) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL: url, Size: len(body)}
}
