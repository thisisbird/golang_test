package calendar

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

type Event struct {
	title string //	不直接存取 用小寫
	Date
}
type Date struct {
	year  int //小寫不給別人存取
	month int
	day   int
}

type MyType struct {
	EmbeddedType
}

type EmbeddedType string

func (e EmbeddedType) ExportedMethod() {
	fmt.Println("Hi from ExportedMethond on EmbeddedType")
}

func (e EmbeddedType) unexportedMethod() {
	fmt.Println("Hi from unexportedMethond on EmbeddedType")
}

func (e Event) Title() string {
	return e.title
}

func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}
