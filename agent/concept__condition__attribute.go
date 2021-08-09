package agent

import (
	"fmt"
)

/*
  condition of observing an object with specific attribute
  - i.e. the following attribute condition:
		attributeCondition {
		  objType: apple
		  attrType:   direction
		  attrVal:    left
		}
    is satisfied if the agent observes an apple to its left
*/
type attributeCondition struct {
	*commonConcept
	objType   objectType
	attrType  int
	attrVal   int
	checked   int
	satisfied int
}

func (c *attributeCondition) toString(indent string, recursive, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += "attributeCondition"
	if recursive {
		result += fmt.Sprintf("\n"+indent+"  objectType: %s,\n"+indent+"  ", c.objType.toString(indent+"  ", recursive, false))
	}
	result += fmt.Sprintf(" attrType: %s,", attrTypes[c.attrType])
	result += fmt.Sprintf(" attrVal: %s,", attrVals[c.attrType][c.attrVal])
	result += fmt.Sprintf(" checked: %d, satisfied: %d", c.checked, c.satisfied)
	return result
}

func (c *attributeCondition) isSatisfied(a *Agent) bool {
	mindObjs := a.mind.objects()
	matchAttrs := map[int]int{
		c.attrType: c.attrVal,
	}
	consideredAttrTypes := map[int]bool{
		c.attrType: true,
	}

	satisfied := false
	for _, mindObj := range mindObjs {
		if mindObj.matchAttrs(matchAttrs, consideredAttrTypes) {
			satisfied = true
			break
		}
	}

	c.checked++
	if satisfied {
		c.satisfied++
	}
	return satisfied
}

func (c *attributeCondition) match(other concept) bool {
	otherAttributeCondition, ok := other.(*attributeCondition)
	if !ok {
		return false
	}

	return c.objType.match(otherAttributeCondition.objType) &&
		c.attrType == otherAttributeCondition.attrType &&
		c.attrVal == otherAttributeCondition.attrVal
}

// change from c to other
func (c *attributeCondition) buildChanges(other condition) []change {
	var result []change
	o, ok := other.(*attributeCondition)
	if !ok {
		return result
	}

	if c.objType.match(o.objType) && c.attrType == o.attrType && c.attrVal != o.attrVal {
		result = append(result, c.agent.newQualitativeAttributeChange(c.objType, c.attrType, c.attrVal, o.attrVal))
	}

	return result
}

func (a *Agent) newAttributeCondition(objType objectType, attrType int, attrVal int) *attributeCondition {
	c := &attributeCondition{
		commonConcept: a.newCommonConcept(),
		objType:       objType,
		attrType:      attrType,
		attrVal:       attrVal,
	}

	c.addAssoc(objType, 0.5)
	objType.addAssoc(c, 0.5)
	c = a.memory.add(c).(*attributeCondition)
	return c
}
