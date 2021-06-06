package agent

import (
	"github.com/z24tao/sapphire_core/world"
	"fmt"
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

func (c *aaiCondition) match(other condition) bool {
	if o, ok := other.(*aaiCondition); ok {
		return c.aai == o.aai && c.enabled == o.enabled
	}

	return false
}

func (c *aaiCondition) toString(indent string, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("aaiCondition: %s", c.aai.Name)
	result += fmt.Sprintf(" enabled: %t", c.enabled)

	return result
}

func newAAICondition(aat *atomicActionType, enabled bool) *aaiCondition {
	c := &aaiCondition{
		commonConcept: newCommonConcept(),
		aai:     aat.aai,
		enabled: enabled,
	}

	c.commonConcept.assocs[aat] = 1.0
	return c
}
