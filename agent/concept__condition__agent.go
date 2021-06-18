package agent

import "fmt"

type agentCondition struct {
	*commonConcept
	stateType int

	// positive - true: agent IS in specified state
	//            false: agent IS NOT in specified state
	positive bool
}

func (c *agentCondition) toString(indent string, _, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("agentCondition: %s", agentStateTypes[c.stateType])
	result += fmt.Sprintf(" positive: %t", c.positive)

	return result
}

func (c *agentCondition) isSatisfied(a *Agent) bool {
	stateVal, seen := a.state.states[c.stateType]
	if !seen {
		return c.positive == false
	}

	if stateInfo, seen := agentStateInfos[c.stateType]; seen {
		return c.positive == (stateVal > stateInfo.threshold)
	}

	if _, seen := agentExperienceInfos[c.stateType]; seen {
		return c.positive == (stateVal > 0)
	}

	return false
}

func (c *agentCondition) match(other concept) bool {
	if o, ok := other.(*agentCondition); ok {
		return c.stateType == o.stateType && c.positive == o.positive
	}

	return false
}

func (a *Agent) newAgentCondition(stateType int, positive bool) *agentCondition {
	c := &agentCondition{
		commonConcept: newCommonConcept(),
		stateType:     stateType,
		positive:      positive,
	}

	c = a.memory.add(c).(*agentCondition)
	return c
}
