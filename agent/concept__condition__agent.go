package agent

import "fmt"

type agentCondition struct {
	*commonConcept
	stateType int

	// positive - true: agent IS in specified state
	//            false: agent IS NOT in specified state
	positive  bool
	checked   int
	satisfied int
}

func (c *agentCondition) toString(indent string, _, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("agentCondition: %s,", agentStateTypes[c.stateType])
	result += fmt.Sprintf(" positive: %t,", c.positive)
	result += fmt.Sprintf(" checked: %d, satisfied: %d", c.checked, c.satisfied)

	return result
}

func (c *agentCondition) isSatisfied(a *Agent) bool {
	stateVal, seen := a.state.states[c.stateType]
	satisfied := false
	if !seen {
		satisfied = c.positive == false
	} else if stateInfo, seen := agentStateInfos[c.stateType]; seen {
		satisfied = c.positive == (stateVal > stateInfo.threshold)
	} else if _, seen := agentExperienceInfos[c.stateType]; seen {
		satisfied = c.positive == (stateVal > 0)
	}

	c.checked++
	if satisfied {
		c.satisfied++
	}
	return satisfied
}

func (c *agentCondition) match(other concept) bool {
	if o, ok := other.(*agentCondition); ok {
		return c.stateType == o.stateType && c.positive == o.positive
	}

	return false
}

func (c *agentCondition) buildChanges(other condition) []change {
	return []change{}
}

func (a *Agent) newAgentCondition(stateType int, positive bool) *agentCondition {
	c := &agentCondition{
		commonConcept: a.newCommonConcept(),
		stateType:     stateType,
		positive:      positive,
	}

	c = a.memory.add(c).(*agentCondition)
	return c
}
