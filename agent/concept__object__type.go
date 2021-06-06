package agent

/*
	This type controls all generalizable knowledge of type
 */
type objectType interface {
	concept
	conceptType
	getAttrs() map[int]int // attribute type -> attribute value

	// similar to object.match, verify if given attributes match this objectType
	//  - newAttrs (attribute type -> attribute value): attributes to match, single mismatch would
	//    return false
	//  - consideredAttrTypes (set of attribute type): attribute types to match, match all
	//    attributes in this object if nil, otherwise ignore all other attribute types
	match(newAttrs map[int]int, consideredAttrTypes map[int]bool) bool
}

// the purpose of this struct is to remove duplicated code from implementations
type commonObjectType struct {
	*commonConcept
	attrs map[int]int // attribute type -> attribute value
}

func (t *commonObjectType) getAttrs() map[int]int {
	return t.attrs
}

// similar to object.match, verify if given attributes match this objectType
//  - newAttrs (attribute type -> attribute value): attributes to match, single mismatch would
//    return false
//  - consideredAttrTypes (set of attribute type): attribute types to match, match all
//    attributes in this object if nil, otherwise ignore all other attribute types
func (t *commonObjectType) match(newAttrs map[int]int, consideredAttrTypes map[int]bool) bool {
	for newAttrType, newAttrVal := range newAttrs {
		if consideredAttrTypes != nil && !consideredAttrTypes[newAttrType] {
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
		attrs: map[int]int{},
	}
}
