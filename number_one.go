package main

import (
	"fmt"
	"strings"
)

func NumberOne(n int, arr []string) string {
	var (
		result         = make(map[string][]string)
		lastMatchCount int
		lastMatchStr   string
	)

	for i := 0; i < n; i++ {
		//validate max array
		if i > len(arr)-1 {
			break
		}

		//make the string insensitive
		str := strings.ToLower(arr[i])

		//store all result
		result[str] = append(result[str], fmt.Sprintf("%v", i+1))

		//store the max result
		if len(result[str]) > lastMatchCount && len(result[str]) > 1 {
			lastMatchCount = len(result[str])
			lastMatchStr = str
		}
	}

	if lastMatchCount > 0 {
		return strings.Join(result[lastMatchStr], " ")
	}

	return "false"
}
