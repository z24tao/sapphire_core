package world

const (
	unitTypeItem = iota
	unitTypeActor
)

type unit interface {
	_id() int
	_type() int
	_board() *board
	_color() int
	_shape() int
	_eatenResponse() []interface{}
}
