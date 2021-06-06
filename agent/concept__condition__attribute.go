package agent

import "fmt"

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
	objType  objectType
	attrType int
	attrVal  int
}

func (c *attributeCondition) toString(indent string, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += "attributeCondition"
	result += fmt.Sprintf(" objectType: %s", c.objType.toString(indent+"  ", false))
	result += fmt.Sprintf(" attrType: %s", attrTypes[c.attrType])
	result += fmt.Sprintf(" attrVal: %s", attrVals[c.attrType][c.attrVal])
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
	for _, mindObj := range mindObjs {
		if mindObj.match(matchAttrs, consideredAttrTypes) {
			return true
		}
	}

	return false
}

func (c *attributeCondition) match(other condition) bool {
	otherAttributeCondition, ok := other.(*attributeCondition)
	if !ok {
		return false
	}

	return c.objType == otherAttributeCondition.objType &&
		c.attrType == otherAttributeCondition.attrType &&
		c.attrVal == otherAttributeCondition.attrVal
}

func newAttributeCondition(objType objectType, attrType int, attrVal int) *attributeCondition {
	c := &attributeCondition{
		objType:  objType,
		attrType: attrType,
		attrVal:  attrVal,
	}

	c.commonConcept.assocs[objType] = 1.0

	return c
}
