package agent

/*
	This type controls all generalizable knowledge of type
*/
type objectType interface {
	concept
	conceptType
	getAttrs() map[int]int // attribute type -> attribute value
}

// the purpose of this struct is to remove duplicated code from implementations
type commonObjectType struct {
	*commonConcept
	attrs map[int]int // attribute type -> attribute value
}

func (t *commonObjectType) getAttrs() map[int]int {
	return t.attrs
}

func (t *commonObjectType) match(other concept) bool {
	otherObjectType, ok := other.(objectType)
	if !ok {
		return false
	}

	for newAttrType, newAttrVal := range otherObjectType.getAttrs() {
		if !visualAttrTypes[newAttrType] {
			continue
		}

		if newAttrVal != t.attrs[newAttrType] {
			return false
		}
	}

	return true
}

func newCommonObjectType() *commonObjectType {
	return &commonObjectType{
		commonConcept: newCommonConcept(),
		attrs:         map[int]int{},
	}
}
