package agent

import "../world"

const imageDefaultImportance = 0.5
const changeDefaultImportance = 0.3

func (a *Agent) _observe() {
	imgs := world.Look(a._unitId)
	for _, img := range imgs {
		a._processImage(img)
	}
}

func (a *Agent) _processImage(img *world.Image) {
	attrs := a._identifyAttributes(img)
	obj := a._identifyObjectInst(attrs)

	if obj == nil {
		a._createObject(attrs)
	}
}

func (a *Agent) _identifyAttributes(img *world.Image) map[int]int {
	attrs := make(map[int]int)

	attrs[attrTypeColor] = img.Color
	attrs[attrTypeShape] = img.Shape

	if img.XDist == 0 && img.ZDist == 0 {
		attrs[attrTypeDirection] = directionOrigin
	} else if img.XDist > 0 && img.ZDist == 0 {
		attrs[attrTypeDirection] = directionXPos
	} else if img.XDist < 0 && img.ZDist == 0 {
		attrs[attrTypeDirection] = directionXNeg
	} else if img.XDist == 0 && img.ZDist > 0 {
		attrs[attrTypeDirection] = directionZPos
	} else if img.XDist == 0 && img.ZDist < 0 {
		attrs[attrTypeDirection] = directionZNeg
	} else if img.XDist > 0 && img.ZDist > 0 {
		attrs[attrTypeDirection] = directionXPosZPos
	} else if img.XDist > 0 && img.ZDist < 0 {
		attrs[attrTypeDirection] = directionXPosZNeg
	} else if img.XDist < 0 && img.ZDist > 0 {
		attrs[attrTypeDirection] = directionXNegZPos
	} else if img.XDist < 0 && img.ZDist < 0 {
		attrs[attrTypeDirection] = directionXNegZNeg
	}

	return attrs
}

func (a *Agent) _identifyObjectInst(attrs map[int]int) object {
	mindObjs := a._mind._objects()
	for _, obj := range mindObjs {
		if obj._match(nil, attrs, visualAttrTypes) {
			changes := obj._update(attrs)
			a._updateActionOutcome(changes)
			return obj
		}
	}

	return nil
}

func (a *Agent) _updateActionOutcome(changes []change) {
	for _, c := range changes {
		for _, ac := range a._activity._activeActions {
			matched := false
			for _, currCausation := range ac._type()._getCausations() {
				if currCausation._change._match(c) {
					matched = true
					currCausation._occurrences ++
					break
				}
			}
			if !matched {
				ac._type()._setCausations(append(ac._type()._getCausations(), &causation{
					_change: c,
					_occurrences: 1,
				}))
			}
		}

		a._mind._addItem(c, changeDefaultImportance)
	}
}

func (a *Agent) _createObject(attrs map[int]int) {
	newType := a._identifyObjectType(attrs)
	if newType == nil {
		newType = a._createObjectType(attrs)
	}

	a._createObjectInst(attrs, newType)
}

func (a *Agent) _identifyObjectType(attrs map[int]int) objectType {
	mindObjTypes := a._mind._objectTypes()
	for _, objType := range mindObjTypes {
		if objType._match(attrs, visualAttrTypes) {
			return objType
		}
	}

	return nil
}

func (a *Agent) _createObjectType(attrs map[int]int) objectType {
	objType := &simpleObjectType{
		__attrs: make(map[int]map[int]int),
		__count: 0,
	}

	a._mind._addItem(objType, imageDefaultImportance)
	return objType
}

func (a *Agent) _createObjectInst(attrs map[int]int, newType objectType) object {
	newInst := newType._instantiate().(object)
	newInst._update(attrs)
	a._mind._addItem(newInst, imageDefaultImportance)
	return newInst
}
