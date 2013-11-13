package calc

type Calculator struct{}

type Button uint32

const (
	Zero Button = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Add
	Subtract
	Multiply
	Divide
	Calculate // =
	Clear
)

func (calculator Calculator) DisplayValue() int64 {
	return 0
}

func (calculator Calculator) Press(button Button) {
}


