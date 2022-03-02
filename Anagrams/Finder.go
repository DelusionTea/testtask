package Anagram

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
func Finder(text []string) []string {
	var resultArr []string
	//check dublicates if result is not empty, check @contains@
	parsedText := make(map[string][]rune)
	for _, word := range text {
		parsedText[word] = quickSort([]rune(word))
	}
	for key, value := range parsedText {
		if len(resultArr) != 0 {
			//we can add only new one
		}

	}
	//1.check size
	//2.v1 sort? or compare
	//2.v2 sort, then use ==  OR check each rune
	return resultArr
}
