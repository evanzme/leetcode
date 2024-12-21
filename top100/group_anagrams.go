package top100

import (
	"sort"
	"strings"
)

func groupAnagrams(str []string) [][]string {
	var res [][]string
	sortedMap := make(map[string][]string)
	for _, s := range str {
		sortedStrArr := strings.Split(s, "")
		sort.Strings(sortedStrArr)
		sortedStr := strings.Join(sortedStrArr, "")
		sortedMap[sortedStr] = append(sortedMap[sortedStr], s)
	}
	for _, strArr := range sortedMap {
		res = append(res, strArr)
	}
	return res
}
