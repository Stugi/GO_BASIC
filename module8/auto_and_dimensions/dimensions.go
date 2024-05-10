package main

type DimensionsCM struct {
	length, width, height float64
}

type DimensionsInch struct {
	length, width, height float64
}

func (d DimensionsCM) Length() Unit {
	return Unit{Value: d.length, T: CM}
}

func (d DimensionsCM) Width() Unit {
	return Unit{Value: d.width, T: CM}
}

func (d DimensionsCM) Height() Unit {
	return Unit{Value: d.height, T: CM}
}

func (d DimensionsInch) Length() Unit {
	return Unit{Value: d.length, T: Inch}
}

func (d DimensionsInch) Width() Unit {
	return Unit{Value: d.width, T: Inch}
}

func (d DimensionsInch) Height() Unit {
	return Unit{Value: d.height, T: Inch}
}
