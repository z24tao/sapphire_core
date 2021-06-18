package agent

import (
	"fmt"
	"github.com/z24tao/sapphire_core/world"
)

// condition of whether an atomic action interface is enabled
type aaiCondition struct {
	*commonConcept
	aai     *world.AtomicActionInterface
	enabled bool
}

func (c *aaiCondition) isSatisfied(*Agent) bool {
	return c.aai.Enabled == c.enabled
}

func (c *aaiCondition) match(other concept) bool {
	if o, ok := other.(*aaiCondition); ok {
		return c.aai == o.aai && c.enabled == o.enabled
	}

	return false
}

func (c *aaiCondition) toString(indent string, _, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("aaiCondition: %s", c.aai.Name)
	result += fmt.Sprintf(" enabled: %t", c.enabled)

	return result
}

func (a *Agent) newAAICondition(aat *atomicActionType, enabled bool) *aaiCondition {
	c := &aaiCondition{
		commonConcept: newCommonConcept(),
		aai:           aat.aai,
		enabled:       enabled,
	}

	c.addAssoc(aat, 0.5)
	aat.addAssoc(c, 0.5)
	c = a.memory.add(c).(*aaiCondition)
	return c
}
