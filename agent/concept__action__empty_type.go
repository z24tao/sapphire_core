package agent

type emptyActionType struct {
	*commonActionType
}

func (t *emptyActionType) toString(indent string, recursive, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += "emptyActionType"
	return result
}

func (t *emptyActionType) match(other concept) bool {
	_, ok := other.(*emptyActionType)
	return ok
}

func (t *emptyActionType) instantiate() concept {
	return &emptyAction{
		commonAction: t.agent.newCommonAction(),
	}
}

var emptyActionTypeSingleton *emptyActionType
