package FindNum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

//simple version
func findMax(arr []string) int {
	max, _ := strconv.Atoi(arr[0])
	for _, nums := range arr {
		value, _ := strconv.Atoi(nums)
		if value > max {
			max = value
		}
	}
	return max
}
func finder(text string) (int, error) {
	if text == "" {
		return 0, errors.New("empty string")
	}
	//var resultArr []string
	//var resultMap

	//only positive numbers
	//nums := regexp.MustCompile("[0-9]+")
	//with negative numbers
	nums := regexp.MustCompile("-?[0-9]+")
	resultArr := nums.FindAllString(text, -1)
	if len(resultArr) == 0 {
		return 0, errors.New("Contains no numbers")
	}
	result := findMax(resultArr)

	fmt.Println(nums.FindAllString(text, -1))
	fmt.Println(result)
	return result, nil
}
