package agent

func (a *Agent) _updateState() {
	for stateKey, stateVal := range a._state._states {
		if stateType, seen := agentStateTypes[stateKey]; seen {
			//if a._state._states[stateKey] < stateType._threshold &&
			//	a._state._states[stateKey] + stateType._perTurn >= stateType._threshold {
			//	// TODO maybe create a feeling and add to mind
			//}
			a._state._states[stateKey] += stateType._perTurn
		} else if stateVal > 0 { // experience fades away if remaining number of turns is positive
			a._state._states[stateKey] --
		}
	}
}

type agentState struct {
	_agent *Agent
	_states map[int]int
}

func (s *agentState) _update(key, value int) {
	var c *agentChange

	if _, seen := agentStateTypes[key]; seen {
		c = &agentChange{
			_type: key,
			_oldVal: s._states[key],
			_newVal: s._states[key] + value,
		}

		s._states[key] += value
	} else if experienceType, seen := agentExperienceTypes[key]; seen {
		c = &agentChange{
			_type: key,
			_oldVal: 0,
			_newVal: 0,
		}

		s._states[key] += experienceType._duration
	} else {
		return
	}

	s._agent._updateActionOutcome([]change{c})
}

func newAgentState(a *Agent) *agentState {
	states := make(map[int]int)

	for stateType := range agentStateTypes {
		states[stateType] = 30
	}

	for experienceType := range agentExperienceTypes {
		states[experienceType] = 0
	}

	return &agentState{
		_agent: a,
		_states: states,
	}
}

type agentStateType struct {
	_type       int
	_threshold  int
	_pointValue int
	_perTurn    int
}

// very basic design, assume every experience has a set duration and it does not affect the value
type agentExperienceType struct {
	_type     int
	_duration int
	_value    int
}

const (
	agentStateTypeHunger = iota
	agentExperienceTypeSweet
)

var agentStateTypes = map[int]*agentStateType{
	agentStateTypeHunger: {
		_type:       agentStateTypeHunger,
		_threshold:  50,
		_pointValue: -1,
		_perTurn:    1,
	},
}

var agentExperienceTypes = map[int]*agentExperienceType{
	agentExperienceTypeSweet: {
		_type:     agentExperienceTypeSweet,
		_duration: 10,
		_value:    15,
	},
}
