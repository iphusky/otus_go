package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var taskWithAsteriskIsCompleted = true

var splitRegexp = regexp.MustCompile(`[^\s]+`)

var sizeOf = 10

type WordsCount struct {
	Str   string
	Count int
}

func Top10(text string) []string {
	if len(text) == 0 {
		return []string{}
	}

	dict := splitWordsByRegexp(text)

	resultList := make([]WordsCount, 0)

	for key, value := range dict {
		resultList = append(resultList, WordsCount{Str: key, Count: value})
	}

	return sortDataSet(resultList)
}

func splitWordsByRegexp(text string) map[string]int {
	dict := make(map[string]int)

	words := splitRegexp.FindAllString(text, -1)

	for _, value := range words {
		if taskWithAsteriskIsCompleted {
			value = strings.Trim(value, ",.!-")
			if len(value) == 0 {
				continue
			}
			value = strings.ToLower(value)
		}

		if val, ok := dict[value]; ok {
			dict[value] = val + 1
		} else {
			dict[value] = 1
		}
	}

	return dict
}

func sortDataSet(dataset []WordsCount) []string {
	sort.Slice(dataset, func(i, j int) bool {
		if dataset[i].Count == dataset[j].Count {
			return dataset[i].Str < dataset[j].Str
		}

		return dataset[i].Count > dataset[j].Count
	})

	var result []string

	for i := 0; i < len(dataset); i++ {
		if i >= sizeOf {
			break
		}
		result = append(result, dataset[i].Str)
	}

	return result
}
