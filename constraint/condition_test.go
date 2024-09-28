package constraint

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

func TestApplyCondition(t *testing.T) {
	var cases = []struct {
		name            string
		conditionString string
		expectError     bool
		successInputs   []string
		failInputs      []string
	}{
		{"numeric-test", "IsNumeric", false, []string{"-19", "40", "9999999999999"}, []string{"wow", "-", "!@!#$!"}},
		{"boolean-test", "IsBoolean", false, []string{"true", "false"}, []string{"yes", "-", "222159", "!@!#$!"}},
		{"unique-test", "Unique", false, []string{"unique1", "unique2", "unique3"}, []string{"unique1", "unique2", "unique3"}},
		{"prefix-test", "HasPrefix(test-)", false, []string{"test-scenario1", "test-case1", "test-please"}, []string{"not-test-", "ttest-failed"}},
		{"suffix-test", "HasSuffix(-test)", false, []string{"some-test", "another-test"}, []string{"not-test-", "failing-tests"}},

		// Make sure an invalid constraint is forced to validate false to everything
		{"invalid-test", "SomethingInvalid(arg1, arg2)", true, []string{}, []string{"test test test", "1237532123", "true", "$!&@#($)"}},
	}

	for _, c := range cases {
		constraint, err := NewConstraint(c.name, all_pattern, c.conditionString)
		if c.expectError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}

		for _, successTest := range c.successInputs {
			assert.True(t, constraint.Condition.ApplyCondition(successTest))
		}
		for _, failTest := range c.failInputs {
			assert.False(t, constraint.Condition.ApplyCondition(failTest))
		}
	}
}
