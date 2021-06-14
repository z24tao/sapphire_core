package agent

import (
	"fmt"
)

type conditionalAction struct {
	*commonAction
	actionType *conditionalActionType
	passAction action
	failAction action
	passed     bool
}

func (a *conditionalAction) toString(indent string, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("conditionalAction")
	result += fmt.Sprintf(" condition passed: %t", a.passed)
	result += fmt.Sprintf(" state: %s", actionStates[a.state])
	result += fmt.Sprintf(" type: %s", a.actionType.toString(indent, indentFirstLine))
	return result
}

func (a *conditionalAction) getType() actionType {
	return a.actionType
}

func (a *conditionalAction) start(agent *Agent) bool {
	if a.state != actionStateIdle {
		return false
	}

	for cond := range a.actionType.conditions[actionConditionTypeStart] {
		if !cond.isSatisfied(agent) {
			return false
		}
	}

	a.passed = a.actionType.condition.isSatisfied(agent)
	a.state = actionStateActive
	a.actionType.attempts++
	return true
}

func (a *conditionalAction) step(agent *Agent) {
	if a.state != actionStateActive {
		return
	}

	childAction := a.failAction
	if a.passed {
		childAction = a.passAction
	}

	if childAction.getState() == actionStateIdle {
		if !childAction.start(agent) {
			a.state = actionStateDone
			a.outcome = newActionOutcome(false, false, false, true)
			return
		}
	}

	if childAction.getState() == actionStateActive {
		childAction.step(agent)
	}

	if childAction.getState() == actionStateDone {
		a.outcome = childAction.getOutcome()
		a.state = actionStateDone
	}
}

func (a *conditionalAction) stop(agent *Agent) bool {
	if a.state != actionStateActive {
		return false
	}

	childAction := a.failAction
	if a.passed {
		childAction = a.passAction
	}

	childStopped := childAction.stop(agent)
	if !childStopped {
		return false
	} else {
		a.outcome = childAction.getOutcome()
		a.state = actionStateDone
	}

	return true
}
