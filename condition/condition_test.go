package condition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetArguments(t *testing.T) {
	var cases = []struct {
		argString     string
		expectedCount uint
		expectsError  bool
		expected      []string
	}{
		{"", 0, false, []string{}},                               // Empty args, non requested
		{"", 1, true, []string{}},                                // Empty args, 1 requested
		{"arg", 0, false, []string{}},                            // 1 arg, with 0 requested
		{"Sometest-1237-", 1, false, []string{"Sometest-1237-"}}, // One arg with no "," all should be returned
		{"point, test", 2, false, []string{"point", "test"}},     // 2 Args requested, with spaces
	}

	for _, c := range cases {
		result, err := getArguments(c.argString, c.expectedCount)
		if c.expectsError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, len(c.expected), len(result))
			for i := range c.expected {
				assert.Equal(t, c.expected[i], result[i])
			}
		}
	}
}

func TestNewCondition(t *testing.T) {
	var cases = []struct {
		conditionString string
		expectError     bool
		expectedType    ConditionType
		expectedArgs    []string
	}{
		{"Unique", false, ConditionTypeUnique, []string(nil)},        // 0 arg condition
		{"Unique(check)", false, ConditionTypeUnique, []string(nil)}, // 0 arg condition with args
		{"NotUnique", true, ConditionTypeUnique, []string(nil)},      // 0 arg but not matching any type

		{"HasPrefix", true, ConditionTypeHasPrefix, []string(nil)},                 // 1 arg required but not provided
		{"HasPrefix(test)", false, ConditionTypeHasPrefix, []string{"test"}},       // 1 arg required and provided
		{"HasPrefix(test, text)", false, ConditionTypeHasPrefix, []string{"test"}}, // 1 arg required and 2 provided

		{"HasPrefix( , text)", false, ConditionTypeHasPrefix, []string{""}}, // 1 arg required and 2 provided but first is a space

		{"HasPrefix(, text)", false, ConditionTypeHasPrefix, []string{"text"}}, // 1 arg required and 2 provided but first is empty string

		{"HasSuffix(test", true, ConditionTypeHasPrefix, []string(nil)}, // 1 arg required but no closing bracket
		{"HasSuffixtest)", true, ConditionTypeHasPrefix, []string(nil)}, // 1 arg required but no open bracket
	}

	for _, c := range cases {
		condition, err := NewCondition(c.conditionString)
		if c.expectError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, c.expectedType, condition.Type)
			assert.Equal(t, c.expectedType.expectedArgsCount(), uint(len(condition.Args)))
			assert.Equal(t, c.expectedArgs, condition.Args)
		}
	}
}
