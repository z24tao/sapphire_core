package agent

import (
	"../util"
	"fmt"
)

type agentChange struct {
	__assocs []concept
	_type   int
	_oldVal int
	_newVal int
}

func (c *agentChange) _assocs() []concept {
	return c.__assocs
}

func (c *agentChange) _toString() string {
	return fmt.Sprintf("agentChange, value: %v", c._getValue())
}

func (c *agentChange) _match(n change) bool {
	nn, ok := n.(*agentChange)
	// TODO reduce granularity
	return ok && c._type == nn._type && c._oldVal == nn._oldVal && c._newVal == nn._newVal
}

func (c *agentChange) _matchOld(a *Agent, oldItem concept) bool {
	if _, seen := agentExperienceTypes[c._type]; seen {
		return a._state._states[c._type] == 0
	}

	oldVal := a._state._states[c._type]
	stateType := agentStateTypes[c._type]
	return (oldVal < stateType._threshold) == (c._oldVal < stateType._threshold)
}

func (c *agentChange) _matchNew(a *Agent, newItem concept) bool {
	if _, seen := agentExperienceTypes[c._type]; seen {
		return a._state._states[c._type] > 0
	}

	newVal := a._state._states[c._type]
	stateType := agentStateTypes[c._type]
	return (newVal < stateType._threshold) == (c._newVal < stateType._threshold)
}

func (c *agentChange) _getValue() float64 {
	if experienceType, seen := agentExperienceTypes[c._type]; seen {
		return float64(experienceType._value)
	}

	stateType := agentStateTypes[c._type]
	oldVal := util.Max(0, c._oldVal-stateType._threshold)
	newVal := util.Max(0, c._newVal-stateType._threshold)
	return float64((newVal - oldVal) * stateType._pointValue)
}

func (c *agentChange) _setValue(newValue float64) {
	// nothing should be able to modify the value of a feeling... yet
}
