package main

import (
	"errors"
)

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
	if arabNum > 3999 || arabNum < 1 {
		return "", errors.New("non-convertible number. >3999 || < 1")
	}

	//Fun fact: in Golang iterations over maps random
	keys := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for i := 0; i < len(keys); i++ {
		if arabNum >= keys[i] {
			romanResult += ArabRoman[keys[i]]
			arabNum -= keys[i]
		}
	}

	return romanResult, nil
}
