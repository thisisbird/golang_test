package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	files, err := ioutil.ReadDir("src")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Directory:", file.Name())
		} else {
			fmt.Println("File:", file.Name())
		}
	}

	// recurses()
	count(1, 3)
	err = scanDirectory("src")
	if err!=nil{
		log.Fatal(err)
	}
}

func recurses() {
	fmt.Println("Oh,no,I'm stuck!")
	recurses()
}
func count(start int, end int) {
	fmt.Printf("count(%d,%d) called\n", start, end)
	fmt.Println(start)
	if start < end {
		count(start+1, end)
	}
	fmt.Printf("Returning from count(%d,%d) call\n", start, end)
}

func scanDirectory(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("%s發生錯誤\n",path)
		return err
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name()) //透過斜線把目錄路徑與檔案合併在一起
		if file.IsDir() {
			err := scanDirectory(filePath)
			if err != nil {
				fmt.Printf("%s發生錯誤2\n",path)
				return err
			}
		} else {
			fmt.Println(filePath)
		}
	}
	return nil
}
