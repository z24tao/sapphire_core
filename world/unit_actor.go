package world

type actorUnit struct {
	__id int
	__board *board
}

func (u *actorUnit) _id() int {
	return u.__id
}

func (u *actorUnit) _type() int {
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

func (u *actorUnit) _eatenResponse() []interface{} {
	return []interface{}{}
}

func NewActor() int {
	id := newUnitId()
	a := &actorUnit{
		__id:    id,
		__board: defaultBoard,
	}

	units[id] = a
	if !defaultBoard.addUnit(a) {
		panic("board out of space")
	}

	return id
}
