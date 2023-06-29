package handler

import (
	"log"
	"time"
)

func IsTimeFormat(str string) bool{
    _, err := time.Parse("2006-01-02 15:04", str)
	if err != nil {
		return false
	}
	return true
}
func StrToTime(str string) time.Time {
    t, err := time.Parse("2006-01-02 15:04", str)
	if err != nil {
		log.Panic(err)
	}
	return t
}