// Package _75 - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* minimum-number-of-steps-to-make-two-strings-anagram - file brief introduce */
/*
modification history
-----------------------------
2020/2/9 10:45 AM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _175

func minSteps(s string, t string) int {
	ms := make(map[uint8]int)
	mt := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		if _, ok := ms[s[i]]; !ok {
			ms[s[i]] = 0
		}
		ms[s[i]] += 1
		if _, ok := mt[t[i]]; !ok {
			mt[t[i]] = 0
		}
		mt[t[i]] += 1
	}
	res := len(s)
	for k, v := range ms {
		if tv, ok := mt[k]; ok {
			if v <= tv {
				res -= v
			} else {
				res -= tv
			}
		}
	}
	return res
}
