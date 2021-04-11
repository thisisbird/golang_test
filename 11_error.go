package main

import (
	"fmt"
	"log"
)

type ComedyError string
type OverheatError float64

func (c ComedyError) Error() string {
	return string(c)
}
func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degress!", o)
}
func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}

func main() {
	var err error
	err = ComedyError("What's a xxxxx xxxxx xxxx")
	fmt.Println(err)

	var err2 error = checkTemperature(121.379,100.0)
	if err2 != nil{
		log.Fatal(err2)
	}
}
