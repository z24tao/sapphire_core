package agent

import (
	"fmt"
)

type qualitativeAttributeChange struct {
	*commonChange
	objectType      objectType
	attrType        int
	oldVal          int
	newVal          int
	beforeCondition *attributeCondition
	afterCondition  *attributeCondition
}

func (c *qualitativeAttributeChange) toString(indent string, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("qualitativeAttributeChange: %s", attrTypes[c.attrType])
	result += fmt.Sprintf(" [%s -> %s],", attrVals[c.attrType][c.oldVal], attrVals[c.attrType][c.newVal])
	result += fmt.Sprintf(" value: %.2f", c.value)
	return result
}

func (c *qualitativeAttributeChange) match(other singletonConcept) bool {
	o, ok := other.(*qualitativeAttributeChange)
	if !ok {
		return false
	}

	return c.objectType.match(o.objectType) && c.attrType == o.attrType && c.oldVal == o.oldVal && c.newVal == o.newVal
}

func (c *qualitativeAttributeChange) before() condition {
	return c.beforeCondition
}

func (c *qualitativeAttributeChange) after() condition {
	return c.afterCondition
}

func (c *qualitativeAttributeChange) precedes(other change) bool {
	if _, ok := other.(*qualitativeAttributeChange); !ok {
		return false
	}

	return c.after().match(other.before())
}

func (a *Agent) newQualitativeAttributeChange(t objectType, attrType, oldVal, newVal int) *qualitativeAttributeChange {
	c := &qualitativeAttributeChange{
		commonChange:    newCommonChange(),
		objectType:      t,
		attrType:        attrType,
		oldVal:          oldVal,
		newVal:          newVal,
		beforeCondition: a.newAttributeCondition(t, attrType, oldVal),
		afterCondition:  a.newAttributeCondition(t, attrType, newVal),
	}

	c = a.memory.add(c).(*qualitativeAttributeChange)
	return c
}
