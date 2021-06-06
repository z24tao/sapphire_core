package agent

/*
when a child action completes, its outcome field is populated for the parent action to decide
  its execution
*/
type actionOutcome struct {
	shouldReturn bool
	shouldContinue bool
	shouldBreak bool
	exception bool
}

func newActionOutcome(shouldReturn, shouldContinue, shouldBreak, exception bool) *actionOutcome {
	return &actionOutcome{
		shouldReturn:   shouldReturn,
		shouldContinue: shouldContinue,
		shouldBreak:    shouldBreak,
		exception:      exception,
	}
}
