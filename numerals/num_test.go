package numerals

import "testing"

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
