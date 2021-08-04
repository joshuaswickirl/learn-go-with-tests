package propertybased_test

import (
	"fmt"
	"testing"
	"testing/quick"

	"github.com/joshuaswickirl/learn-go-with-tests/internal/propertybased"
)

var cases = []struct {
	arabic int
	roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{9, "IX"},
	{10, "X"},
	{40, "XL"},
	{50, "L"},
	{90, "XC"},
	{100, "C"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
}

func TestConvertingToRomanNumerals(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprintf("%d to %s", c.arabic, c.roman), func(t *testing.T) {
			got := propertybased.ConvertToRoman(c.arabic)
			if got != c.roman {
				t.Errorf("got %q, want %q", got, c.roman)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprintf("%s to %d", c.roman, c.arabic), func(t *testing.T) {
			got := propertybased.ConvertToArabic(c.roman)
			if got != c.arabic {
				t.Errorf("got %d, want %d", got, c.arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic int) bool {
		if arabic < 0 || arabic > 3999 {
			return true
		}
		roman := propertybased.ConvertToRoman(arabic)
		fromRoman := propertybased.ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
		t.Error("failed checks", err)
	}
}
