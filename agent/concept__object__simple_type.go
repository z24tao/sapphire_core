package agent

import "fmt"

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

func newSimpleObjectType(debugName string) *simpleObjectType {
	return &simpleObjectType{
		commonObjectType: newCommonObjectType(),
		debugName:        debugName,
	}
}
