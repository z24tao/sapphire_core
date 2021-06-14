package agent

type emptyAction struct {
	*commonAction
}

func (a *emptyAction) toString(indent string, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += "emptyAction"
	return result
}

func (a *emptyAction) getType() actionType {
	return emptyActionTypeSingleton
}

func (a *emptyAction) start(*Agent) bool {
	a.state = actionStateDone
	return true
}

func (a *emptyAction) step(*Agent) {
	return
}

func (a *emptyAction) stop(agent *Agent) bool {
	return a.start(agent)
}
