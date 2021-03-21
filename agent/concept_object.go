package agent

type object interface {
	concept
	_type() objectType
	_attrs() map[int]int // type -> value

	// check if given object type match this object's type
	//   and given object attr include all of this object's attribute
	// newType == nil -> skip type match
	// attrTypes != nil -> only match these attribute types
	_match(newType objectType, newAttrs map[int]int, attrTypes map[int]bool) bool
	_update(newAttrs map[int]int) []change
}

type objectType interface {
	concept
	conceptType
	_attrs() map[int]map[int]int // type -> {val -> frequency}
	_match(newAttrs map[int]int, attrTypes map[int]bool) bool
	_update(oldAttrs map[int]int, newAttrs map[int]int)
}
