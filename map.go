package main
import "fmt"
func main() {
	monthdays := map[string]int {
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31, // 逗号
	}
	year := 0
	for _, v := range monthdays {
		year += v
	}
	fmt.Printf("Year days: %d\n", year)
	fmt.Printf("days in Dec: %d\n", monthdays["Dec"])
	monthdays["Feb"] = 29
	year = 0
	for _, v := range monthdays {
		year += v
	}
	fmt.Printf("Year days: %d\n", year)
}
