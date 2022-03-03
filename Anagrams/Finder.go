package Anagram

import "errors"

func sort(arr []rune, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if arr[i] < pivot {
			temp := arr[splitIndex]
			arr[splitIndex] = arr[i]
			arr[i] = temp
			splitIndex++
		}
	}

	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	sort(arr, start, splitIndex-1)
	sort(arr, splitIndex+1, end)
}

func quickSort(arr []rune) []rune {

	newArr := make([]rune, len(arr))
	for i, v := range arr {
		newArr[i] = v
	}
	sort(newArr, 0, len(arr)-1)
	return newArr
}
func finder(text []string) ([]string, error) {
	var resultArr []string
	if len(text) <= 1 {
		return []string{""}, errors.New("not enough items. <2")
	}
	parsedText := make(map[string][]rune)
	for _, word := range text {
		parsedText[word] = quickSort([]rune(word))
	}
	var keys []string
	for key := range parsedText {
		if len(resultArr) != 0 {
			//we can add only new one
		}
		keys = append(keys, key)
	}
	for i, j := 0, 1; j < len(keys); i, j = i+1, j+1 {
		if len(keys[i]) == len(keys[j]) {
			if string(parsedText[keys[i]]) == string(parsedText[keys[j]]) {
				//log.Println(len(resultArr))
				//if len(resultArr) == 0 {
				resultArr = append(resultArr, keys[i])
				//}
				resultArr = append(resultArr, keys[j])
			}
		}
	}

	if len(resultArr) == 0 {
		return []string{""}, errors.New("Can't Find Anagrams")
	}

	return resultArr, nil
}
