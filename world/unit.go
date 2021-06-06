package world

const (
	unitTypeItem = iota
	unitTypeActor
)

type unit interface {
	_id() int
	getType() int
	_board() *board
	_color() int
	_shape() int
	_name() string
	_eatenResponse() []interface{}
}
