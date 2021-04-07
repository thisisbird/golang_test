package main
import "fmt"
const DaysInWeek = 7

func WeeksToDays(weeks int) int {
	return weeks * DaysInWeek
}

func DaysToWeeks(days int) float64 {
	return float64(days) / float64(DaysInWeek)
}

func main() {
	a := WeeksToDays(2)
	b := DaysToWeeks(3)
	fmt.Println(a)
	fmt.Println(b)
}
