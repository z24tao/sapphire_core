package agent

import "fmt"

type quantitativeAttributeChange struct {
	*commonChange
	objectType objectType
	attrType   int
	increase   bool
}

func (c *quantitativeAttributeChange) toString(indent string, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("qualitativeAttributeChange: %s", attrTypes[c.attrType])

	changeText := "decrease"
	if c.increase {
		changeText = "increase"
	}
	result += fmt.Sprintf(" [%s],", changeText)
	result += fmt.Sprintf(" value: %.2f", c.value)
	return result
}

func (c *quantitativeAttributeChange) match(other change) bool {
	o, ok := other.(*quantitativeAttributeChange)
	if !ok {
		return false
	}

	return c.objectType == o.objectType && c.attrType == o.attrType && c.increase == o.increase
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

func newQuantitativeAttributeChange(objType objectType, attrType int, increase bool) *quantitativeAttributeChange {
	return &quantitativeAttributeChange{
		commonChange: newCommonChange(),
		objectType:   objType,
		attrType:     attrType,
		increase:     increase,
	}
}
