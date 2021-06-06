package agent

import "fmt"

type conditionalActionType struct {
	*commonActionType
	condition condition
	passActionType actionType
	failActionType actionType
}

func (t *conditionalActionType) toString(indent string, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("conditionalActionType")
	result += fmt.Sprintf(" condition: %s\n", t.condition.toString(indent + "  ", false))
	result += fmt.Sprintf(" passActionType: %s\n", t.passActionType.toString(indent + "  ", false))
	result += fmt.Sprintf(" failActionType: %s\n", t.failActionType.toString(indent + "  ", false))
	result += fmt.Sprintf(" value: %.2f", actionTypeValue(t))
	result += t.commonActionType.toString(indent, indentFirstLine)
	return result
}

func (t *conditionalActionType) instantiate() concept {
	return &conditionalAction{
		commonAction: newCommonAction(),
		actionType: t,
		passAction: t.passActionType.instantiate().(action),
		failAction: t.failActionType.instantiate().(action),
		passed: false,
	}
}

func newConditionalActionType(c condition, p, f actionType) *conditionalActionType {
	t := &conditionalActionType{
		commonActionType: newCommonActionType(),
		condition:        c,
		passActionType:   p,
		failActionType:   f,
	}

	return t
}
