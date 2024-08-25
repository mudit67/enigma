package utils

import "fmt"

func StringToRune(str string) (rune, error) {
	if len(str) != 1 {
		return 0, fmt.Errorf("string must have exactly one character")
	}
	return []rune(str)[0], nil
}
