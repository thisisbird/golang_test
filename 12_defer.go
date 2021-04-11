package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// err := Socialize()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	numbers, err := getFloat("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(numbers)
}
func Socialize() error {
	defer fmt.Println("Goodbye") //延遲調用
	fmt.Println("Hello")
	return fmt.Errorf("I dont want to talk")
}

func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("opening", fileName)
	return os.Open(fileName)
}
func CloseFile(file *os.File) {
	fmt.Println("closing file")
	file.Close()
}

func getFloat(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	defer CloseFile(file) // 在getFloat結束前調用
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}
