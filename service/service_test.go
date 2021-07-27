package service

import (
	"github.com/muhammadisa/mocklearn/calculator/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

// serviceTS embedding suite.Suite of testify
type serviceTS struct {
	suite.Suite
}

// TestServiceTS initialize test suite
func TestServiceTS(t *testing.T) {
	suite.Run(t, new(serviceTS))
}

func (s *serviceTS) TestDoingSum() {

	tests := []struct {
		Name    string
		X, Y, R int
	}{
		{
			Name: "test doing sum number 1",
			X:    5, Y: 5, R: 10,
		},
		{
			Name: "test doing sum number 2",
			X:    4, Y: 5, R: 9,
		},
		{
			Name: "test doing sum number 3",
			X:    2, Y: 2, R: 4,
		},
	}

	for _, test := range tests {
		s.Run(test.Name, func() {
			calcMock := &mocks.Calculator{}

			calcMock.On("Sum", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(func(x, y int) int {
				return test.R
			})

			sv := NewService(calcMock)
			res := sv.DoingSum(test.X, test.Y)
			s.Assert().Equal(test.R, res)
		})
	}

}

func (s *serviceTS) TestDoingDiv() {

	tests := []struct {
		Name    string
		X, Y, R int
	}{
		{
			Name: "test doing div number 1",
			X:    4, Y: 2, R: 2,
		},
		{
			Name: "test doing div number 2",
			X:    10, Y: 2, R: 5,
		},
		{
			Name: "test doing div number 3",
			X:    16, Y: 2, R: 8,
		},
	}

	for _, test := range tests {
		s.Run(test.Name, func() {
			calcMock := &mocks.Calculator{}

			calcMock.On("Div", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(func(x, y int) int {
				return test.R
			})

			sv := NewService(calcMock)
			res := sv.DoingDiv(test.X, test.Y)
			s.Assert().Equal(test.R, res)
		})
	}

}

func (s *serviceTS) TestDoingTimes() {

	tests := []struct {
		Name    string
		X, Y, R int
	}{
		{
			Name: "test doing times number 1",
			X:    5, Y: 5, R: 25,
		},
		{
			Name: "test doing times number 2",
			X:    10, Y: 2, R: 20,
		},
		{
			Name: "test doing times number 3",
			X:    5, Y: 10, R: 50,
		},
	}

	for _, test := range tests {
		s.Run(test.Name, func() {
			calcMock := &mocks.Calculator{}

			calcMock.On("Times", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(func(x, y int) int {
				return test.R
			})

			sv := NewService(calcMock)
			res := sv.DoingTimes(test.X, test.Y)
			s.Assert().Equal(test.R, res)
		})
	}

}

func (s *serviceTS) TestDoingModulo() {

	tests := []struct {
		Name    string
		X, Y, R int
	}{
		{
			Name: "test doing modulo number 1",
			X:    5, Y: 5, R: 1,
		},
		{
			Name: "test doing modulo number 2",
			X:    15, Y: 5, R: 3,
		},
		{
			Name: "test doing modulo number 3",
			X:    20, Y: 4, R: 5,
		},
	}

	for _, test := range tests {
		s.Run(test.Name, func() {
			calcMock := &mocks.Calculator{}

			calcMock.On("Modulo", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(func(x, y int) int {
				return test.R
			})

			sv := NewService(calcMock)
			res := sv.DoingModulo(test.X, test.Y)
			s.Assert().Equal(test.R, res)
		})
	}

}
