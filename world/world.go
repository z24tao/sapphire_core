package world

var unitIds = 0
var units = make(map[int]unit)
var currTime int
var scheduledEvents map[int][]func(int)
var scheduledEventParams map[int][]int

func newUnitId() int {
	unitIds++
	return unitIds
}

func Look(unitId int) []*Image {
	currTime++
	if currTimeEvents, seen := scheduledEvents[currTime]; seen {
		for i, currTimeEvent := range currTimeEvents {
			currTimeEvent(scheduledEventParams[currTime][i])
		}
		delete(scheduledEvents, currTime)
		delete(scheduledEventParams, currTime)
	}

	var result []*Image

	u, seen := units[unitId]
	b := u._board()
	if seen && b != nil {
		for v := range b.units {
			img := b.look(u, v)
			if img != nil {
				result = append(result, img)
			}
		}
	}

	return result
}

func scheduleEvent(eventTime int, event func(int), param int) {
	if _, seen := scheduledEvents[eventTime]; !seen {
		scheduledEvents = map[int][]func(int){}
		scheduledEventParams = map[int][]int{}
	}
	scheduledEvents[eventTime] = append(scheduledEvents[eventTime], event)
	scheduledEventParams[eventTime] = append(scheduledEventParams[eventTime], param)
}

func init() {
	side := 10
	defaultTiles := make([][]*tile, side)
	for i := 0; i < side; i++ {
		defaultTiles[i] = make([]*tile, side)
		for j := 0; j < side; j++ {
			defaultTiles[i][j] = newTile()
		}
	}

	defaultBoard = &board{
		tiles: defaultTiles,
		units: make(map[unit][2]int),
	}

	addRandomItem(newApple())
	//addRandomItem(newOrange())
	//addRandomItem(newLemon())
	//addRandomItem(newLime())
	//addRandomItem(newBlueberry())
	currTime = 0
	scheduledEvents = map[int][]func(int){}
	scheduledEventParams = map[int][]int{}
}
