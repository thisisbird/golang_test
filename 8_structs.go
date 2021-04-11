package main

import (
	"fmt"
)

func main() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}
	fmt.Printf("%#v\n", myStruct)

	myStruct.number = 3.14
	myStruct.toggle = true
	myStruct.word = "pie"
	fmt.Println(myStruct.number)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	showInfo(bolts)
	p := minimumOrder("Hex bolts2")
	fmt.Println(p.description, p.count)

	myPart := part{description: "abc", count: 5}
	myPart2 := part{description: "abcd"}
	fmt.Println(myPart)
	fmt.Println(myPart2)

	subscriber := Subscriber{}
	subscriber.HomeAddress.City = "Kaohsiung"
	fmt.Println(subscriber)
	subscriber.HomeAddress.PostalCod = "812"
	fmt.Println(subscriber.HomeAddress)

	//匿名取法
	subscriber2 := Subscriber{name:"Aman Singh"}
	subscriber2.Address.Street = "123 Oak St"
	subscriber2.Address.City = "Omaha"
	fmt.Println(subscriber2.Address.Street)
	fmt.Println(subscriber2.Address.City)

}

//定義名為“part”的型別
type part struct {
	description string
	count       int
}
type Subscriber struct { //若要被引用要大寫
	name        string
	Rate        float64 //若要被引用要大寫
	active      bool
	HomeAddress Address //把結構作為欄位加到別的型別
	Address //把結構作為欄位加到別的型別(匿名)
}

type Rmployee struct {
	Name        string
	Salary      float64
	HomeAddress Address
}

type Address struct {
	Street    string
	City      string
	State     string
	PostalCod string
}

func showInfo(p part) {
	fmt.Println(p.description)
	fmt.Println(p.count)
}

func minimumOrder(description string) part {
	var p part
	p.description = description
	p.count = 100
	return p
}
