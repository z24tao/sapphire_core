package agent

import (
	"fmt"
)

type simpleObjectType struct {
	*commonObjectType
	debugName string
}

func (t *simpleObjectType) toString(indent string, recursive, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("simpleObjectType: %s", t.debugName)

	if !recursive {
		return result
	}

	result += fmt.Sprintf("\n"+indent+"  attributes: (%d) [\n", len(t.attrs))

	for attrType, attrVal := range t.attrs {
		if qualitativeAttrTypes[attrType] {
			result += fmt.Sprintf(indent+"    %s: %s\n", attrTypes[attrType], attrVals[attrType][attrVal])
		} else {
			result += fmt.Sprintf(indent+"    %s: %d\n", attrTypes[attrType], attrVal)
		}
	}
	result += indent + "  ]"

	return result
}

func (t *simpleObjectType) instantiate() concept {
	return &simpleObject{
		commonObject: t.agent.newCommonObject(),
		objectType:   t,
	}
}

func (a *Agent) newSimpleObjectType(debugName string) *simpleObjectType {
	t := &simpleObjectType{
		commonObjectType: a.newCommonObjectType(),
		debugName:        debugName,
	}

	t = a.memory.add(t).(*simpleObjectType)
	return t
}
