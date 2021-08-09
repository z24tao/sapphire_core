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
	getPreConditions() map[condition]bool
	getPostConditions() map[condition]bool
	buildCausations() // for multi-step actions to build causations using preConditions and postConditions
}

// the purpose of this struct is to remove duplicated code from implementations
type commonAction struct {
	*commonConcept
	state          int
	outcome        *actionOutcome
	preConditions  map[condition]bool
	postConditions map[condition]bool
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

func (a *commonAction) getPreConditions() map[condition]bool {
	return a.preConditions
}

func (a *commonAction) getPostConditions() map[condition]bool {
	return a.postConditions
}

func (a *commonAction) buildCausations() {

}

// action instances are constructed in idle state with nil outcome
func (a *Agent) newCommonAction() *commonAction {
	return &commonAction{
		commonConcept:  a.newCommonConcept(),
		state:          actionStateIdle,
		outcome:        nil,
		preConditions:  map[condition]bool{},
		postConditions: map[condition]bool{},
	}
}
