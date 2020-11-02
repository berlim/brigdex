package helpers

import (
	"regexp"
)

func RemoveSpecialCharacters(str string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9_-]+")

	processedString := reg.ReplaceAllString(str, "")

	return processedString, err
}
