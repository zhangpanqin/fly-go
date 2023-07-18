package test

import (
	"testing"
)

func Test_Sum1(t *testing.T) {
	t.Setenv("a", "b")
	if sum := Sum1(2); sum != 3 {
		t.FailNow()
	}
}
