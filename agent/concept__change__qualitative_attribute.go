package agent

import (
	"fmt"
)

type qualitativeAttributeChange struct {
	*commonChange
	objectType objectType
	attrType   int
	oldVal     int
	newVal     int
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

func (c *qualitativeAttributeChange) match(other change) bool {
	o, ok := other.(*qualitativeAttributeChange)
	if !ok {
		return false
	}

	return c.objectType == o.objectType && c.attrType == o.attrType && c.oldVal == o.oldVal && c.newVal == o.newVal
}

func (c *qualitativeAttributeChange) before() condition {
	return newAttributeCondition(c.objectType, c.attrType, c.oldVal)
}

func (c *qualitativeAttributeChange) after() condition {
	return newAttributeCondition(c.objectType, c.attrType, c.newVal)
}

func (c *qualitativeAttributeChange) precedes(other change) bool {
	if _, ok := other.(*qualitativeAttributeChange); !ok {
		return false
	}

	return c.after().match(other.before())
}

func newQualitativeAttributeChange(objType objectType, attrType, oldVal, newVal int) *qualitativeAttributeChange {
	return &qualitativeAttributeChange{
		commonChange: newCommonChange(),
		objectType:   objType,
		attrType:     attrType,
		oldVal:       oldVal,
		newVal:       newVal,
	}
}
