package agent

import "fmt"

type hypothesis struct {
	condition      condition
	change         change
	conditionMatch int  // number of times the condition is observed
	changeMatch    int  // number of times the change is observed
	bothMatch      int  // number of times both the condition and the change is observed

	// true (forward): we hypothesize that taking this action under this condition would cause this change
	// false (backward): we hypothesize that if we took this action and caused this change, it must have
	//   previously been this condition
	forward        bool
}

func (h *hypothesis) evaluate() float64 {
	if h.forward {
		if h.conditionMatch < 5 {
			return 0.5
		}

		return float64(h.bothMatch) / float64(h.conditionMatch)
	} else {
		if h.changeMatch < 5 {
			return 0.5
		}

		return float64(h.bothMatch) / float64(h.changeMatch)
	}
}

func (h *hypothesis) toString(indent string, indentFirstLine bool) string {
	if h.evaluate() < 0.1 {
		return ""
	}

	result := ""
	if indentFirstLine {
		result += indent
	}
	if h.forward {
		result += fmt.Sprintf(indent + "forward hypothesis:\n")
	} else {
		result += fmt.Sprintf(indent + "backward hypothesis:\n")
	}
	result += fmt.Sprintf(indent + "  matches: causation = %d, change = %d, both = %d\n",
		h.conditionMatch, h.changeMatch, h.bothMatch)
	result += fmt.Sprintf(indent + "  condition: %s\n", h.condition.toString(indent + "  ", false))
	result += fmt.Sprintf(indent + "  change: %s\n", h.change.toString(indent + "  ", false))
	result += fmt.Sprintf(indent + "  score: %.3f\n", h.evaluate())
	return result
}

var hypothesisCount = 0 // TODO remove this

func newHypothesis(condition condition, change change, forward bool) *hypothesis {
	hypothesisCount ++
	return &hypothesis{
		condition:      condition,
		change:         change,
		conditionMatch: 0,
		changeMatch:    0,
		bothMatch:      0,
		forward:        forward,
	}
}
