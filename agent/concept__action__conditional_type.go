package agent

import "fmt"

type conditionalActionType struct {
	*commonActionType
	condition      condition
	passActionType actionType
	failActionType actionType
}

func (t *conditionalActionType) match(other singletonConcept) bool {
	o, ok := other.(*conditionalActionType)
	if !ok {
		return false
	}

	return t.condition.match(o.condition) &&
		t.passActionType.match(o.passActionType) &&
		t.failActionType.match(o.failActionType)
}

func (t *conditionalActionType) toString(indent string, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("conditionalActionType")
	result += fmt.Sprintf(" condition: %s\n", t.condition.toString(indent+"  ", false))
	result += fmt.Sprintf(" passActionType: %s\n", t.passActionType.toString(indent+"  ", false))
	result += fmt.Sprintf(" failActionType: %s\n", t.failActionType.toString(indent+"  ", false))
	result += fmt.Sprintf(" value: %.2f", actionTypeValue(t))
	result += t.commonActionType.toString(indent, indentFirstLine)
	return result
}

func (t *conditionalActionType) instantiate() concept {
	return &conditionalAction{
		commonAction: newCommonAction(),
		actionType:   t,
		passAction:   t.passActionType.instantiate().(action),
		failAction:   t.failActionType.instantiate().(action),
		passed:       false,
	}
}

func (a *Agent) newConditionalActionType(c condition, p, f actionType) *conditionalActionType {
	t := &conditionalActionType{
		commonActionType: newCommonActionType(),
		condition:        c,
		passActionType:   p,
		failActionType:   f,
	}

	t = a.memory.add(t).(*conditionalActionType)
	return t
}
