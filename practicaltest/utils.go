package practicaltest

import (
	"log"
	"time"
)

func parsetime(timeStr1, timestr2 string) int {
	duration1, err := time.Parse("2006-01-02T15:04:05.000Z", timeStr1)
	if err != nil {
		log.Println("error when parsing time1: ", err.Error())
	}
	duration2, err := time.Parse("2006-01-02T15:04:05.000Z", timestr2)
	if err != nil {
		log.Println("error when parsing time2: ", err.Error())
	}
	duration := duration2.Hour() - duration1.Hour()
	return duration
}
