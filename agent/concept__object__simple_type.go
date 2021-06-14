package agent

import (
	"fmt"
)

type simpleObjectType struct {
	*commonObjectType
	debugName string
}

func (t *simpleObjectType) toString(indent string, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += fmt.Sprintf("simpleObjectType: %s", t.debugName)
	return result
}

func (t *simpleObjectType) instantiate() concept {
	return &simpleObject{
		commonObject: newCommonObject(),
		objectType:   t,
	}
}

func (a *Agent) newSimpleObjectType(debugName string) *simpleObjectType {
	t := &simpleObjectType{
		commonObjectType: newCommonObjectType(),
		debugName:        debugName,
	}

	t = a.memory.add(t).(*simpleObjectType)
	return t
}
