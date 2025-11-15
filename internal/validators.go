package internal

import (
	"regexp"
)

func ValidateSizeStr(s string) bool {
	if s != "" {
		re := regexp.MustCompile(`^\d+[mkg]?$`)
		if !re.MatchString(s) {
			return false
		}
	}
	return true
}

func ValidateReturn(s string) bool {
	re := regexp.MustCompile(`^[1-5]{1,1}[0-9]{1,1}[0-9]{1,1} https?://[a-z./]+$`)
	if !re.MatchString(s) {
		return false
	}
	return true
}

func ValidateURL(s string) bool {
	re := regexp.MustCompile(`^https?://[a-z./]+$`)
	if !re.MatchString(s) {
		return false
	}
	return true
}
