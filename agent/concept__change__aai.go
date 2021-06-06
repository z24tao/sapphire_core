package agent

import (
	"fmt"
)

type aaiChange struct {
	*commonChange
	actionType *atomicActionType
	enabling   bool // true: disabled -> enabled, false: enabled -> disabled
}

func (c *aaiChange) toString(indent string, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}

	result += fmt.Sprintf("aaiChange: %s", c.actionType.aai.Name)
	result += fmt.Sprintf(" enabling: %t", c.enabling)

	return result
}

func (c *aaiChange) before() condition {
	panic("implement me")
}

func (c *aaiChange) after() condition {
	panic("implement me")
}

func (c *aaiChange) precedes(change) bool {
	return false
}

func (c *aaiChange) match(other change) bool {
	otherChange, ok := other.(*aaiChange)
	return ok && c.actionType == otherChange.actionType && c.enabling == otherChange.enabling
}

func newAAIChange(t *atomicActionType, enabling bool) *aaiChange {
	return &aaiChange{
		commonChange: newCommonChange(),
		actionType:   t,
		enabling:     enabling,
	}
}
