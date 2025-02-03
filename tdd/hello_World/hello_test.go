package helloworld

import "testing"

func TestHelloWorld(t *testing.T) {
	got := Hello()
	expected := "hello world !"
	if got != expected {
		t.Errorf("got [%v] expected [%v]", got, expected)
	}
}
