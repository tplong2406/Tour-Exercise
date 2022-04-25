package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
 * Complete the 'ModifyString' function below and add imports if needed.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING str as parameter.
 */

func ModifyString(str string) string {

	temp := []rune(str)
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	for k := 0; k < len(temp); k++ {
		if temp[k] > 47 && temp[k] < 58 {
			temp[k] = 32
		}
	}
	fmt.Println(temp)

	s := strings.TrimFunc(string(temp), func(r rune) bool {
		return unicode.IsDigit(r)
	})

	return s
}

func main() {
	kQ := ModifyString("    o65ll321eH   ")
	fmt.Println(kQ)
}
