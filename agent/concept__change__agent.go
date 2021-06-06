package agent

import (
	"fmt"
)

type agentChange struct {
	*commonChange
	t        int
	deltaVal int
}

func (c *agentChange) toString(indent string, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("agentChange: %s", agentStateTypeNames[c.t])
	result += fmt.Sprintf(" [%d],", c.deltaVal)
	result += fmt.Sprintf(" value: %.2f", c.getValue())
	return result
}

func (c *agentChange) match(other change) bool {
	o, ok := other.(*agentChange)

	return ok && c.t == o.t && c.deltaVal == o.deltaVal
}

// agentChange value is defined, cannot be set
func (c *agentChange) setValue(float64) {
	return
}

func (c *agentChange) getValue() float64 {
	if experienceType, seen := agentExperienceTypes[c.t]; seen {
		return float64(experienceType.value)
	}

	stateType := agentStateTypes[c.t]
	return float64(c.deltaVal * stateType.pointValue)
}

// agentChange does not require condition therefore does not require before and after
func (c *agentChange) before() condition {
	return nil
}

func (c *agentChange) after() condition {
	return nil
}

// agent changes cannot proceed each other
func (c *agentChange) precedes(change) bool {
	return false
}
