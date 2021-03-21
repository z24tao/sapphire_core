package world

type itemUnit struct {
	__id int
	__board *board
	__color int
	__shape int
	__nutrition int
	__isSweet bool
}

func (u *itemUnit) _id() int {
	return u.__id
}

func (u *itemUnit) _type() int {
	return unitTypeItem
}

func (u *itemUnit) _board() *board {
	return u.__board
}

func (u *itemUnit) _color() int {
	return u.__color
}

func (u *itemUnit) _shape() int {
	return u.__shape
}

func (u *itemUnit) _eatenResponse() []interface{} {
	return []interface{}{
		&Taste{
			u.__nutrition,
			u.__isSweet,
		},
	}
}

func newApple() int {
	id := newUnitId()
	a := &itemUnit{
		__id:    id,
		__board: defaultBoard,
		__color: red,
		__shape: circle,
		__isSweet: true,
		__nutrition: 15,
	}

	units[id] = a
	if !defaultBoard.addUnit(a) {
		panic("board out of space")
	}

	return id
}
