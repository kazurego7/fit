package util

import "strings"

func SplitLn(str string) []string {
	return strings.Split(strings.Trim(strings.ReplaceAll(string(str), `"`, ""), "\n"), "\n")
}
