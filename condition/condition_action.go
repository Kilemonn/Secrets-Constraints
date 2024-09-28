package condition

import (
	"strconv"
	"strings"
)

type ConditionAction interface {
	CheckCondition(input string, args []string) bool
}

type HasPrefixConditionAction struct{}

func (a HasPrefixConditionAction) CheckCondition(input string, args []string) bool {
	return strings.HasPrefix(input, args[0])
}

type HasSuffixConditionAction struct{}

func (a HasSuffixConditionAction) CheckCondition(input string, args []string) bool {
	return strings.HasSuffix(input, args[0])
}

type InvalidConditionAction struct{}

func (a InvalidConditionAction) CheckCondition(input string, args []string) bool {
	return false
}

type UniqueConditionAction struct {
	seenValues map[string]bool
}

func NewUniqueConditionAction() (a UniqueConditionAction) {
	a.seenValues = make(map[string]bool)
	return
}

func (a UniqueConditionAction) CheckCondition(input string, args []string) bool {
	if _, exists := a.seenValues[input]; exists {
		return false
	} else {
		// TODO: Set in the credential name as the value so we can log it if there is a collision
		a.seenValues[input] = true
		return true
	}
}

type IsNumericConditionAction struct{}

func (a IsNumericConditionAction) CheckCondition(input string, args []string) bool {
	_, err := strconv.ParseInt(input, 10, 64)
	return err == nil
}

type IsBooleanConditionAction struct{}

func (a IsBooleanConditionAction) CheckCondition(input string, args []string) bool {
	_, err := strconv.ParseBool(input)
	return err == nil
}
