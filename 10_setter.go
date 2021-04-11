package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	year  int //小寫不給別人存取
	month int
	day   int
}

func (d *Date) SetYear(year int) error { //使用指標*Date
	if year < 1 {
		return errors.New("invaild year")
	}
	d.year = year
	return nil
}
func (d *Date) SetMonth(month int) error { //使用指標*Date
	if month < 1 || month > 12 {
		return errors.New("invaild month")
	}
	d.month = month
	return nil
}
func (d *Date) SetDay(day int) error { //使用指標*Date
	if day < 1 || day > 31 {
		return errors.New("invaild day")
	}
	d.day = day
	return nil
}

func (d *Date) Year() int{ //go允許使用GetYear，但不必要這樣做，go社群習慣不加
	return d.year
}
func (d *Date) Month() int{
	return d.month
}
func (d *Date) Day() int{
	return d.day
}

func main() {
	date := Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year)
}
