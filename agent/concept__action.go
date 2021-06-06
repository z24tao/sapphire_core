package agent

// action instance lifecycle enums
const (
	actionStateIdle = iota
	actionStateActive
	actionStatePaused
	actionStateDone
)

// enum -> name map for debug
var actionStates = map[int]string{
	actionStateIdle:   "idle",
	actionStateActive: "active",
	actionStatePaused: "paused",
	actionStateDone:   "done",
}

/*
	The action interface is used to control individual instances of actions, this interface is
	designed to recursively construct larger actions from smaller ones, in order to simulate human
	behavior.
*/
type action interface {
	concept
	// TODO add subject and object
	getType() actionType
	start(agent *Agent) bool // returns success
	step(agent *Agent)
	// TODO add pause()
	stop(agent *Agent) bool // returns success
	getState() int
	getOutcome() *actionOutcome
}

// the purpose of this struct is to remove duplicated code from implementations
type commonAction struct {
	*commonConcept
	state   int
	outcome *actionOutcome
}

func (a *commonAction) getState() int {
	return a.state
}

func (a *commonAction) getOutcome() *actionOutcome {
	if a.state != actionStateDone {
		return nil
	}
	return a.outcome
}

// action instances are constructed in idle state with nil outcome
func newCommonAction() *commonAction {
	return &commonAction{
		commonConcept: newCommonConcept(),
		state:         actionStateIdle,
		outcome:       nil,
	}
}
