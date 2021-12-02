package utils

import "strings"

func GetOptions(options string) []string {
	return strings.Split(options, ",")
}

func Pick(options []string, randomNum int) string {
	return options[randomNum]
}
