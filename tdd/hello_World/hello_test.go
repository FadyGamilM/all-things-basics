package helloworld

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestHelloWorld(t *testing.T) {
	asserts := assert.New(t)

	t.Run("is_it_greeting_everyone", func(t *testing.T) {
		got := Hello("everyone")
		expected := "hello everyone !"
		asserts.Equal(expected, got)
	})

	t.Run("testing_hello_world", func(t *testing.T) {
		got := Hello("world")
		expected := "hello world !?"
		asserts.Equal(expected, got)
	})

}

// to act as our test suite
type SimpleTestSuite struct {
	suite.Suite // so the test runner treats this struct as test suite
}

// to tell the test runner to execute this suite when we run the go test cmd
func TestSimpleTestSuite(t *testing.T) {
	suite.Run(t, &SimpleTestSuite{})
}

func (s *SimpleTestSuite) TestTrue() {
	s.True(true)
}
