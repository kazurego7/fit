package util

import (
	"strings"
)

func SplitLn(str string) []string {
	res := strings.Split(strings.Trim(strings.ReplaceAll(string(str), `"`, ""), "\n"), "\n")
	if len(res) == 1 && res[0] == "" {
		return []string{}
	} else {
		return res
	}
}

func Difference(a, b []string) (diff []string) {
	m := make(map[string]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}
