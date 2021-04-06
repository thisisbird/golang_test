package main

import (
	"fmt"
	"reflect"
)

func main() {
	amount := 6
	fmt.Println(amount)
	fmt.Println(&amount)
	fmt.Println(reflect.TypeOf(&amount))

	var myInt int
	fmt.Println(reflect.TypeOf(&myInt))
	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat))
	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool))
	fmt.Println("----------test------------")
	test()
	fmt.Println("----------test2-------------")
	test2()
	fmt.Println("---------createPointer--------")

	var myFloatPoint *float64 = createPointer()
	fmt.Println(*myFloatPoint)
	fmt.Println("----------createPointer2----------")
	createPointer2(&myBool)

	fmt.Println("----------double2------------")
	double2(&amount)
	fmt.Println(amount)

}

func test() {
	var myInt int
	var myIntPoint *int
	myIntPoint = &myInt
	fmt.Println(myIntPoint)

	var myFloat float64
	var myFloatPoint *float64
	myFloatPoint = &myFloat
	fmt.Println(myFloatPoint)

	var myBool bool
	myBoolPoint := &myBool
	fmt.Println(myBoolPoint)
}

func test2() {
	myInt := 4
	myIntPoint := &myInt
	fmt.Println(myIntPoint)
	fmt.Println(*myIntPoint)

	*myIntPoint = 8
	fmt.Println(*myIntPoint)
	fmt.Println(myInt)
}

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func createPointer2(myFloatPoint *bool) {
	fmt.Println(*myFloatPoint)
}

func double2(number *int) {
	*number *= 2
}
