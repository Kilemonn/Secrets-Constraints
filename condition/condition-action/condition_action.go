package condition_action

var (
	// TODO: Create a new instance for each different constraint definition that uses Unique
	// Need to create and hold this variable, since it's state needs to be retained
	UniqueConditionActionObj UniqueConditionAction = NewUniqueConditionAction()
)

type ConditionAction interface {
	CheckCondition(input string, args []string) bool
}
