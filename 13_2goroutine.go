package main

import (
	"fmt"
	"time"
)
func a(){
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}
func b(){
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}
func main() {
	//go 陳述句無法使用回傳值
	go a()
	go b()
	time.Sleep(5 * time.Second) //暫停5秒
	fmt.Println("end!")

	
}