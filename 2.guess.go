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
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	traget := rand.Intn(100) + 1
	fmt.Println("guess between 1 and 100")

	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("you have", 10-guesses, "guesses left")
		fmt.Print("Mak a guess:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)  //移除換行
		guess, err := strconv.Atoi(input) //轉整數
		if err != nil {
			log.Fatal(err)
		}

		if guess < traget {
			fmt.Println("too low")
		} else if guess > traget {
			fmt.Println("too high")
		} else {
			success = true
			fmt.Println("guess it")
			break
		}
	}
	if !success {
		fmt.Println("Sorry,too noob")
	}
}
