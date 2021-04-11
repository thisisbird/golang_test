package main

import "fmt"

func main() {
	one()
}

func one(){
	defer fmt.Println("defer in one()")
	two()
}

func two(){
	defer fmt.Println("defer in two()")
	panic("panic!!")
}