package electronic

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	ButtonCount() int
}

type Smartphone interface {
	OS() string
}
