package chapter8

import (
	"github.com/go-bdd/gobdd"
	"testing"
)

func add(t gobdd.StepTest, ctx gobdd.Context, first int, second int) {
	res := first + second
	ctx.Set("result", res)
}

func check(t gobdd.StepTest, ctx gobdd.Context, sum int) {
	received, err := ctx.GetInt("result")
	if err != nil {
		t.Fatal(err)
		return
	}

	if sum != received {
		t.Fatalf("Expected %d, received %d", sum, received)
	}
}

func TestScenarios(t *testing.T) {
	suite := gobdd.NewSuite(t)
	suite.AddStep(`I add (\d+) and (\d+)`, add)
	suite.AddStep(`the result should equal (\d+)`, check)
	suite.Run()
}
