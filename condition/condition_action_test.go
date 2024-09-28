package condition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyCondition(t *testing.T) {
	var cases = []struct {
		conditionString string
		expectError     bool
		successInputs   []string
		failInputs      []string
	}{
		{"IsNumeric", false, []string{"-19", "40", "9999999999999"}, []string{"wow", "-", "!@!#$!"}},
		{"IsBoolean", false, []string{"true", "false"}, []string{"yes", "-", "222159", "!@!#$!"}},
		{"Unique", false, []string{"unique1", "unique2", "unique3"}, []string{"unique1", "unique2", "unique3"}},
		{"HasPrefix(test-)", false, []string{"test-scenario1", "test-case1", "test-please"}, []string{"not-test-", "ttest-failed"}},
		{"HasSuffix(-test)", false, []string{"some-test", "another-test"}, []string{"not-test-", "failing-tests"}},

		// Make sure an invalid constraint is forced to validate false to everything
		{"SomethingInvalid(arg1, arg2)", true, []string{}, []string{"test test test", "1237532123", "true", "$!&@#($)"}},
	}

	for _, c := range cases {
		condition, err := NewCondition(c.conditionString)
		if c.expectError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}

		for _, successTest := range c.successInputs {
			assert.True(t, condition.ApplyCondition(successTest))
		}
		for _, failTest := range c.failInputs {
			assert.False(t, condition.ApplyCondition(failTest))
		}
	}
}
