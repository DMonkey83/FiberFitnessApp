package util

import (
	"regexp"
)

// credits to this -> https://ihateregex.io/expr/

func MatchesRegex(regex string, value string) bool {
	isValid := regexp.MustCompile(regex).MatchString(value)
	return isValid
}
