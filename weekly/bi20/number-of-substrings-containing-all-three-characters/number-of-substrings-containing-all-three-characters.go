// Package number_of_substrings_containing_all_three_characters - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* number-of-substrings-containing-all-three-characters - file brief introduce */
/*
modification history
-----------------------------
2020/2/22 10:57 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package number_of_substrings_containing_all_three_characters

import (
	"strings"
)

func numberOfSubstrings(s string) int {
	if len(s) < 3 {
		return 0
	}
	if s[len(s)-1] == 'c' && !strings.Contains(s[:len(s)-1], "c") {
		return len(s) - 2
	}
	res := 0
	for i := 0; i < len(s)-2; i++ {
		k := map[uint8]bool{}
		k[s[i]] = true
		k[s[i+1]] = true
		for j := i + 1; j < len(s); j++ {
			k[s[j]] = true
			if len(k) == 3 {
				res += len(s) - j
				break
			}
		}
	}
	return res
}
