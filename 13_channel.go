package main

import (
	"fmt"
	"time"
)

func main() {
	myChannel := make(chan string)
	go greeting(myChannel)
	receivedValue := <-myChannel
	// fmt.Println(<-myChannel)
	fmt.Println(receivedValue)

	channel1 := make(chan string)
	channel2 := make(chan string)
	go abc(channel1)
	go def(channel2)

	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Println("-----------------------------")
	myChannelNew := make(chan string)
	go send(myChannelNew)
	reportNap("receiving goroutine",5)
	fmt.Println(<-myChannelNew)
	fmt.Println(<-myChannelNew)

}
func greeting(myChannel chan string) {
	myChannel <- "hi"
}

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}
func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}
func reportNap(name string,delay int){
	for i := 0; i < delay; i++ {
		fmt.Println(name,"sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name,"wakes up!")
}
func send(myChannel chan string){
	reportNap("sending goroutine",2)
	fmt.Println("****sending value****1")
	myChannel <- "aa"
	fmt.Println("****sending value****2")
	myChannel <- "bb"

}