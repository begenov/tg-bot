package services

import (
	"regexp"
	"strings"
)

type Validater struct {
}

func (u Validater) IsNameValid(name string) bool {
	parts := strings.Split(name, " ")
	if len(parts) == 1 {
		pattern := "^[A-Za-zА-Яа-я]+$"
		regex := regexp.MustCompile(pattern)
		// Check if the name matches the pattern
		return regex.MatchString(name)
	} else if len(parts) == 2 {
		pattern := "^[A-Za-zА-Яа-я]+ [A-Za-zА-Яа-я]+$"
		regex := regexp.MustCompile(pattern)
		// Check if the name and surname match the pattern
		return regex.MatchString(name)
	}
	return false
}
