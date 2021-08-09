package agent

import "fmt"

type conditionalActionType struct {
	*commonActionType
	condition      condition
	passActionType actionType
	failActionType actionType
}

func (t *conditionalActionType) match(other concept) bool {
	o, ok := other.(*conditionalActionType)
	if !ok {
		return false
	}

	return t.condition.match(o.condition) &&
		t.passActionType.match(o.passActionType) &&
		t.failActionType.match(o.failActionType)
}

func (t *conditionalActionType) toString(indent string, recursive, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("conditionalActionType\n")
	result += fmt.Sprintf(indent+"  condition: %s\n", t.condition.toString(indent+"  ", false, false))
	result += fmt.Sprintf(indent+"  passActionType: %s\n", t.passActionType.toString(indent+"  ", false, false))
	result += fmt.Sprintf(indent+"  failActionType: %s\n", t.failActionType.toString(indent+"  ", false, false))
	result += fmt.Sprintf(indent+"  value: %.2f", actionTypeValue(t))

	if !recursive {
		return result
	}

	result += t.commonActionType.toString(indent+"  ", false, indentFirstLine)
	return result
}

func (t *conditionalActionType) instantiate() concept {
	return &conditionalAction{
		commonAction: t.agent.newCommonAction(),
		actionType:   t,
		passAction:   t.passActionType.instantiate().(action),
		failAction:   t.failActionType.instantiate().(action),
		passed:       false,
	}
}

func (a *Agent) newConditionalActionType(c condition, p, f actionType) *conditionalActionType {
	t := &conditionalActionType{
		commonActionType: a.newCommonActionType(),
		condition:        c,
		passActionType:   p,
		failActionType:   f,
	}

	t.commonActionType.conditions[actionConditionTypeStart][c] = true

	t = a.memory.add(t).(*conditionalActionType)
	return t
}
