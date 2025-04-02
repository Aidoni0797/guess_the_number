package foo

import (
	"testing"
)

func TestFooFunc(t *testing.T) {
	expectedFooResult := "bar"
	if actualFooResult := Foo(); actualFooResult != expectedFooResult {
		t.Errorf("expected %s; got: %s", expectedFooResult, actualFooResult)
	}
}
