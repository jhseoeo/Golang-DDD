package chapter3_test

import (
	"github.com/jhseoeo/Golang-DDD/chapter3"
	"testing"
)

func Test_Point(t *testing.T) {
	a := chapter3.NewPoint(1, 2)
	b := chapter3.NewPoint(1, 2)

	if a != b {
		t.Fatal("a and b were not equal")
	}
}
