package agent

type emptyActionType struct {
	*commonActionType
}

func (t *emptyActionType) toString(indent string, indentFirstLine bool) string {
	result := ""
	if indentFirstLine {
		result += indent
	}
	result += "emptyActionType"
	return result
}

func (t *emptyActionType) instantiate() concept {
	return &emptyAction{
		commonAction: newCommonAction(),
	}
}

var emptyActionTypeSingleton = &emptyActionType{
	commonActionType: newCommonActionType(),
}
