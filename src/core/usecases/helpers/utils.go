package helpers

import (
	"regexp"
	"time"
)

func RemoveSpecialCharacters(str string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9_-]+")

	processedString := reg.ReplaceAllString(str, "")

	return processedString, err
}

func GetDateTimeStr() string {
	var str string

	t := time.Now()

	str = t.Format("20060102150405")

	return str
}
