package main

import "fmt"

func main() {
	fmt.Println(recover())
	defer calmDown()
	freakOut()
	panic("QQQ")
}
func calmDown(){
	fmt.Println(recover())
}
func freakOut(){
	defer calmDown()
	panic("oh no")
	fmt.Println("I won't be run")
}