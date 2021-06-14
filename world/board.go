package world

import (
	"fmt"
)

type board struct {
	tiles [][]*tile
	units map[unit][2]int
}

func (b *board) addUnit(u unit) bool {
	for x, row := range b.tiles {
		for z, t := range row {
			if t.occupants[u.getType()] == nil {
				t.occupants[u.getType()] = u
				b.units[u] = [2]int{x, z}
				return true
			}
		}
	}

	return false
}

func (b *board) addUnitAt(u unit, pos [2]int) bool {
	t := b.tiles[pos[0]][pos[1]]
	fmt.Printf("create %s at %d, %d \n", u._name(), pos[0], pos[1])
	if t.occupants[u.getType()] == nil {
		t.occupants[u.getType()] = u
		b.units[u] = pos
		return true
	}

	return false
}

func (b *board) removeUnit(u unit) {
	b.tiles[b.units[u][0]][b.units[u][1]].occupants[u.getType()] = nil
	delete(b.units, u)
}

func (b *board) moveUnit(u unit, dx, dz int) bool {
	src, seen := b.units[u]
	if !seen {
		//fmt.Println("failed to move, unit does not exist")
		return false
	}

	dest := [2]int{src[0] + dx, src[1] + dz}
	fmt.Println("unit moving from ", src, " to ", dest, " type ", u._name())
	if dest[0] < 0 || dest[0] >= len(b.tiles) ||
		dest[1] < 0 || dest[1] >= len(b.tiles[0]) {
		//fmt.Println("failed to move, destination out of bounds")
		return false
	}

	if b.tiles[dest[0]][dest[1]].occupants[u.getType()] != nil {
		//fmt.Println("failed to move, destination occupied by " + b.tiles[dest[0]][dest[1]].occupants[u.getType()]._name())
		//fmt.Println(b.tiles[dest[0]][dest[1]].occupants)
		//fmt.Println(b.tiles[dest[0]][dest[1]].occupants[u.getType()] == u)
		//for _, row := range b.tiles {
		//	for _, col := range row {
		//		if col.occupants[unitTypeActor] != nil && col.occupants[unitTypeItem] != nil {
		//			fmt.Printf("3 ")
		//		} else if col.occupants[unitTypeActor] != nil {
		//			fmt.Printf("2 ")
		//		} else if col.occupants[unitTypeItem] != nil {
		//			fmt.Printf("1 ")
		//		} else {
		//			fmt.Printf("0 ")
		//		}
		//	}
		//	fmt.Println()
		//}
		return false
	}

	b.tiles[dest[0]][dest[1]].occupants[u.getType()] = u
	delete(b.tiles[src[0]][src[1]].occupants, u.getType())
	b.units[u] = dest
	return true
}

func (b *board) unitEat(u unit) []interface{} {
	t := b.tiles[b.units[u][0]][b.units[u][1]]
	var response []interface{}
	eatenItem := t.occupants[unitTypeItem]
	if eatenItem != nil {
		fmt.Println("ate apple at", b.units[u][0], b.units[u][1])
		response = eatenItem._eatenResponse()
		b.removeUnit(eatenItem)
		scheduleEvent(currTime+2, addRandomApple)
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
		XDist:     vPos[0] - uPos[0],
		ZDist:     vPos[1] - uPos[1],
		Color:     v._color(),
		Shape:     v._shape(),
		DebugName: v._name(),
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
				fmt.Println("AAI DISABLED ", aai.Name)
				response = append(response, &AtomicActionInterfaceChange{
					Interface: aai,
					Enabling:  false,
				})
			}
			if !oldEnabled && newEnabled {
				fmt.Println("AAI ENABLED", aai.Name)
				response = append(response, &AtomicActionInterfaceChange{
					Interface: aai,
					Enabling:  true,
				})
			}
			aai.Enabled = newEnabled
		}
	}
	return response
}

func (b *board) getState() *BoardState {
	units := make([]*unitState, 0)

	for u, pos := range b.units {
		units = append(units, &unitState{
			UnitType: u._name(),
			XPos:     pos[0],
			ZPos:     pos[1],
		})
	}

	return &BoardState{
		XDim:  len(b.tiles),
		ZDim:  len(b.tiles[0]),
		Units: units,
	}
}

var defaultBoard *board
