package world

const (
	black = iota
	white
	red
	orange
	yellow
	green
	blue
	purple
)

const (
	circle = iota
	triangle
	rectangle
)

type Image struct {
	XDist     int
	ZDist     int
	Color     int
	Shape     int
	DebugName string
}

var Colors = map[int]string {
	black: "black",
	white: "white",
	red: "red",
	orange: "orange",
	yellow: "yellow",
	green: "green",
	blue: "blue",
	purple: "purple",
}

var Shapes = map[int]string {
	circle: "circle",
	triangle: "triangle",
	rectangle: "rectangle",
}

type Taste struct {
	Nutrition int
	Sweet     bool
}
