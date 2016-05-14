package problems

import "time"


/**
You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
 */


func Problem019() int {

	date, _ := time.Parse("2006-1-2", "1901-01-01")
	endDate, _ := time.Parse("2006-1-2", "2000-12-31")

	dayCounter := 0
	for date.Before(endDate) {
		if date.Weekday() == time.Sunday{
			dayCounter++
		}
		date = date.AddDate(0,1,0)
	}

	return dayCounter
}