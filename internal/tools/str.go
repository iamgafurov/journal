package tools

import "strings"

func StrEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}
