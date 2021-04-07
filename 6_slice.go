package main

import (
	"fmt"
)

func main() {
	var notes []string
	notes = make([]string, 7)
	notes[6] = "ti"

	primes := make([]int, 5)
	primes[4] = 5
	fmt.Println(notes, primes)
	fmt.Println(len(notes), len(primes))

	Array1 := []string{"a", "b", "c", "d", "e"}
	i, j := 1, 4
	fmt.Println(Array1[0:3])
	fmt.Println(Array1[i:j])
	fmt.Println(Array1[1:])
	fmt.Println(Array1[:3])

	fmt.Println("--------影響切片---------")
	Array1[2] = "X"
	fmt.Println(Array1[0:3])
	fmt.Println(Array1[i:j])
	fmt.Println(Array1[1:])
	fmt.Println(Array1[:3])
	fmt.Println("--------append---------")
	Array1 = append(Array1, "f", "g", "h")
	fmt.Println(Array1)
}
