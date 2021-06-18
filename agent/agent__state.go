package agent

const (
	agentStateTypeHunger = iota
	agentExperienceTypeSweet
)

var agentStateTypes = map[int]string{
	agentStateTypeHunger:     "hunger",
	agentExperienceTypeSweet: "sweet",
}

var agentStateInfos = map[int]*agentStateInfo{
	agentStateTypeHunger: {
		stateType:  agentStateTypeHunger,
		threshold:  50,
		pointValue: -1,
		perTurn:    1,
	},
}

var agentExperienceInfos = map[int]*agentExperienceInfo{
	agentExperienceTypeSweet: {
		expType:  agentExperienceTypeSweet,
		duration: 10,
		value:    15,
	},
}

// agentState - Agent States
type agentState struct {
	agent  *Agent
	states map[int]int
}

// agentStateInfo - Info about agent state
type agentStateInfo struct {
	stateType  int
	threshold  int
	pointValue int
	perTurn    int
}

// agentExperienceInfo - very basic design, assume every experience has a set duration and it does not affect the value
type agentExperienceInfo struct {
	expType  int
	duration int
	value    int
}

// TODO: Be more descriptive

// update - Updates Agent State
func (s *agentState) update(key, value int) {
	c := s.agent.newAgentChange(key, value)

	if info, seen := agentStateInfos[key]; seen {
		s.states[key] += value
		if s.states[key] > info.threshold {
			s.agent.mind.addItem(s.agent.newAgentCondition(key, true), 1.0)
		}
	} else if experienceType, seen := agentExperienceInfos[key]; seen {
		s.states[key] += experienceType.duration
		if s.states[key] > 0 {
			s.agent.mind.addItem(s.agent.newAgentCondition(key, true), 1.0)
		}
	} else {
		return
	}

	s.agent.recordActionChanges([]change{c})
}

// newAgentState - New instance of Agent State
func newAgentState(a *Agent) *agentState {
	states := make(map[int]int)

	for stateType := range agentStateInfos {
		states[stateType] = 30
	}

	for experienceType := range agentExperienceInfos {
		states[experienceType] = 0
	}

	return &agentState{
		agent:  a,
		states: states,
	}
}

