package utils

import (
	"strconv"
	"time"
)

func FormatInteger(num int) string {
	numStr := strconv.Itoa(num)
	length := len(numStr)

	dots := (length - 1) / 3

	formattedStr := ""
	for i := 0; i < dots; i++ {
		formattedStr = "." + numStr[length-3*(i+1):length-3*i] + formattedStr
	}
	formattedStr = numStr[:length-3*dots] + formattedStr

	return formattedStr
}

func MillisecondsToTime(ms int) string {
	duration := time.Duration(ms) * time.Millisecond

	roundedDuration := duration.Round(time.Second)

	t := time.Time{}.Add(roundedDuration)

	formattedTime := t.Format("15:04:05")

	return formattedTime
}
