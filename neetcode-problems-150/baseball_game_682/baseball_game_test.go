package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type baseballGameSuite struct {
	suite.Suite
	inp []string
}

func TestBaseballGameSuite(t *testing.T) {
	suite.Run(t, new(baseballGameSuite))
}

func (s *baseballGameSuite) TestEmptyOperations() {
	tc := []string{}
	s.inp = tc
	got := calPoints(s.inp)
	s.Equal(0, got)
}

func (s *baseballGameSuite) TestOnlyOperationsWithoutNumbers() {
	
}
