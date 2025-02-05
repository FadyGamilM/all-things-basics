package main

import (
	"strconv"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	out := [][]string{}
	tracker := map[string][]string{}
	for _, str := range strs {
		anchor := 'a'
		freq := make([]int, 26)
		freqStr := make([]string, 26)
		var i = 25
		for i >= 0 {
			freq[i] = 0
			i--
		}

		for _, char := range str {
			freq[char-anchor] += 1
		}

		for i, item := range freq {
			freqStr[i] = strconv.Itoa(item)
		}

		freqKey := strings.Join(freqStr, "#")
		if v, exists := tracker[freqKey]; exists {
			v = append(v, str)
			tracker[freqKey] = v
		} else {
			tracker[freqKey] = []string{str}
		}
	}

	for _, v := range tracker {
		out = append(out, v)
	}

	return out
}

func main() {
	groupAnagrams([]string{"bdddddddddd", "bbbbbbbbbbc"})
}
