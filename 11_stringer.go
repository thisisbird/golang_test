package main

import "fmt"

type Stringer interface {
	String() string
}
type CoffeePot string

type Gallons float64

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

type Liters float64

func (g Liters) String() string {
	return fmt.Sprintf("%0.2f L", g)
}

type Milliliters float64

func (g Milliliters) String() string {
	return fmt.Sprintf("%0.2f mL", g)
}

func (c CoffeePot) String() string {
	return string(c) + " coffee pot"
}
func main() {
	coffeePot := CoffeePot("LuxBrew")
	fmt.Println(coffeePot.String())
	fmt.Print(coffeePot, "\n")
	fmt.Println(coffeePot)
	fmt.Printf("%s\n", coffeePot)

	fmt.Println(Gallons(12.09248342))
	fmt.Println(Liters(12.09248342))
	fmt.Println(Milliliters(12.09248342))
}
