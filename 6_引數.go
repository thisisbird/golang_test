package main

import (
	"fmt"
	"math"
	"os"
)

//輸入 go run xxx.go arg1 arg2 arg3
//arg1 arg2 arg3會被列印出來
func main() {
	fmt.Println(os.Args[1:])

	severalInts(1)
	severalInts(1, 2, 3)
	severalInts()
	max := maximum(12, 13, 9, 1, 55, 0)
	fmt.Println(max)

	new_string := []string{"a", "b", "c"}
	mix(1, true, new_string...)
	mix(1, false, "x","y","z")
}

func severalInts(numbers ...int) {
	fmt.Println(numbers)
}

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1) //無限小
	for _, value := range numbers {
		if value > max {
			max = value
		}
	}
	return max
}

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}
