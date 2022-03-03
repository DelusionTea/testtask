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
func getMapKey(m map[string]int, value int) (key string, result bool) {
	for k, v := range m {
		if v == value {
			key = k
			result = true
			return
		}
	}
	return
}
func finder(text []string) ([]string, error) {
	var resultArr []string
	if len(text) <= 1 {
		return []string{""}, errors.New("not enough items. <2")
	}
	parsedText := make(map[string]string)
	for _, word := range text {
		sortedWord := string(quickSort([]rune(word)))
		if parsedText[sortedWord] != "" {
			resultArr = append(resultArr, parsedText[sortedWord])
			resultArr = append(resultArr, word)
		}
		parsedText[sortedWord] = word
	}

	if len(resultArr) == 0 {
		return []string{""}, errors.New("Can't Find Anagrams")
	}

	return resultArr, nil
}
