package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/calendar"
)

func main() {
	event := calendar.Event{}
	// event.Date.year = 2019
	fmt.Println(event)
	value := calendar.MyType{}
	value.ExportedMethod() //調用晉升方法

	err := event.SetTitle("1414141")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())
}
