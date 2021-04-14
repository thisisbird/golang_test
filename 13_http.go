package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	response,err := http.Get("http://example.com")
	if err!=nil{
		log.Fatal(err)
	}
	defer response.Body.Close()
	body,err:=ioutil.ReadAll(response.Body)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(string(body))

	go responseSize("http://example.com")
	go responseSize("http://golang.org/")
	go responseSize("http://golang.org/doc")
	time.Sleep(5 * time.Second) //暫停5秒
}

func responseSize(url string){
	fmt.Println("Getting",url)
	response,err := http.Get(url)
	if err!=nil{
		log.Fatal(err)
	}
	defer response.Body.Close()
	body,err:=ioutil.ReadAll(response.Body)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(len(body))
}