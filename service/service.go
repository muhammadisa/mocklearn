package service

import "github.com/muhammadisa/mocklearn/calculator"

type Service interface {
	DoingSum(int, int) int
	DoingDiv(int, int) int
	DoingTimes(int, int) int
	DoingModulo(int, int) int
}

type service struct {
	Calculator calculator.Calculator
}

func (s service) DoingSum(x, y int) int {
	return s.Calculator.Sum(x, y)
}

func (s service) DoingDiv(x, y int) int {
	return s.Calculator.Div(x, y)
}

func (s service) DoingTimes(x, y int) int {
	return s.Calculator.Times(x, y)
}

func (s service) DoingModulo(x, y int) int {
	return s.Calculator.Modulo(x, y)
}

func NewService(calc calculator.Calculator) Service {
	return &service{Calculator: calc}
}
