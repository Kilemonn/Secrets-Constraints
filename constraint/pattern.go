package constraint

import (
	"fmt"
	"regexp"
)

const (
	all_pattern = "ALL"
)

type Pattern struct {
	pattern string
	regex   *regexp.Regexp
}

func NewPattern(p string) (pattern Pattern, err error) {
	pattern.pattern = p
	regex, err := regexp.Compile(pattern.pattern)
	if err != nil {
		fmt.Printf("Failed to compile regex pattern: [%s].\n", err.Error())
	} else {
		pattern.regex = regex
	}
	return
}

func (p Pattern) Matches(input string) bool {
	if p.pattern == all_pattern {
		return true
	}
	if p.regex == nil {
		return false
	}

	return p.regex.Match([]byte(input))
}
