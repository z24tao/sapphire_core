package agent

import "fmt"

type causation struct {
	*commonConcept
	change      change
	occurrences int
}

func (c *causation) toString(indent string, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("causation occurrences: %d,", c.occurrences)
	result += fmt.Sprintf(" change: %s", c.change.toString(indent+"  ", false))
	return result
}

func newCausation(change change, actionType actionType) *causation {
	c := &causation{
		commonConcept: newCommonConcept(),
		change:        change,
		occurrences:   1,
	}

	c.commonConcept.assocs[actionType] = 1.0
	return c
}
