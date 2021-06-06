package agent

type agentState struct {
	agent  *Agent
	states map[int]int
}

type agentStateType struct {
	stateType  int
	threshold  int
	pointValue int
	perTurn    int
}

// very basic design, assume every experience has a set duration and it does not affect the value
type agentExperienceType struct {
	expType  int
	duration int
	value    int
}

func (s *agentState) _update(key, value int) {
	var c *agentChange

	if _, seen := agentStateTypes[key]; seen {
		c = &agentChange{
			t:        key,
			deltaVal: value,
		}

		s.states[key] += value
	} else if experienceType, seen := agentExperienceTypes[key]; seen {
		c = &agentChange{
			t:        key,
			deltaVal: value,
		}

		s.states[key] += experienceType.duration
	} else {
		return
	}

	s.agent.recordActionChanges([]change{c})
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
		agent:  a,
		states: states,
	}
}

const (
	agentStateTypeHunger = iota
	agentExperienceTypeSweet
)

var agentStateTypeNames = map[int]string{
	agentStateTypeHunger:     "hunger",
	agentExperienceTypeSweet: "sweet",
}

var agentStateTypes = map[int]*agentStateType{
	agentStateTypeHunger: {
		stateType:  agentStateTypeHunger,
		threshold:  50,
		pointValue: -1,
		perTurn:    1,
	},
}

var agentExperienceTypes = map[int]*agentExperienceType{
	agentExperienceTypeSweet: {
		expType:  agentExperienceTypeSweet,
		duration: 10,
		value:    15,
	},
}
