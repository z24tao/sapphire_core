package agent

import (
	"fmt"
)

type agentChange struct {
	*commonChange
	stateType       int
	deltaVal        int
	beforeCondition *agentCondition
	afterCondition  *agentCondition
}

func (c *agentChange) toString(indent string, _, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("agentChange: %s", agentStateTypes[c.stateType])
	result += fmt.Sprintf(" [%d],", c.deltaVal)
	result += fmt.Sprintf(" value: %.2f", c.getValue())
	return result
}

func (c *agentChange) match(other concept) bool {
	o, ok := other.(*agentChange)

	return ok && c.stateType == o.stateType && c.deltaVal == o.deltaVal
}

// agentChange value is defined, cannot be set
func (c *agentChange) setValue(float64) {
	return
}

func (c *agentChange) getValue() float64 {
	if experienceType, seen := agentExperienceInfos[c.stateType]; seen {
		return float64(experienceType.value)
	}

	stateType := agentStateInfos[c.stateType]
	return float64(c.deltaVal * stateType.pointValue)
}

// agentChange does not require condition therefore does not require before and after
func (c *agentChange) before() condition {
	return c.beforeCondition
}

func (c *agentChange) after() condition {
	return c.afterCondition
}

// agent changes cannot proceed each other
func (c *agentChange) precedes(change) bool {
	return false
}

func (a *Agent) newAgentChange(stateType, deltaVal int) *agentChange {
	c := &agentChange{
		commonChange: a.newCommonChange(),
		stateType:    stateType,
		deltaVal:     deltaVal,
	}

	beforeCondition := a.newAgentCondition(stateType, deltaVal >= 0)
	afterCondition := a.newAgentCondition(stateType, deltaVal < 0)
	beforeCondition.addAssoc(c, 0.5)
	c.addAssoc(beforeCondition, 0.5)
	afterCondition.addAssoc(c, 0.5)
	c.addAssoc(afterCondition, 0.5)
	c.beforeCondition = beforeCondition
	c.afterCondition = afterCondition

	c = a.memory.add(c).(*agentChange)
	return c
}
