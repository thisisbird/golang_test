package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	sizes := make(chan int)
	urls := []string{"http://example.com", "http://golang.org/", "http://golang.org/doc"}

	for _, url := range urls {
		go responseSize(url, sizes)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-sizes)
	}
}

func responseSize(url string, channel chan int) {
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
	channel <- len(body)
}
