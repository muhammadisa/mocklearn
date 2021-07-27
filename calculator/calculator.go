package calculator

type Calculator interface {
	Sum(int, int) int
	Div(int, int) int
	Times(int, int) int
	Modulo(int, int) int
}

type calculator struct{}

func (c calculator) Sum(x, y int) int {
	return x + y
}

func (c calculator) Div(x, y int) int {
	return x / y
}

func (c calculator) Times(x, y int) int {
	return x * y
}

func (c calculator) Modulo(x, y int) int {
	return x % y
}

func NewCalculator() Calculator {
	return &calculator{}
}