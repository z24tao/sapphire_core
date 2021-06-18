package agent

import (
	"fmt"
)

type quantitativeAttributeChange struct {
	*commonChange
	objectType objectType
	attrType   int
	increase   bool
	// TODO add scale e.g. 10^0, 10^1, etc.
}

func (c *quantitativeAttributeChange) toString(indent string, _, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("quantitativeAttributeChange: %s", attrTypes[c.attrType])

	changeText := "decrease"
	if c.increase {
		changeText = "increase"
	}
	result += fmt.Sprintf(" [%s],", changeText)
	result += fmt.Sprintf(" value: %.2f", c.value)
	return result
}

func (c *quantitativeAttributeChange) match(other concept) bool {
	o, ok := other.(*quantitativeAttributeChange)
	if !ok {
		return false
	}

	return c.objectType.match(o.objectType) && c.attrType == o.attrType && c.increase == o.increase
}

func (c *quantitativeAttributeChange) before() condition {
	return nil
}

func (c *quantitativeAttributeChange) after() condition {
	return nil
}

func (c *quantitativeAttributeChange) precedes(other change) bool {
	return c.match(other)
}

func (a *Agent) newQuantitativeAttributeChange(t objectType, attrType int, increase bool) *quantitativeAttributeChange {
	c := &quantitativeAttributeChange{
		commonChange: newCommonChange(),
		objectType:   t,
		attrType:     attrType,
		increase:     increase,
	}

	c.value = 20 // TODO DELETE HARDCODED VALUE
	if c.increase {
		c.value = 0
	}

	c = a.memory.add(c).(*quantitativeAttributeChange)
	return c
}
