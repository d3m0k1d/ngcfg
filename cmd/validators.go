package cmd

import (
	"regexp"
)

func ValidateSizeStr(s string) bool {
	re := regexp.MustCompile(`^\d+[mkg]?$`)
	if !re.MatchString(s) {
		panic("invalid size please check the documentation")
	}
	return true
}
