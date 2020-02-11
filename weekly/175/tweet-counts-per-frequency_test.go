// Package main - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* tweet-counts-per-frequency_test - file brief introduce */
/*
modification history
-----------------------------
2020/2/9 11:59 AM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _175

import "testing"

func TestTweet(t *testing.T) {
	obj := Constructor()
	obj.RecordTweet("tweet0", 99)
	obj.RecordTweet("tweet1", 80)
	obj.RecordTweet("tweet2", 29)
	obj.RecordTweet("tweet3", 81)
	obj.RecordTweet("tweet4", 56)
	t.Errorf("%v", obj.GetTweetCountsPerFrequency("day", "tweet4", 56, 2667))
}