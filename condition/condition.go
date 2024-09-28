package condition

import (
	"fmt"
	"slices"
	"strings"
)

type ConditionType uint

const (
	ConditionTypeInvalid   ConditionType = iota
	ConditionTypeUnique    ConditionType = iota
	ConditionTypeHasPrefix ConditionType = iota
	ConditionTypeHasSuffix ConditionType = iota
	ConditionTypeIsNumeric ConditionType = iota
	ConditionTypeIsBoolean ConditionType = iota
)

var (
	// TODO: Create a new instance for each different constraint definition that uses Unique
	// Need to create and hold this variable, since it's state needs to be retained
	uniqueConditionAction UniqueConditionAction = NewUniqueConditionAction()
)

func ConditionTypeFromString(input string) ConditionType {
	index := slices.Index(conditionTypeStrings(), strings.ToLower(input))
	if index == -1 {
		return ConditionTypeInvalid
	} else {
		return ConditionType(index)
	}
}

func conditionTypeStrings() []string {
	return []string{"invalid", "unique", "hasprefix", "hassuffix", "isnumeric", "isboolean"}
}

func (c ConditionType) IsValid() bool {
	return c != ConditionTypeInvalid
}

type Condition struct {
	Type ConditionType
	Args []string
}

func (c ConditionType) expectedArgsCount() uint {
	args := []uint{0, 0, 1, 1, 0, 0}
	return args[uint(c)]
}

func (c ConditionType) conditionActions() []ConditionAction {
	return []ConditionAction{
		InvalidConditionAction{},
		uniqueConditionAction,
		HasPrefixConditionAction{},
		HasSuffixConditionAction{},
		IsNumericConditionAction{},
		IsBooleanConditionAction{},
	}
}

func (c ConditionType) getConditionAction() ConditionAction {
	return c.conditionActions()[uint(c)]
}

func NewCondition(conditionString string) (condition Condition, err error) {
	index := strings.Index(conditionString, "(")
	conditionTypeString := conditionString
	if index != -1 {
		conditionTypeString = conditionString[0:index]
	}
	conditionType := ConditionTypeFromString(conditionTypeString)
	if !conditionType.IsValid() {
		err = fmt.Errorf("provided condition type [%s] is invalid", conditionTypeString)
		return
	}
	argumentsString := ""
	if index != -1 {
		argumentsString = conditionString[index+1:]

		closingBracketIndex := strings.LastIndex(argumentsString, ")")
		if closingBracketIndex != -1 {
			argumentsString = argumentsString[:len(argumentsString)-1]
		} else {
			err = fmt.Errorf("no closing bracket provided in condition [%s]", conditionString)
			return
		}
	}

	args, err := getArguments(argumentsString, conditionType.expectedArgsCount())
	if err != nil {
		fmt.Printf("Failed to create condition. Error: [%s]\n", err.Error())
		return
	}

	condition.Type = conditionType
	condition.Args = args

	return
}

func getArguments(arg string, expectedArgsCount uint) (args []string, err error) {
	split := strings.Split(arg, ",")
	// Remove empty
	split = slices.DeleteFunc(split, func(e string) bool { return e == "" })
	if len(split) > int(expectedArgsCount) {
		fmt.Printf("Too many arguments provided, expected [%d] but found [%d]. Only the first [%d] will be used.\n", expectedArgsCount, len(split), expectedArgsCount)
	} else if len(split) < int(expectedArgsCount) {
		err = fmt.Errorf("not enough arguments supplied for condition type")
		return
	}
	for i := range expectedArgsCount {
		args = append(args, strings.TrimSpace(split[i]))
	}
	return
}

func (c Condition) ApplyCondition(input string) bool {
	return c.Type.getConditionAction().CheckCondition(input, c.Args)
}
