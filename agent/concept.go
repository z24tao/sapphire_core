package agent

type concept interface {
	_assocs() []concept // TODO: maybe this needs to be a map from concept to float64 to store assoc strength
	_toString() string
}

type conceptType interface {
	_instantiate() concept
}
