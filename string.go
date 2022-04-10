package utils

import (
// "strings"
)

func InListIsExist(source string, check []string) bool {
	for _, s := range check {
		if source == s {
			return true
		}
	}
	return false
}
