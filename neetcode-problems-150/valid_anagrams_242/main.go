package main

import "log"

func main() {
	s := "abc"
	log.Println(s[0] - byte('a'))
	log.Println(s[1] - byte('a'))
	log.Println(s[2] - byte('a'))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s_records, t_records := map[byte]int{}, map[byte]int{}

	for i := range s {
		if _, exists := s_records[s[i]]; exists {
			s_records[s[i]]++
		} else {
			s_records[s[i]] = 1
		}

		if _, exists := t_records[t[i]]; exists {
			t_records[t[i]]++
		} else {
			t_records[t[i]] = 1
		}
	}

	for s_key, s_val := range s_records {
		t_val, exists := t_records[s_key]
		if !exists {
			return false
		}

		if t_val != s_val {
			return false
		}
	}

	return true
}

// since we already know the cap of the characters we want to compare to each others, we can define a slots slice that store the frequency of each char of the 26 english character
func isAnagram_optimized(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	slots := make([]int, 26)
	anchor := byte('a')

	for i := range s {
		s_charSlot := s[i] - anchor
		t_charSlot := t[i] - anchor
		slots[s_charSlot]++
		slots[t_charSlot]--
	}

	for i := range slots {
		if slots[i] != 0 {
			return false
		}
	}

	return true
}
