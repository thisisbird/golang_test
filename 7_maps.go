package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	lines, err := GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	fmt.Println(counts)
	for name,value := range counts {
		fmt.Println(name,value) //counts 隨機跑的 沒有順序
	}

	grades := map[string]float64{"Edward":99.9,"Alma":72.2,"Carl":55.2}
	var names []string
	for name := range grades{
		names = append(names,name)
	}
	sort.Strings(names) //排序後 按照順序 不會隨機
	for _,name := range names{
		fmt.Printf("%s has a grade of %0.1f%%\n",name,grades[name])
	}

	ranks := map[string]int{"bronze": 3, "silver": 2, "gold": 1}
	fmt.Println(ranks["gold"])

	elements := map[string]string{
		"H":  "Hydrogen",
		"Li": "Lithium",
	}
	fmt.Println(elements["H"])

	emptyMap := map[string]float64{}
	fmt.Println(emptyMap, elements)

	numbers := make(map[string]int)
	numbers["qqq aaa"] = 12
	fmt.Printf("%#v\n", numbers["qqq aaa"])
	fmt.Printf("%#v\n", numbers["qqq aaa ccc"])

	numbers2 := make(map[int]string)
	numbers2[3] = "three"
	fmt.Printf("%#v\n", numbers2[3])

	status("Alma")
	status("qqqq")
	deleteFunc()
}

func GetStrings(filename string) ([]string, error) {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	// i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}

func status(name string) {
	grades := map[string]float64{"Alma": 0, "Rohit": 86.5}
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("NO grade recorded for %s.\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing!\n", name)
	}
}

func deleteFunc() {
	var ok bool
	ranks := make(map[string]int)
	var rank int
	ranks["bronze"] = 3
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d,ok: %v\n", rank, ok)

	delete(ranks, "bronze")
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d,ok: %v\n", rank, ok)

}
