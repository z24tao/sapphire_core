package world

import (
	"encoding/json"
	"fmt"
)

type board struct {
	tiles [][]*tile
	units map[unit][2]int
}

func (b *board) addUnit(u unit) bool {
	for x, row := range b.tiles {
		for z, t := range row {
			if t.occupants[u._type()] == nil {
				t.occupants[u._type()] = u
				b.units[u] = [2]int{x, z}
				return true
			}
		}
	}

	return false
}

func (b *board) addUnitAt(u unit, pos [2]int) bool {
	t := b.tiles[pos[0]][pos[1]]
	if t.occupants[u._type()] == nil {
		t.occupants[u._type()] = u
		b.units[u] = pos
		return true
	}

	return false
}

func (b *board) removeUnit(u unit) {
	b.tiles[b.units[u][0]][b.units[u][1]].occupants[u._type()] = nil
	delete(b.units, u)
}

func (b *board) moveUnit(u unit, dx, dz int) bool {
	src, seen := b.units[u]
	if !seen {
		return false
	}

	dest := [2]int{src[0] + dx, src[1] + dz}
	fmt.Println("unit moving from ", src, " to ", dest)
	if dest[0] < 0 || dest[0] >= len(b.tiles) ||
		dest[1] < 0 || dest[1] >= len(b.tiles[0]) {
		return false
	}

	if b.tiles[dest[0]][dest[1]].occupants[u._type()] != nil {
		return false
	}

	b.tiles[dest[0]][dest[1]].occupants[u._type()] = u
	b.tiles[src[0]][src[1]].occupants[u._type()] = u
	b.units[u] = dest
	return true
}

func (b *board) unitEat(u unit) []interface{} {
	t := b.tiles[b.units[u][0]][b.units[u][1]]
	var response []interface{}
	eatenItem := t.occupants[unitTypeItem]
	if eatenItem != nil {
		response = eatenItem._eatenResponse()
		b.removeUnit(eatenItem)
	}
	return response
}

func (b *board) look(u, v unit) *Image {
	if u == v {
		return nil
	}
	uPos, seen := b.units[u]
	if !seen {
		return nil
	}
	vPos, seen := b.units[v]
	if !seen {
		return nil
	}

	return &Image{
		XDist: vPos[0] - uPos[0],
		ZDist: vPos[1] - uPos[1],
		Color: v._color(),
		Shape: v._shape(),
	}
}

func (b *board) updateAAI(u unit) []interface{} {
	aais := unitAAI[u._id()]
	var response []interface{}
	for aai := range aais {
		if unitAAIAction[aai] == aaiEat {
			oldEnabled := aai.Enabled
			newEnabled := b.tiles[b.units[u][0]][b.units[u][1]].occupants[unitTypeItem] != nil
			if oldEnabled && !newEnabled {
				response = append(response, &AtomicActionInterfaceChange{
					Interface: aai,
					Enabling: false,
				})
			}
			if !oldEnabled && newEnabled {
				response = append(response, &AtomicActionInterfaceChange{
					Interface: aai,
					Enabling: true,
				})
			}
			aai.Enabled = newEnabled
		}
	}
	return response
}

func (b *board) getState() *BoardState {
	actors := make([][]string, 0)
	items := make([][]string, 0)
	for range b.tiles {
		actors = append(actors, make([]string, len(b.tiles[0])))
		items = append(items, make([]string, len(b.tiles[0])))
	}

	for x, row := range b.tiles {
		for z, t := range row {
			if t.occupants[unitTypeItem] != nil {
				items[x][z] = "apple"
			}
			if t.occupants[unitTypeActor] != nil {
				actors[x][z] = "actor"
			}
		}
	}

	return &BoardState{
		XDim: len(b.tiles),
		ZDim: len(b.tiles[0]),
		Actors: actors,
		Items: items,
	}
}

var defaultBoard *board

type BoardState struct {
	XDim int `json:"x_dim"`
	ZDim int `json:"z_dim"`
	Actors [][]string `json:"actors"`
	Items [][]string `json:"items"`
}

func GetDefaultBoardState() string {
	data, err := json.Marshal(defaultBoard.getState())
	if err != nil {
		return ""
	}
	return string(data)
}
