package propertybased

import (
	"strings"
)

type number struct {
	arabic int
	roman  string
}

var numbers = []number{
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

func ConvertToRoman(arabic int) string {
	var romanNumeral strings.Builder
	for _, n := range numbers {
		if arabic >= n.arabic {
			romanNumeral.WriteString(strings.Repeat(n.roman, arabic/n.arabic))
			arabic = arabic % n.arabic
		}
	}
	return romanNumeral.String()
}

var numeralToArabic = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func ConvertToArabic(romanNumeral string) int {
	var arabic int
	for i := 0; i < len(romanNumeral); i += 2 {
		if i+2 > len(romanNumeral) {
			arabic += numeralToArabic[string(romanNumeral[i])]
			return arabic
		}

		switch romanNumeral[i : i+2] {
		case "IV":
			arabic += 4
			continue
		case "IX":
			arabic += 9
			continue
		case "XL":
			arabic += 40
			continue
		case "XC":
			arabic += 90
			continue
		case "CD":
			arabic += 400
			continue
		case "CM":
			arabic += 900
			continue
		}

		arabic += numeralToArabic[string(romanNumeral[i])]
		i--
	}
	return arabic
}
