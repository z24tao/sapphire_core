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
	XDist int
	ZDist int
	Color int
	Shape int
}

var Colors = []int{
	black,
	white,
	red,
	orange,
	yellow,
	green,
	blue,
	purple,
}

var Shapes = []int{
	circle,
	triangle,
	rectangle,
}

type Taste struct {
	Nutrition int
	Sweet bool
}
