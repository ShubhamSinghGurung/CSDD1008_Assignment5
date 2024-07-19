package main

import "testing"

// TestIntToRoman tests the IntToRoman function with various cases.
func TestIntToRoman(t *testing.T) {
	testCases := []struct {
		num      int
		expected string
	}{
		{1, "I"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
		{3999, "MMMCMXCIX"},
		{2024, "MMXXIV"},
		{44, "XLIV"},
		{76, "LXXVI"},
		{499, "CDXCIX"},
	}

	for _, tc := range testCases {
		result := IntToRoman(tc.num)
		if result != tc.expected {
			t.Errorf("IntToRoman(%d) = %s; expected %s", tc.num, result, tc.expected)
		}
	}
}
