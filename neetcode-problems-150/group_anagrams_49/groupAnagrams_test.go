package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type groupAnagramsSuite struct {
	suite.Suite
	inp []string
}

func TestGroupAnagramsSuite(t *testing.T) {
	suite.Run(t, new(groupAnagramsSuite))
}

func (s *groupAnagramsSuite) SetupSuite() {
	s.inp = []string{
		"eat",
		"tea",
		"tan",
		"ate",
		"nat",
		"bat",
	}
}

// func (s *groupAnagramsSuite) TestNotEmptyResponse() {
// 	got := groupAnagrams(s.inp)
// 	s.NotEmpty(got)
// }

// func (s *groupAnagramsSuite) TestExpectedOutput() {
// 	got := groupAnagrams(s.inp)
// 	s.Equal([][]string{
// 		{"eat", "tea", "ate"},
// 		{"nat", "tan"},
// 		{"bat"},
// 	}, got)
// }
