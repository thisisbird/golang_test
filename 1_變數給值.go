/**
* * go run xxx.go  		直接執行
* ? go fmt xxx.go		格式化檔案(建議 非必要)
* ? go build xxx.go		打包檔案
* ! go /.xxx			執行檔案
*/
package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	var quantity int
	var length, width float64
	var customerName string

	// quantity = 5 //若不給值 預設0,false,空字串
	length, width = 1.2, 2.4
	customerName = "qqqqq"

	fmt.Println("hello")
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("head"))
	fmt.Println(reflect.TypeOf("head"))

	fmt.Println(quantity, length*width, customerName)
	fmt.Println("-----------------------------")
	shortVar()
	fmt.Println("-----------------------------")
	changeType()
}

func shortVar() { //短命名
	var quantity int = 4
	length, width := 1.2, 2.4
	length = 2.2
	customerName := "shortVar"
	fmt.Println(quantity, length*width, customerName)
}

func changeType() {
	myInt := 2
	fmt.Println(reflect.TypeOf(myInt))
	fmt.Println(reflect.TypeOf(float64(myInt)))
}
