package constraint

type Constraint struct {
	Name      string
	Pattern   Pattern
	Condition Condition
}

func NewConstraint(name string, pattern string, condition string) (constraint Constraint, err error) {
	conditionObj, err := NewCondition(condition)
	if err != nil {
		return
	}
	patternObj, err := NewPattern(pattern)
	if err != nil {
		return
	}
	return Constraint{
		Name:      name,
		Pattern:   patternObj,
		Condition: conditionObj,
	}, nil
}
