package utils

import (
	"strings"
)

// 截图字符串前N个字符内容
func GetStringRuneN(textStr string, n int) (retText string) {
	if len(textStr) <= 0 {
		return ""
	}
	retText = textStr
	textR := []rune(retText)
	if len(textR) > n {
		retText = string(textR[:n-1])
	}
	retText = strings.TrimSpace(retText)
	return retText
}
