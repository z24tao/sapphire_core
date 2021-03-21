package world

type tile struct {
	occupants map[int]unit
}

func newTile() *tile {
	return &tile{
		occupants: make(map[int]unit),
	}
}
