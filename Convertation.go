package main

import (
	"errors"
)

//Check < 4000 - otherwise error or warning?

//Romance rules:
//
//I=1 V=5 X=10 L=50 C=100 D=500  M=1000
//
//Example : XLVII = XL (40 = 50-10) + V (5)  +II(2)
//
//!!!!number worthwhile on the right(that is from which it is deducated)
//Should be no more than what is multiplied by ten!!!
//----example:
// 99!=IC   99 ==XCIX
// 49!-IL   49 == LXIX
//
//!!!It is imposible to repeat 4 digits in a row. That is the number 40 cannot be recorded as XXXX, only as LX.
// Of all these rules, it follows that the max number that can be recorded by Roman numbers is MMMCMXCIX = 3999

//need counter of digits in a row
//check num < 4 || 9
//options 1 11 111 5-1 5 5+1 5+2, 5+3, 10-1, 10 ....

//when we use subtraction
//1. I before V and X (5/10) 4 9
//2. X before L and C (50/100) 40 90
//3. C before D and M (500/1000) 400 900

//just because we can't create const map in go
//func getRomanValue() func(int) string{
//
//}
//func getArabValue func(string) int {
//
//}

func romanToArab(romanNum string) int {

	romanNumRune := []rune(romanNum)
	RomanArab := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	arabNum := 0

	for i, j := 0, 1; j < len(romanNum); i, j = i+1, j+1 {
		arabNum += RomanArab[romanNumRune[j]]
		if RomanArab[romanNumRune[j]] > RomanArab[romanNumRune[j]] {
			arabNum -= RomanArab[romanNumRune[j]] * 2
		}
	}
	return arabNum
}

func ArabToRoman(arabNum int) (string, error) {

	var romanResult string

	ArabRoman := map[int]string{

		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	if arabNum > 3999 {
		return "", errors.New("non-convertible number. >3999")
	}
	for key, value := range ArabRoman {
		if arabNum >= key {
			romanResult += value
			arabNum -= key
		}
	}

	return romanResult, nil
}

//}
