package agent

import "fmt"

type sequentialActionType struct {
	*commonActionType
	first      actionType
	next       actionType
	isFunction bool
}

func (t *sequentialActionType) match(other concept) bool {
	o, ok := other.(*sequentialActionType)
	if !ok {
		return false
	}

	// this match could be called during construction, in which case for recursive sequential actions next would be nil,
	//   after construction the next would be self
	return t.first.match(o.first) &&
		((t.next == nil && o.next == o) ||
			(o.next == nil && t.next == t) ||
			(o.next == o && t.next == t) ||
			t.next.match(o.next)) &&
		t.isFunction == o.isFunction
}

func (t *sequentialActionType) toString(indent string, recursive, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("sequentialActionType")
	if !recursive {
		return result
	}

	result += fmt.Sprintf("\n"+indent+" first: %s\n", t.first.toString(indent+"  ", false, false))
	result += fmt.Sprintf(indent+" next: %s\n", t.next.toString(indent+"  ", false, false))
	result += fmt.Sprintf(indent+" value: %.2f", actionTypeValue(t))

	result += t.commonActionType.toString(indent+"  ", false, indentFirstLine)
	return result
}

func (t *sequentialActionType) instantiate() concept {
	a := &sequentialAction{
		commonAction: t.agent.newCommonAction(),
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
		commonActionType: a.newCommonActionType(),
		first:            f,
		next:             n,
	}

	t = a.memory.add(t).(*sequentialActionType)
	return t
}
