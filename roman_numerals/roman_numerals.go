package romannumerals

import "strings"

func RomanToArabic(roman string) (total uint16) {
	for _, sym := range romanString(roman).Symbols() {
		total += RomanNumeralsTable.ValueOf(sym...)
	}
	return
}

type romanNumeral struct {
	Val    uint16
	Symbol string
}

type romanNumerals []romanNumeral

var RomanNumeralsTable = romanNumerals{
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

type romanString string

func (r romanString) Symbols() (syms [][]byte) {
	for i := 0; i < len(r); i++ {
		sym := r[i]
		notEnd := i+1 < len(r)

		if notEnd && isSubtractive(sym) && RomanNumeralsTable.Exists(sym, r[i+1]) {
			syms = append(syms, []byte{sym, r[i+1]})
			i++
		} else {
			syms = append(syms, []byte{sym})
		}
	}
	return
}

func (r romanNumerals) Exists(syms ...byte) bool {
	sym := string(syms)
	for _, s := range r {
		if s.Symbol == sym {
			return true
		}
	}
	return false
}

func isSubtractive(sym uint8) bool {
	return sym == 'I' || sym == 'X' || sym == 'C'
}

func (r romanNumerals) ValueOf(syms ...byte) uint16 {
	sym := string(syms)
	for _, s := range r {
		if s.Symbol == sym {
			return s.Val
		}
	}

	return 0
}

func ArabicToRoman(arabic uint16) string {
	var res strings.Builder

	for _, num := range RomanNumeralsTable {
		for arabic >= num.Val {
			res.WriteString(num.Symbol)
			arabic -= num.Val
		}
	}

	return res.String()
}
