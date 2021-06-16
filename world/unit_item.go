package world

import "math/rand"

type itemUnit struct {
	__id        int
	__board     *board
	__color     int
	__shape     int
	__nutrition int
	__isSweet   bool
	__name      string
}

func (u *itemUnit) _id() int {
	return u.__id
}

func (u *itemUnit) getType() int {
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

func (u *itemUnit) _name() string {
	return u.__name
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
		__id:        id,
		__board:     defaultBoard,
		__color:     red,
		__shape:     circle,
		__isSweet:   true,
		__nutrition: 15,
		__name:      "apple",
	}

	units[id] = a
	return id
}

func newOrange() int {
	id := newUnitId()
	a := &itemUnit{
		__id:        id,
		__board:     defaultBoard,
		__color:     orange,
		__shape:     circle,
		__isSweet:   true,
		__nutrition: 15,
		__name:      "orange",
	}

	units[id] = a
	return id
}

func newLemon() int {
	id := newUnitId()
	a := &itemUnit{
		__id:        id,
		__board:     defaultBoard,
		__color:     yellow,
		__shape:     circle,
		__isSweet:   true,
		__nutrition: 15,
		__name:      "lemon",
	}

	units[id] = a
	return id
}

func newLime() int {
	id := newUnitId()
	a := &itemUnit{
		__id:        id,
		__board:     defaultBoard,
		__color:     green,
		__shape:     circle,
		__isSweet:   true,
		__nutrition: 15,
		__name:      "lime",
	}

	units[id] = a
	return id
}

func newBlueberry() int {
	id := newUnitId()
	a := &itemUnit{
		__id:        id,
		__board:     defaultBoard,
		__color:     blue,
		__shape:     circle,
		__isSweet:   true,
		__nutrition: 15,
		__name:      "blueberry",
	}

	units[id] = a
	return id
}

func addRandomItem(itemId int) {
	xLen, zLen := len(defaultBoard.tiles), len(defaultBoard.tiles[0])
	for defaultBoard.addUnitAt(units[itemId], [2]int{rand.Intn(xLen-2) + 1, rand.Intn(zLen-2) + 1}) == false {
	}
}
