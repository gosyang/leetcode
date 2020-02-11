// Package _75 - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* tweet-counts-per-frequency - file brief introduce */
/*
modification history
-----------------------------
2020/2/9 11:15 AM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package _175

type TweetCounts struct {
	record map[string]map[int]int
}


func Constructor() TweetCounts {
	return TweetCounts{
		recordMinute: make(map[string]map[int]int),
		recordHour: make(map[string]map[int]int),
		recordDay: make(map[string]map[int]int),
	}
}


func (this *TweetCounts) RecordTweet(tweetName string, time int)  {
	if _, ok := this.recordMinute[tweetName]; !ok {
		this.recordMinute[tweetName] = make(map[int]int)
		this.recordHour[tweetName] = make(map[int]int)
		this.recordDay[tweetName] = make(map[int]int)
	}
	this.recordMinute[tweetName][time/60] += 1
	this.recordHour[tweetName][time/3600] += 1
	this.recordDay[tweetName][time/86400] += 1
}


func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	res := make([]int, 0)
	switch freq {
	case "minute":
		for i := startTime/60; i <= endTime/60; i++ {
			res = append(res, this.recordMinute[tweetName][i])
		}
	case "hour":
		for i := startTime/3600; i <= endTime/3600; i++ {
			res = append(res, this.recordHour[tweetName][i])
		}
	case "day":
		for i := startTime/86400; i <= endTime/86400; i++ {
			res = append(res, this.recordDay[tweetName][i])
		}
	}
	return res
}

// 一开始按照小时存，虽然简单但是不对，需要具体到秒

/**
 * Your TweetCounts object will be instantiated and called as such:
 * obj := Constructor();
 * obj.RecordTweet(tweetName,time);
 * param_2 := obj.GetTweetCountsPerFrequency(freq,tweetName,startTime,endTime);
 */