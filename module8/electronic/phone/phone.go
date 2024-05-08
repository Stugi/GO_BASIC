package electronic

type applePhone struct {
	model string
}

func NewApplePhone(model string) *applePhone {
	return &applePhone{
		model: model,
	}
}

func (a applePhone) Brand() string {
	return "apple"
}

func (a applePhone) Model() string {
	return a.model
}

func (a applePhone) Type() string {
	return "smartphone"
}

func (a applePhone) OS() string {
	return "ios"
}

type androidPhone struct {
	brand string
	model string
}

func NewAndroidPhone(brand, model string) *androidPhone {
	return &androidPhone{
		brand: brand,
		model: model,
	}
}

func (a androidPhone) Brand() string {
	return a.brand
}

func (a androidPhone) Model() string {
	return a.model
}

func (a androidPhone) Type() string {
	return "smartphone"
}

func (a androidPhone) OS() string {
	return "android"
}

type radioPhone struct {
	brand       string
	model       string
	buttonCount int
}

func NewRadioPhone(brand, model string, buttonCount int) *radioPhone {
	return &radioPhone{
		brand:       brand,
		model:       model,
		buttonCount: buttonCount,
	}
}

func (a radioPhone) Brand() string {
	return a.brand
}

func (a radioPhone) Model() string {
	return a.model
}

func (a radioPhone) Type() string {
	return "station"
}

func (a radioPhone) ButtonCount() int {
	return a.buttonCount
}
