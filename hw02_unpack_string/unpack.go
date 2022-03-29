package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

var isNeeded = regexp.MustCompile(`\d{2,}`)

func Unpack(changedString string) (string, error) {
	if len(changedString) == 0 {
		return "", nil
	}

	if !isStringCorrect(changedString) {
		return "", ErrInvalidString
	}

	var result strings.Builder

	runes := []rune(changedString)

	for i, str := range runes {
		if unicode.IsDigit(str) {
			repeatCount, err := strconv.Atoi(string(str))
			if err != nil {
				return "", err
			}
			if repeatCount == 0 {
				newString := removeLastElement(result.String())
				result.Reset()
				result.WriteString(newString)
			} else {
				result.WriteString(strings.Repeat(string(runes[i-1]), repeatCount-1))
			}
		} else {
			result.WriteString(string(str))
		}
	}

	return result.String(), nil
}

func isStringCorrect(checkString string) bool {
	if unicode.IsDigit(rune(checkString[0])) {
		return false
	}

	matched := isNeeded.MatchString(checkString)

	return !matched
}

func removeLastElement(str string) string {
	runes := []rune(str)

	if last := len(runes) - 1; last >= 0 {
		str = string(runes[:last])
	}
	return str
}
