package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Print("enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n') //可用 _ 忽略用不到的變數
	input, err := reader.ReadString('\n') //顯示斷行以前的文字(包含斷行)
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)            //移除空白
	grade, err := strconv.ParseFloat(input, 64) //文字轉float64
	var status string                           //if迴圈才宣靠變數,迴圈結束就拿不到了,所以要在迴圈外宣告
	if grade >= 60 {
		status = "pass"
	} else {
		status = "failing"
	}

	fmt.Println(input, status)
	randomNum()
	for1()
	for2()
	for3()
	format()
}

func randomNum() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println(target)
}

func for1() {
	for x := 1; x < 3; x++ {
		fmt.Println(x)
	}
}
func for2() {
	x := 1
	for x < 3 {
		fmt.Println(x)
		x++
	}
}
func for3() {
	var x int
	for x := 1; x < 3; x++ {
		fmt.Println(x)
	}
	fmt.Println(x)
}


