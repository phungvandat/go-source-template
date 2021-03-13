package helper

import "strings"

// GetStringInBetween func
func GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return str[s:]
	}
	e += s
	return str[s:e]
}
