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
	//var resultMap
	if len(text) <= 1 {
		return []string{""}, errors.New("not enough items. <2")
	}
	parsedText := make(map[string][]string)
	for _, word := range text {
		sortedWord := string(quickSort([]rune(word)))
		parsedText[sortedWord] = append(parsedText[sortedWord], word)

		// if len(parsedText[sortedWord]) == 2 {
		// 	resultArr = append(resultArr, parsedText[sortedWord]...)
		// }
		// if len(parsedText[sortedWord]) > 2 {
		// 	resultArr = append(resultArr, word)
		//}
	}
	for _, words := range parsedText {
		if len(words) >= 2 {
			resultArr = append(resultArr, words...)
		}
	}

	if len(resultArr) == 0 {
		return []string{""}, errors.New("Can't Find Anagrams")
	}

	return resultArr, nil
}
