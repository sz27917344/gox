package xstring

import "strings"

func IsBlank(text string) bool {
	return len(strings.Trim(text, " ")) == 0
}

func IsNotBlank(text string) bool {
	return !IsBlank(text)
}
