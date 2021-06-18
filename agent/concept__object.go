package agent

/*
	The object type is used to reference individual object instances, general knowledge on the object would be
	controlled by the associated objectType.
*/
type object interface {
	concept
	getType() objectType
	getAttrs() map[int]int // attribute type -> attribute value

	// attrs: attribute type -> attribute value
	// updates object attributes, construct and return appropriate changes
	setAttrs(a *Agent, attrs map[int]int) []change
	match(other concept) bool

	// check if given information matches this object
	//  - newAttrs (attribute type -> attribute value): attributes to match, single mismatch would
	//    return false
	//  - consideredAttrTypes (set of attribute type): attribute types to match, match all
	//    attributes in this object if nil, otherwise ignore all other attribute types
	matchAttrs(newAttrs map[int]int, consideredAttrTypes map[int]bool) bool
}

// the purpose of this struct is to remove duplicated code from implementations
type commonObject struct {
	*commonConcept
	attrs map[int]int // attribute type -> attribute value
}

func (o *commonObject) getAttrs() map[int]int {
	return o.attrs
}

// check if given information matches this object
//  - newAttrs (attribute type -> attribute value): attributes to match, single mismatch would
//    return false
//  - consideredAttrTypes (set of attribute type): attribute types to match, match all
//    attributes in this object if nil, otherwise ignore all other attribute types
func (o *commonObject) match(other concept) bool {
	otherObject, ok := other.(object)
	if !ok {
		return false
	}

	for oldType, oldVal := range o.attrs {
		if !visualAttrTypes[oldType] {
			continue
		}

		if newVal, seen := otherObject.getAttrs()[oldType]; !seen || oldVal != newVal {
			return false
		}
	}

	return true
}

// check if given information matches this object
//  - newAttrs (attribute type -> attribute value): attributes to match, single mismatch would
//    return false
//  - consideredAttrTypes (set of attribute type): attribute types to match, match all
//    attributes in this object if nil, otherwise ignore all other attribute types
func (o *commonObject) matchAttrs(newAttrs map[int]int, consideredAttrTypes map[int]bool) bool {
	for oldType, oldVal := range o.attrs {
		if consideredAttrTypes != nil && !consideredAttrTypes[oldType] {
			continue
		}

		if newVal, seen := newAttrs[oldType]; !seen || oldVal != newVal {
			return false
		}
	}

	return true
}

func newCommonObject() *commonObject {
	return &commonObject{
		commonConcept: newCommonConcept(),
		attrs:         map[int]int{},
	}
}
