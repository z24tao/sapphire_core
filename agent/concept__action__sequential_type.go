package agent

import "fmt"

type sequentialActionType struct {
	*commonActionType
	first      actionType
	next       actionType
	isFunction bool
}

func (t *sequentialActionType) match(other singletonConcept) bool {
	o, ok := other.(*sequentialActionType)
	if !ok {
		return false
	}

	return t.first.match(o.first) &&
		t.next.match(o.next) &&
		t.isFunction == o.isFunction
}

func (t *sequentialActionType) toString(indent string, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("conditionalActionType\n")
	result += fmt.Sprintf(" first: %s\n", t.first.toString(indent+"  ", false))
	result += fmt.Sprintf(" next: %s\n", t.next.toString(indent+"  ", false))
	result += fmt.Sprintf(" value: %.2f", actionTypeValue(t))
	result += t.commonActionType.toString(indent, indentFirstLine)
	return result
}

func (t *sequentialActionType) instantiate() concept {
	a := &sequentialAction{
		commonAction: newCommonAction(),
		actionType:   t,
		first:        t.first.instantiate().(action),
		doneFirst:    false,
	}

	if t.next == t {
		a.next = a
	} else {
		a.next = t.next.instantiate().(action)
	}

	return a
}

func (a *Agent) newSequentialActionType(f, n actionType) *sequentialActionType {
	t := &sequentialActionType{
		commonActionType: newCommonActionType(),
		first:            f,
		next:             n,
	}

	t = a.memory.add(t).(*sequentialActionType)
	return t
}