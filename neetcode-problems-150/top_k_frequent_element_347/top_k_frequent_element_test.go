package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type topKFreqElementTestSuite struct {
	suite.Suite
	inp []int
	k   int
}

func TestTopKFreqEleemntTestSuite(t *testing.T) {
	suite.Run(t, new(topKFreqElementTestSuite))
}

func (s *topKFreqElementTestSuite) SetupSuite() {
	s.inp = []int{1, 1, 1, 2, 2, 3}
	s.k = 2
}

func (s *topKFreqElementTestSuite) TestInvalidData() {
	s.inp = []int{}
	s.k = 2
	got := topKFrequent(s.inp, s.k)
	s.T().Log(s.inp)
	s.Empty(got)
}

func (s *topKFreqElementTestSuite) TestHappyScenarioOneLengthArrayInput() {
	s.inp = []int{1}
	s.k = 1
	got := topKFrequent(s.inp, s.k)
	s.Len(got, 1)
	s.Equal([]int{1}, got)
}

func (s *topKFreqElementTestSuite) TestInvalidKLength() {
	s.inp = []int{1}
	s.k = 2
	got := topKFrequent(s.inp, s.k)
	s.Len(got, 1)
	s.Equal([]int{1}, got)
}

func (s *topKFreqElementTestSuite) TestCorrectOutput() {
	got := topKFrequent(s.inp, s.k)
	s.Len(got, s.k) // the valid case actually and should be always this if we are sending the correct k
	s.Equal([]int{1, 2}, got)
}
