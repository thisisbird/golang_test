package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", viewHandler)
	http.HandleFunc("/english", englishHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("hello,web!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello , english")
}

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}
