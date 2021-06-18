package agent

import (
	"fmt"
)

type atomicAction struct {
	*commonAction
	actionType *atomicActionType
}

func (a *atomicAction) match(_ concept) bool {
	return false
}

func (a *atomicAction) getType() actionType {
	return a.actionType
}

func (a *atomicAction) start(agent *Agent) bool {
	if a.state != actionStateIdle {
		return false
	}

	for cond := range a.actionType.conditions[actionConditionTypeStart] {
		if !cond.isSatisfied(agent) {
			return false
		}
	}

	a.state = actionStateActive
	return true
}

func (a *atomicAction) step(agent *Agent) {
	if a.state != actionStateActive {
		return
	}

	if !a.actionType.aai.Enabled {
		return
	}

	responses := a.actionType.aai.Step()
	for _, response := range responses {
		agent.processActionResponse(response)
	}

	a.state = actionStateDone
	a.outcome = newActionOutcome(false, false, false, false)
}

func (a *atomicAction) stop(agent *Agent) bool {
	if a.state != actionStateActive {
		return false
	}

	for cond := range a.actionType.conditions[actionConditionTypeStop] {
		if !cond.isSatisfied(agent) {
			return false
		}
	}

	a.state = actionStateDone
	a.outcome = newActionOutcome(false, false, false, true)
	return true
}

func (a *atomicAction) toString(indent string, _, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("atomicAction: %s,", a.actionType.aai.Name)
	result += fmt.Sprintf(" state: %s", actionStates[a.state])
	return result
}
