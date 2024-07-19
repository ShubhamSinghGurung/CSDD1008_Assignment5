package main

import (
	"fmt"
)

func IntToRoman(num int) string {

	valueSymbols := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	result := ""
	for _, vs := range valueSymbols {
		for num >= vs.value {
			result += vs.symbol
			num -= vs.value
		}
	}

	return result
}

func main() {

	nums := []int{1, 4, 9, 58, 1994}
	for _, num := range nums {
		fmt.Printf("%d in Roman Numerals is %s\n", num, IntToRoman(num))
	}
}
