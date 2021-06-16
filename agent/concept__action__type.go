package agent

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	actionConditionTypeStart = iota
	actionConditionTypeStop
	actionConditionTypeObservedAtStart // conditions observed when action is started, used to train hypothesis
	actionConditionTypeObservedAtDone  // conditions observed when action is done, used to train hypothesis
)

var actionConditionTypes = map[int]string{
	actionConditionTypeStart:           "actionConditionTypeStart",
	actionConditionTypeStop:            "actionConditionTypeStop",
	actionConditionTypeObservedAtStart: "actionConditionTypeObservedAtStart",
	actionConditionTypeObservedAtDone:  "actionConditionTypeObservedAtDone",
}

type actionType interface {
	concept
	conceptType
	getConditions() map[int]map[condition]bool // condition type enum -> set of conditions
	getCausations() map[*causation]bool        // set of causations
	getAttempts() int
	attempt()
	getHypotheses() (map[condition]map[change]*hypothesis, map[change]map[condition]*hypothesis)
}

// the purpose of this struct is to remove duplicated code from implementations
type commonActionType struct {
	*commonConcept
	conditions         map[int]map[condition]bool
	causations         map[*causation]bool
	attempts           int
	forwardHypotheses  map[condition]map[change]*hypothesis // hypotheses of "this condition would cause this change"
	backwardHypotheses map[change]map[condition]*hypothesis // hypotheses of "this change is caused by this condition"
}

func (t *commonActionType) match(_ concept) bool {
	return false
}

func (t *commonActionType) toString(indent string, recursive, _ bool) string {
	result := fmt.Sprintf(" attempts %d,", t.attempts)
	result += fmt.Sprintf(" causations (%d)", len(t.causations))

	if len(t.causations) == 0 {
		return result
	}

	result += ": [\n"
	for c := range t.causations {
		result += c.toString(indent+"  ", recursive, true) + "\n"
	}
	result += indent + "]"
	return result
}

// returns: condition type -> set of conditions
func (t *commonActionType) getConditions() map[int]map[condition]bool {
	return t.conditions
}

// set of causations
func (t *commonActionType) getCausations() map[*causation]bool {
	return t.causations
}

func (t *commonActionType) getAttempts() int {
	return t.attempts
}

func (t *commonActionType) attempt() {
	t.attempts++
}

func (t *commonActionType) getHypotheses() (map[condition]map[change]*hypothesis, map[change]map[condition]*hypothesis) {
	return t.forwardHypotheses, t.backwardHypotheses
}

func newCommonActionType() *commonActionType {
	t := &commonActionType{
		commonConcept:      newCommonConcept(),
		conditions:         map[int]map[condition]bool{}, // condition type enum -> set of conditions
		causations:         map[*causation]bool{},        // set of causations
		attempts:           0,
		forwardHypotheses:  map[condition]map[change]*hypothesis{},
		backwardHypotheses: map[change]map[condition]*hypothesis{},
	}

	for actionConditionType := range actionConditionTypes {
		t.conditions[actionConditionType] = map[condition]bool{}
	}

	return t
}

// the expected value of taking this action once
func actionTypeValue(t actionType) float64 {
	curiosityBase := rand.Float64() * 30
	//if learned == 0 {
	//	curiosityBase = rand.Float64() * 1
	//}
	v := math.Max(curiosityValue*math.Pow(0.8, float64(t.getAttempts())), curiosityBase)

	for c := range t.getCausations() {
		v += c.change.getValue() * float64(c.occurrences) / float64(t.getAttempts())
	}

	return v
}
