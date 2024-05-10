package main

type BMW struct {
	model       string
	maxSpeed    int
	enginePower int
	dimensions  Dimensions
}

func NewBMW(model string, maxSpeed, enginePower int, length, width, height float64) *BMW {
	return &BMW{
		model:       model,
		maxSpeed:    maxSpeed,
		enginePower: enginePower,
		dimensions:  DimensionsCM{length, width, height},
	}
}

func (b *BMW) Brand() string {
	return "BMW"
}

func (b *BMW) Model() string {
	return b.model
}

func (b *BMW) MaxSpeed() int {
	return b.maxSpeed
}

func (b *BMW) EnginePower() int {
	return b.enginePower
}

func (b *BMW) Dimensions() Dimensions {
	return b.dimensions
}

type Mercedes struct {
	model       string
	maxSpeed    int
	enginePower int
	dimensions  Dimensions
}

func NewMercedes(model string, maxSpeed, enginePower int, length, width, height float64) *Mercedes {
	return &Mercedes{
		model:       model,
		maxSpeed:    maxSpeed,
		enginePower: enginePower,
		dimensions:  DimensionsCM{length, width, height},
	}
}

func (m *Mercedes) Brand() string {
	return "Mercedes"
}

func (m *Mercedes) Model() string {
	return m.model
}

func (m *Mercedes) MaxSpeed() int {
	return m.maxSpeed
}

func (m *Mercedes) EnginePower() int {
	return m.enginePower
}

func (m *Mercedes) Dimensions() Dimensions {
	return m.dimensions
}

type Dodge struct {
	model       string
	maxSpeed    int
	enginePower int
	dimensions  Dimensions
}

func NewDodge(model string, maxSpeed, enginePower int, length, width, height float64) *Dodge {
	return &Dodge{
		model:       model,
		maxSpeed:    maxSpeed,
		enginePower: enginePower,
		dimensions:  DimensionsInch{length, width, height},
	}
}

func (d *Dodge) Brand() string {
	return "Dodge"
}

func (d *Dodge) Model() string {
	return d.model
}

func (d *Dodge) MaxSpeed() int {
	return d.maxSpeed
}

func (d *Dodge) EnginePower() int {
	return d.enginePower
}

func (d *Dodge) Dimensions() Dimensions {
	return d.dimensions
}
