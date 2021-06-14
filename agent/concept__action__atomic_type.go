package agent

import (
	"fmt"
	"github.com/z24tao/sapphire_core/world"
)

type atomicActionType struct {
	*commonActionType
	aai *world.AtomicActionInterface
}

func (t *atomicActionType) instantiate() concept {
	return &atomicAction{
		commonAction: newCommonAction(),
		actionType:   t,
	}
}

func (t *atomicActionType) match(other singletonConcept) bool {
	o, ok := other.(*atomicActionType)
	if !ok {
		return false
	}

	return t.aai == o.aai
}

func (t *atomicActionType) toString(indent string, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("atomicActionType: %s,", t.aai.Name)
	result += fmt.Sprintf(" enabled: %t,", t.aai.Enabled)
	result += fmt.Sprintf(" value: %.2f,", actionTypeValue(t))
	result += fmt.Sprintf(" attempted: %d,", t.attempts)
	result += t.commonActionType.toString(indent, indentFirstLine)
	return result
}

func (a *Agent) newAtomicActionType(aai *world.AtomicActionInterface) *atomicActionType {
	t := &atomicActionType{
		commonActionType: newCommonActionType(),
		aai:              aai,
	}

	t.conditions[actionConditionTypeStart][a.newAAICondition(t, true)] = true
	t = a.memory.add(t).(*atomicActionType)
	return t
}
