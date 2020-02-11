// Package _75 - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* check-if-n-and-its-double-exist - file brief introduce */
/*
modification history
-----------------------------
2020/2/9 10:31 AM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _175

func checkIfExist(arr []int) bool {
	hash := make(map[int]bool)
	zero := 0
	for _, i := range arr {
		if i == 0 {
			zero += 1
		} else {
			hash[2*i] = true
		}
	}
	if zero > 1 {
		return true
	}
	for _, i := range arr {
		if _, ok := hash[i]; ok {
			return true
		}
	}
	return false
}
