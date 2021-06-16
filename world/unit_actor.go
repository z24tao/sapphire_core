package world

import "math/rand"

type actorUnit struct {
	__id    int
	__board *board
	__name  string
}

func (u *actorUnit) _id() int {
	return u.__id
}

func (u *actorUnit) getType() int {
	return unitTypeActor
}

func (u *actorUnit) _board() *board {
	return u.__board
}

func (u *actorUnit) _color() int {
	return blue
}

func (u *actorUnit) _shape() int {
	return rectangle
}

func (u *actorUnit) _name() string {
	return u.__name
}

func (u *actorUnit) _eatenResponse() []interface{} {
	return []interface{}{}
}

func NewActor() int {
	id := newUnitId()
	a := &actorUnit{
		__id:    id,
		__board: defaultBoard,
		__name:  "actor",
	}

	units[id] = a
	if !defaultBoard.addUnitAt(a, [2]int{rand.Intn(10), rand.Intn(10)}) {
		panic("board out of space")
	}

	return id
}
