package agent

import (
	"../util"
	"strconv"
)

type simpleObject struct {
	__assocs []concept
	__type   *simpleObjectType
	__attrs  map[int]int
}

func (o *simpleObject) _assocs() []concept {
	return o.__assocs
}

func (o *simpleObject) _toString() string {
	return "simpleObject"
}

func (o *simpleObject) _type() objectType {
	return o.__type
}

func (o *simpleObject) _attrs() map[int]int {
	return o.__attrs
}

func (o *simpleObject) _match(newType objectType, newAttrs map[int]int, attrTypes map[int]bool) bool {
	if newType != nil && o.__type != newType {
		return false
	}

	for oldType, oldVal := range o.__attrs {
		if attrTypes != nil && !attrTypes[oldType] {
			continue
		}

		if newVal, seen := newAttrs[oldType]; !seen || oldVal != newVal {
			return false
		}
	}
	return true
}

func (o *simpleObject) _update(newAttrs map[int]int) []change {
	var changes []change

	oldAttrs := make(map[int]int)
	for attrType, attrVal := range o.__attrs {
		oldAttrs[attrType] = attrVal
	}

	for newType, newVal := range newAttrs {
		if oldVal, seen := o.__attrs[newType]; seen {
			if oldVal != newVal {
				changes = append(changes, &simpleObjectChange{
					__objType:    o.__type,
					__attrType:   newType,
					__oldAttrVal: &oldVal,
					__newAttrVal: &newVal,
				})
			}
		} else {
			changes = append(changes, &simpleObjectChange{
				__objType:    o.__type,
				__attrType:   newType,
				__oldAttrVal: nil,
				__newAttrVal: &newVal,
			})
		}
		o.__attrs[newType] = newVal
	}

	o.__type._update(oldAttrs, newAttrs)

	return changes
}

type simpleObjectType struct {
	__assocs []concept
	__attrs  map[int]map[int]int // attr type -> {attr val -> frequency}
	__count  int
}

func (t *simpleObjectType) _assocs() []concept {
	return t.__assocs
}

func (t *simpleObjectType) _toString() string {
	return "simpleObjectType"
}

func (t *simpleObjectType) _instantiate() concept {
	defaultAttrs := make(map[int]int)

	for attrType, attrDistribution := range t.__attrs {
		for attrVal, freq := range attrDistribution {
			if util.IsSignificant(t.__count, freq) {
				defaultAttrs[attrType] = attrVal
				break
			}
		}
	}

	newObj := &simpleObject{
		__type:  t,
		__attrs: defaultAttrs,
	}

	// self-reinforcement could be bad, not sure, but then again if we keep imagining purple apples
	//    maybe we're eventually going to believe apples are purple too?
	t._update(nil, newObj.__attrs)
	return newObj
}

func (t *simpleObjectType) _attrs() map[int]map[int]int {
	return t.__attrs
}

func (t *simpleObjectType) _match(newAttrs map[int]int, attrTypes map[int]bool) bool {
	for oldType, oldDistribution := range t.__attrs {
		if attrTypes != nil && !attrTypes[oldType] {
			continue
		}

		if newVal, seen := newAttrs[oldType]; !seen || oldDistribution[newVal] == 0 {
			return false
		}
	}
	return true
}

func (t *simpleObjectType) _update(oldAttrs map[int]int, newAttrs map[int]int) {
	for attrType, attrVal := range newAttrs {
		if _, seen := t.__attrs[attrType]; !seen {
			t.__attrs[attrType] = make(map[int]int)
		}
		t.__attrs[attrType][attrVal]++
	}

	if oldAttrs == nil {
		t.__count++
		for attrType, attrVal := range oldAttrs {
			if _, seen := t.__attrs[attrType]; !seen {
				t.__attrs[attrType] = make(map[int]int)
			}
			if t.__attrs[attrType][attrVal] > 0 {
				t.__attrs[attrType][attrVal]--
			}
		}
	}
}

type simpleObjectChange struct {
	__assocs      []concept
	__objType     *simpleObjectType
	__attrType    int
	__oldAttrVal  *int
	__newAttrVal  *int
	__changeValue float64
}

func (c *simpleObjectChange) _assocs() []concept {
	return c.__assocs
}

func (c *simpleObjectChange) _toString() string {
	return "simpleObjectChange [" + strconv.Itoa(c.__attrType) + ", " + strconv.Itoa(*c.__oldAttrVal) + "->" + strconv.Itoa(*c.__newAttrVal) + "]"
}

func (c *simpleObjectChange) _match(n change) bool {
	nn, ok := n.(*simpleObjectChange)
	// TODO reduce granularity
	return ok && c.__objType == nn.__objType && c.__attrType == nn.__attrType &&
		c.__oldAttrVal == nn.__oldAttrVal && c.__newAttrVal == nn.__newAttrVal
}

func (c *simpleObjectChange) _matchOld(a *Agent, oldItem concept) bool {
	oldObj, ok := oldItem.(*simpleObject)
	if !ok {
		return false
	}

	if oldObj.__type != c.__objType {
		return false
	}

	if oldVal, seen := oldObj.__attrs[c.__attrType]; !seen || oldVal != *c.__oldAttrVal {
		return false
	}

	return true
}

func (c *simpleObjectChange) _matchNew(a *Agent, newItem concept) bool {
	newObj, ok := newItem.(*simpleObject)
	if !ok {
		return false
	}

	if newObj.__type != c.__objType {
		return false
	}

	if newVal, seen := newObj.__attrs[c.__attrType]; !seen || newVal != *c.__newAttrVal {
		return false
	}

	return true
}

func (c *simpleObjectChange) _getValue() float64 {
	return c.__changeValue
}

func (c *simpleObjectChange) _setValue(newValue float64) {
	c.__changeValue = newValue
}
