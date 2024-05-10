package main

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {
	value := u.Value

	if t != u.T {
		// сконвертировать value в заданный в параметре UnitType
		if u.T == CM && t == Inch {
			value = value * 0.39
		}

		if u.T == Inch && t == CM {
			value = value * 2.54
		}

	}

	return value
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}
