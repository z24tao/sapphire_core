package agent

import "fmt"

type sequentialActionType struct {
	*commonActionType
	first actionType
	next  actionType
	isFunction bool
}

func (t *sequentialActionType) toString(indent string, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("conditionalActionType\n")
	result += fmt.Sprintf(" first: %s\n", t.first.toString(indent + "  ", false))
	result += fmt.Sprintf(" next: %s\n", t.next.toString(indent + "  ", false))
	result += fmt.Sprintf(" value: %.2f", actionTypeValue(t))
	result += t.commonActionType.toString(indent, indentFirstLine)
	return result
}

func (t *sequentialActionType) instantiate() concept {
	a := &sequentialAction{
		commonAction: newCommonAction(),
		actionType: t,
		first: t.first.instantiate().(action),
		doneFirst: false,
	}

	if t.next == t {
		a.next = a
	} else {
		a.next = t.next.instantiate().(action)
	}

	return a
}

func newSequentialActionType(f, n actionType) *sequentialActionType {
	t := &sequentialActionType{
		commonActionType: newCommonActionType(),
		first:            f,
		next:             n,
	}

	return t
}
