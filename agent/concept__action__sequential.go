package agent

import "fmt"

type sequentialAction struct {
	*commonAction
	actionType *sequentialActionType
	first      action
	next       action
	doneFirst  bool
}

func (a *sequentialAction) match(_ concept) bool {
	return false
}

func (a *sequentialAction) toString(indent string, recursive, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("sequentialAction")
	result += fmt.Sprintf(" done first: %t", a.doneFirst)
	result += fmt.Sprintf(" state: %s", actionStates[a.state])
	result += fmt.Sprintf(" type: %s", a.actionType.toString(indent, recursive, indentFirstLine))
	return result
}

func (a *sequentialAction) getType() actionType {
	return a.actionType
}

func (a *sequentialAction) start(agent *Agent) bool {
	if a.state != actionStateIdle {
		return false
	}

	for cond := range a.actionType.conditions[actionConditionTypeStart] {
		if !cond.isSatisfied(agent) {
			return false
		}
	}

	a.doneFirst = false
	a.state = actionStateActive
	return true
}

func (a *sequentialAction) step(agent *Agent) {
	if a.state != actionStateActive {
		return
	}

	childAction := a.first
	if a.doneFirst {
		childAction = a.next
	}

	if childAction.getState() == actionStateIdle {
		if !childAction.start(agent) {
			a.state = actionStateDone

			// if action is a loop, do not throw an exception
			a.outcome = newActionOutcome(false, false, false, !(a.next == a))
			return
		}
	}

	if childAction.getState() == actionStateActive {
		childAction.step(agent)
	}

	if childAction.getState() == actionStateDone {
		childOutcome := childAction.getOutcome()
		if a.doneFirst { // next finished
			a.state = actionStateDone
			a.outcome = childOutcome
		} else { // first finished
			a.doneFirst = true
			if childOutcome.exception {
				a.state = actionStateDone
				a.outcome = newActionOutcome(false, false, false, true)
				return
			}

			if a.next == a {
				a.next = a.actionType.instantiate().(action)
			}
		}
	}
}

func (a *sequentialAction) stop(agent *Agent) bool {
	if a.state != actionStateActive {
		return false
	}

	childAction := a.first
	if a.doneFirst {
		childAction = a.next
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
