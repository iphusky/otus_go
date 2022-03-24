package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(changedString string) (string, error) {

	if len(changedString) == 0 {
		return "", nil
	}

	if isStringCorrect(changedString) == false {
		return "", ErrInvalidString
	}

	var result strings.Builder

	runes := []rune(changedString)

	for i, str := range runes {

		if unicode.IsDigit(str) {

			repeatCount, _ := strconv.Atoi(string(str))

			if repeatCount == 0 {

				newString, _ := removeLastElement(result)
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

	matched, _ := regexp.MatchString(`\d{2,}`, checkString)

	if matched == true {
		return false
	}

	return true
}

func removeLastElement(buildString strings.Builder) (string, error) {

	str := buildString.String()
	if len(str) > 0 {
		return str[:len(str)-1], nil
	}

	return "", nil
}
