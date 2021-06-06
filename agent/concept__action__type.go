package agent

import (
	"fmt"
	"math"
)

const (
	actionConditionTypeStart = iota
	actionConditionTypeStop
)

var actionConditionTypes = map[int]string{
	actionConditionTypeStart: "actionConditionTypeStart",
	actionConditionTypeStop:  "actionConditionTypeStop",
}

type actionType interface {
	concept
	conceptType
	getConditions() map[int]map[condition]bool // condition type enum -> set of conditions
	getCausations() map[*causation]bool        // set of causations
	getAttempts() int
}

// the purpose of this struct is to remove duplicated code from implementations
type commonActionType struct {
	*commonConcept
	conditions map[int]map[condition]bool
	causations map[*causation]bool
	attempts   int
}

func (t *commonActionType) toString(indent string, _ bool) string {
	result := fmt.Sprintf(" causations (%d)", len(t.causations))

	if len(t.causations) == 0 {
		return result
	}

	result += ": [\n"
	for c := range t.causations {
		result += c.toString(indent+"  ", true) + "\n"
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

func newCommonActionType() *commonActionType {
	t := &commonActionType{
		commonConcept: newCommonConcept(),
		conditions:    map[int]map[condition]bool{}, // condition type enum -> set of conditions
		causations:    map[*causation]bool{},        // set of causations
		attempts:      0,
	}

	for actionConditionType := range actionConditionTypes {
		t.conditions[actionConditionType] = map[condition]bool{}
	}

	return t
}

// the expected value of taking this action once
func actionTypeValue(t actionType) float64 {
	v := math.Max(curiosityValue*math.Pow(0.8, float64(t.getAttempts())), 10.0)

	for c := range t.getCausations() {
		v += c.change.getValue() * float64(c.occurrences) / float64(t.getAttempts())
	}

	return v
}
