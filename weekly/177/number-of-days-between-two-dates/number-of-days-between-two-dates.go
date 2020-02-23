// Package number_of_days_between_two_dates - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* number-of-days-between-two-dates - file brief introduce */
/*
modification history
-----------------------------
2020/2/23 10:31 AM, by gosyang, create
*/
/*
DESCRIPTION
请你编写一个程序来计算两个日期之间隔了多少天。

日期以字符串形式给出，格式为 YYYY-MM-DD
*/
package number_of_days_between_two_dates

import (
	"time"
)

func daysBetweenDates(date1 string, date2 string) int {
	day1, _ := time.Parse("2006-01-02 15:04:05", date1+" 15:04:05")
	day2, _ := time.Parse("2006-01-02 15:04:05", date2+" 15:04:05")
	diff := (day1.Unix() - day2.Unix()) / 3600 / 24
	if diff < 0 {
		diff = 0 - diff
		return int(diff)
	}
	return int(diff)
}
