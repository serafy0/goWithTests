package numerals

import (
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {

	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 returns I", 1, "I"},
		{"2 returns I", 2, "II"},
		{"3 returns I", 3, "III"},
		{"4 returns IV", 4, "IV"},
		{"5 returns V", 5, "V"},
		{"6 returns VI", 6, "VI"},
		{"7 returns VII", 7, "VII"},
		{"8 returns VIII", 8, "VIII"},
		{"9 returns IX", 9, "IX"},
		{"10 returns X", 10, "X"},
		{"14 returns XIV", 14, "XIV"},
		{"18 returns XVIII", 18, "XVIII"},
		{"20 returns XX", 20, "XX"},
		{"39 returns XXXIX", 39, "XXXIX"},
		{"40 returns XL", 40, "XL"},
		{"47 returns XLVII", 47, "XLVII"},
		{"49 returns XLIX", 49, "XLIX"},
		{"50 returns L", 50, "L"},
		{"100", 100, "C"},
		{"90", 90, "XC"},
		{"400", 400, "CD"},
		{"500", 500, "D"},
		{"900", 900, "CM"},
		{"1000", 1000, "M"},
		{"1984", 1984, "MCMLXXXIV"},
		{"3999", 3999, "MMMCMXCIX"},
		{"2014", 2014, "MMXIV"},
		{"1006", 1006, "MVI"},
		{"798", 798, "DCCXCVIII"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {

			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}

}
func TestConvertingToArabic(t *testing.T) {

	cases := []struct {
		Description string
		Arabic      int
		Roman       string
	}{
		{"1 returns I", 1, "I"},
		{"2 returns I", 2, "II"},
		{"3 returns I", 3, "III"},
		{"4 returns IV", 4, "IV"},
		{"5 returns V", 5, "V"},
		{"6 returns VI", 6, "VI"},
		{"7 returns VII", 7, "VII"},
		{"8 returns VIII", 8, "VIII"},
		{"9 returns IX", 9, "IX"},
		{"10 returns X", 10, "X"},
		{"14 returns XIV", 14, "XIV"},
		{"18 returns XVIII", 18, "XVIII"},
		{"20 returns XX", 20, "XX"},
		{"39 returns XXXIX", 39, "XXXIX"},
		{"40 returns XL", 40, "XL"},
		{"47 returns XLVII", 47, "XLVII"},
		{"49 returns XLIX", 49, "XLIX"},
		{"50 returns L", 50, "L"},
		{"100", 100, "C"},
		{"90", 90, "XC"},
		{"400", 400, "CD"},
		{"500", 500, "D"},
		{"900", 900, "CM"},
		{"1000", 1000, "M"},
		{"1984", 1984, "MCMLXXXIV"},
		{"3999", 3999, "MMMCMXCIX"},
		{"2014", 2014, "MMXIV"},
		{"1006", 1006, "MVI"},
		{"798", 798, "DCCXCVIII"},
	}
	for _, test := range cases[:1] {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}
