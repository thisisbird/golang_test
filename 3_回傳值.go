package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	format()
	dozen := double(6)
	fmt.Println(dozen)
	// errorMsg()
	myInt, myBool, myString := manyReturns()
	fmt.Println(myInt, myBool, myString)
	myInt2, myBool2, myString2 := manyReturns2()
	fmt.Println(myInt2, myBool2, myString2)
	amount, err := paintNeeded(4.2, -3.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f liters needed\n", amount)

}

func format() {
	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("An integer: %d\n", 3.1415)
	fmt.Printf("A string: %s\n", 3.1415)
	fmt.Printf("A boolen: %t\n", 3.1415)
	fmt.Printf("Values: %v %v %v %v %v\n", 3.1415, "\t", true, "", "\n")
	fmt.Printf("Values: %#v %#v %#v %#v %#v\n", 3.1415, "\t", true, "", "\n")
	fmt.Printf("Types: %T %T %T\n", 3.1415, "\t", true)
	fmt.Printf("APercent sign: %%\n")

	fmt.Printf("%%7.3f: %7.3f\n", 12.3456)
	fmt.Printf("%%7.2f: %7.2f\n", 12.3456)
	fmt.Printf("%%7.1f: %7.1f\n", 12.3456)
	fmt.Printf("%%.1f: %.1f\n", 12.3456)
	fmt.Printf("%%.2f: %.2f\n", 12.3456)
}

func double(number float64) float64 {
	return number * 2
}

func errorMsg() {
	err := errors.New("heigh can't be negative")
	fmt.Println(err.Error())
	fmt.Println(err)
	log.Fatal(err)
}

func manyReturns() (int, bool, string) {
	return 1, true, "hello"
}

func manyReturns2() (theInt int, theBool bool, theString string) { //可增加回傳的名稱,方便閱讀
	return 1, true, "hello"
}

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil
}
