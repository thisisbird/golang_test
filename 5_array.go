package main

import (
	"fmt"
)

func main() {
	var notes [7]string
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[2])

	notes2 := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	primes := [5]int{2, 3, 4, 5, 6}
	text := [3]string{
		"qqqqq",
		"aaaaaa",
		"d  a a s s d ",
	}
	fmt.Println(notes2)
	fmt.Println(primes)
	fmt.Println(text)

	fmt.Printf("%#v\n",notes2)
	fmt.Printf("%#v\n",primes)
	fmt.Printf("%#v\n",text)

	for index,value := range notes2{
		fmt.Println(index,value)
	}

	average()
}

func average(){
	number:=[3]float64{71.8,56.2,89.5}
	var sum float64 = 0
	for _,value := range number{
		sum += value
	}
	count := float64(len(number))
	fmt.Printf("平均%0.2f\n",sum/count)
}