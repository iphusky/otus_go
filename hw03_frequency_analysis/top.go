package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var splitRegexp = regexp.MustCompile(`[^\\s]+`)

var sizeOf = 10

type WordsCount struct {
	Str   string
	Count int
}

func Top10(text string) []string {
	dict := splitWordsByRegexp(text)

	resultList := make([]WordsCount, sizeOf, sizeOf*2)

	for key, value := range dict {
		resultList = append(resultList, WordsCount{Str: key, Count: value})
	}

	sort.Slice(resultList, func(i, j int) bool {
		if resultList[i].Count == resultList[j].Count {
			return resultList[i].Str < resultList[j].Str
		}

		return resultList[i].Count > resultList[j].Count
	})

	var result []string

	for i := 0; i < len(resultList); i++ {
		result = append(result, resultList[i].Str)
		if i >= sizeOf {
			break
		}
	}

	return result
}

func splitWordsByRegexp(text string) map[string]int {
	dict := make(map[string]int)

	words := splitRegexp.FindAllString(text, -1)

	for _, value := range words {
		value = strings.Trim(value, ",.!")

		if val, ok := dict[value]; ok {
			dict[value] = val + 1
		} else {
			dict[value] = 1
		}
	}

	return dict
}
